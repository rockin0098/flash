package tlservice

import (
	"time"

	"github.com/rockin0098/meow/server/model"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_signIn
func (s *TLService) TL_auth_signIn_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_signIn)

	Log.Infof("TL_auth_signIn = %+v", FormatStruct(tl))

	phone := tl.Get_phone_number()

	if tl.Get_phone_code() == "" {
		Log.Warn("phone code is empty")
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EMPTY), "code empty")
	}

	hash := tl.Get_phone_code_hash()
	mm := model.GetModelManager()
	pt := mm.GetAuthPhoneTransactionByHash(hash)
	if pt == nil || pt.Code != tl.Get_phone_code() || pt.PhoneNumber != phone {
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
	}

	var tluser mtproto.TLObject
	user := mm.GetUserByPhoneNumber(phone)
	if user == nil {
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_NUMBER_UNOCCUPIED)
	} else {

		// set csess userid
		csess.UserID = user.ID

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
	}

	authAuthorization := &mtproto.TL_auth_authorization{
		M_user: tluser,
	}

	return authAuthorization, nil
}
