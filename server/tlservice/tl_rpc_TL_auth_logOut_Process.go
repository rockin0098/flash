package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

func (s *TLService) TL_auth_logOut_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := msg.TLObject
	// tl := tlobj.(*mtproto.TL_auth_logOut)

	tltrue := mtproto.ToBool(true)

	return tltrue, nil
}
