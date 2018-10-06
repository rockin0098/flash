package session

import (
	"sync"

	"github.com/rockin0098/flash/base/guid"
	"github.com/rockin0098/flash/proto/mtproto"
)

func GenerateSessionID() string {
	return guid.GenerateStringUID()
}

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
var sessionStore = &sync.Map{}

type Session struct {
	rw         *sync.RWMutex
	sessionID  string
	connID     string
	mtpCryptor *mtproto.MTProtoCryptor
	mtpState   *mtproto.MTProtoState
}

func NewSession(connID string) *Session {

	sess := &Session{
		rw:         &sync.RWMutex{},
		sessionID:  GenerateSessionID(),
		connID:     connID,
		mtpCryptor: mtproto.NewMTProtoCryptor(),
		mtpState:   mtproto.NewMTProtoState(),
	}

	sessionStore.Store(sess.sessionID, sess)

	return sess
}

func GetSession(sessid string) *Session {
	sobj, ok := sessionStore.Load(sessid)
	if !ok {
		return nil
	}

	return sobj.(*Session)
}

func (s *Session) SessionID() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.sessionID
}

func (s *Session) ConnectionID() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.connID
}

func (s *Session) MTProtoCryptor() *mtproto.MTProtoCryptor {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.mtpCryptor
}

func (s *Session) MTProtoState() *mtproto.MTProtoState {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.mtpState
}

func (s *Session) SetMTProtoState(state *mtproto.MTProtoState) {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.mtpState = state
}
