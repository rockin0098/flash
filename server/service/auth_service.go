package service

import (
	"math/rand"
	"sync"
	"time"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
)

const (
	SALT_PREFIX  = "salt"
	SALT_TIMEOUT = 30 * 60 // salt timeout
)

type AuthService struct {
	authStorage *sync.Map
}

var authService = &AuthService{
	authStorage: &sync.Map{},
}

func AuthServiceInstance() *AuthService {
	return authService
}

func (s *AuthService) Store(authid int64, auth *Auth) {
	s.authStorage.Store(authid, auth)
}

func (s *AuthService) Load(authid int64) (*Auth, bool) {
	obj, ok := s.authStorage.Load(authid)
	if !ok {
		Log.Warnf("authid = [%v] does not exist", authid)
		return nil, false
	}

	return obj.(*Auth), true
}

// https://core.telegram.org/mtproto/description#server-salt

// Server Salt
//
// A (random) 64-bit number periodically (say, every 24 hours) changed
// (separately for each session) at the request of the server.
// All subsequent messages must contain the new salt
// (although, messages with the old salt are still accepted for a further 300 seconds).
// Required to protect against replay attacks and certain tricks
// associated with adjusting the client clock to a moment in the distant future.
//
func (s *AuthService) CheckBySalt(authid int64, salt int64) bool {
	var date = int32(time.Now().Unix())

	// var salt int64 = 0
	auth, ok := s.Load(authid)
	if !ok {
		mm := model.GetModelManager()
		a := mm.GetAuthKeyByAuthID(authid)
		if a == nil {
			return false
		}

		auth = &Auth{
			AuthKeyID: authid,
		}
		s.Store(authid, auth)
	}

	salts := auth.Salts

	Log.Infof("date = %v, salts = %v, salt = %v", date, salts, salt)

	if len(salts) > 0 {
		for _, v := range salts {
			// old salt are still accepted for a further 300 seconds
			if v.M_valid_since <= date && v.M_valid_until+SALT_TIMEOUT > date && salt == v.M_salt {
				return true
			}
		}
	}

	return false
}

func (s *AuthService) GetOrInsertSalt(authid int64) (int64, error) {
	var date = int32(time.Now().Unix())

	var salt int64 = 0
	auth, ok := s.Load(authid)
	if ok {
		if len(auth.Salts) > 0 {
			for _, v := range auth.Salts {
				if v.M_valid_since <= date && v.M_valid_until > date {
					salt = v.M_salt
					break
				}
			}
		}
	}

	if salt == 0 {
		salt = rand.Int63()
		saltData := &mtproto.TL_future_salt{
			M_valid_since: date,
			M_valid_until: date + SALT_TIMEOUT + SALT_TIMEOUT,
			M_salt:        salt,
		}

		auth.Salts = []*mtproto.TL_future_salt{saltData}
	}

	return salt, nil
}

func (s *AuthService) GetOrInsertSaltList(authid int64, size int) ([]*mtproto.TL_future_salt, error) {
	var (
		salts = make([]*mtproto.TL_future_salt, size)
		// saltKey =

		date           = int32(time.Now().Unix())
		lastValidUntil = date

		// ok = false
		saltsData []*mtproto.TL_future_salt
	)

	v, ok := s.Load(authid)
	if ok {
		if len(v.Salts) > 0 {
			for _, salt := range v.Salts {
				if salt.M_valid_until >= date {
					saltsData = append(saltsData, salt)
					if lastValidUntil < salt.M_valid_until {
						lastValidUntil = salt.M_valid_until
					}
				}
			}
		}
	}

	left := size - len(saltsData)
	if left > 0 {
		for i := 0; i < size; i++ {
			salt := &mtproto.TL_future_salt{
				M_valid_since: lastValidUntil,
				M_valid_until: lastValidUntil + SALT_TIMEOUT,
				M_salt:        rand.Int63(),
			}
			saltsData = append(saltsData, salt)
			lastValidUntil += SALT_TIMEOUT
		}
	}

	for i := 0; i < size; i++ {
		salts[i] = saltsData[i]
	}

	if left > 0 {
		// err := cacheSalts.cache.Put(genCacheSaltKey(keyId), saltsData, time.Duration(len(saltsData))*kSaltTimeout*time.Second)
		// if err != nil {
		// 	glog.Error(err)
		// 	return nil, err
		// }
		v.Salts = salts
	}
	return salts, nil
}

// auth
type Auth struct {
	// rw        *sync.RWMutex
	AuthKeyID int64
	UserID    int32
	Salts     []*mtproto.TL_future_salt
}
