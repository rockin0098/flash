package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_resendCode
func (s *TLService) TL_auth_resendCode_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_resendCode)

	phone := tl.Get_phone_number()

	registered := s.Dao.UserDao.CheckPhoneExists(phone)

	Log.Infof("registered = %v", registered)

	codeType := &mtproto.TL_auth_sentCodeTypeApp{
		M_length: 5,
	}
	authSentCode := &mtproto.TL_auth_sentCode{
		// M_phone_registered: mtproto.ToBool(registered),
		M_type:            codeType,
		M_phone_code_hash: tl.Get_phone_code_hash(),
		// M_next_type:       codeType,
		M_timeout: 60,
	}

	if registered {
		authSentCode.M_phone_registered = mtproto.ToBool(registered)
	}

	return authSentCode, nil
}
