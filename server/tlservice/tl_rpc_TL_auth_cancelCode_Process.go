package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_cancelCode
func (s *TLService) TL_auth_cancelCode_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_auth_cancelCode)

	// phone := tl.Get_phone_number()
	// hash := tl.Get_phone_code_hash()

	// delete code record

	return mtproto.ToBool(true), nil
}
