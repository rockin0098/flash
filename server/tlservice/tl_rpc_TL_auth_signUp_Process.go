package tlservice

import (
	"math/rand"
	"time"

	"github.com/rockin0098/meow/server/model"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_signUp
func (s *TLService) TL_auth_signUp_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_signUp)

	Log.Infof("%T = %+v", tl, FormatStruct(tl))

	phone := tl.Get_phone_number()

	var tluser mtproto.TLObject
	mm := model.GetModelManager()
	user := mm.GetUserByPhoneNumber(phone)
	if user != nil {
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_NUMBER_OCCUPIED)
	} else {
		user = &model.User{
			AccessHash:  rand.Int63(),
			Phone:       phone,
			FirstName:   tl.Get_first_name(),
			LastName:    tl.Get_last_name(),
			CountryCode: "UN",
		}

		err := mm.CreateNewUser(user, csess.AuthKeyID)
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		// set csess userid
		csess.UserID = user.ID
	}

	userStatus := &mtproto.TL_userStatusOnline{
		M_expires: int32(time.Now().Unix()) + 1800,
	}

	tluser = &mtproto.TL_user{
		M_self:           mtproto.ToBool(true),
		M_contact:        mtproto.ToBool(true),
		M_mutual_contact: mtproto.ToBool(true),
		M_id:             int32(user.ID),
		M_access_hash:    user.AccessHash,
		M_first_name:     user.FirstName,
		M_last_name:      user.LastName,
		M_username:       user.UserName,
		M_phone:          user.Phone,
		M_status:         userStatus,
	}

	authAuthorization := &mtproto.TL_auth_authorization{
		M_user: tluser,
	}

	return authAuthorization, nil
}
