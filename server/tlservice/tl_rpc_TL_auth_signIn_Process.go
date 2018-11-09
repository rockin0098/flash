package tlservice

import (
	"time"

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

	if tl.Get_phone_code() == "" {

	}

	userStatus := &mtproto.TL_userStatusOnline{
		M_expires: int32(time.Now().Unix()) + 1800,
	}

	user := &mtproto.TL_user{
		M_self:           mtproto.ToBool(true),
		M_contact:        mtproto.ToBool(true),
		M_mutual_contact: mtproto.ToBool(true),
		M_id:             2,
		M_access_hash:    66666666666666,
		M_first_name:     "qqq",
		M_last_name:      "aaa",
		M_username:       "dadada",
		M_phone:          "+8613333333333",
		M_status:         userStatus,
	}

	authAuthorization := &mtproto.TL_auth_authorization{
		M_user: user,
	}

	return authAuthorization, nil
}
