package tlservice

import (
	"github.com/rockin0098/meow/base/crypto"
	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_sendCode
func (s *TLService) TL_auth_sendCode_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_sendCode)

	Log.Infof("TL_auth_sendCode = %+v", FormatStruct(tl))

	phone := tl.Get_phone_number()

	mm := model.GetModelManager()
	registered := mm.CheckPhoneExists(phone)

	Log.Infof("registered = %v", registered)

	codeType := &mtproto.TL_auth_sentCodeTypeApp{
		M_length: 5,
	}
	authSentCode := &mtproto.TL_auth_sentCode{
		// M_phone_registered: mtproto.ToBool(registered),
		M_type:            codeType,
		M_phone_code_hash: crypto.GenerateStringNonce(16),
		// M_next_type:       codeType,
		M_timeout: 60,
	}

	if registered {
		authSentCode.M_phone_registered = mtproto.ToBool(registered)
	}

	return authSentCode, nil
}
