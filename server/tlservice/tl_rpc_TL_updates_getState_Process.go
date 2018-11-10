package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_updates_getState
func (s *TLService) TL_updates_getState_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_updates_getState)

	// Log.Infof("TL_users_getFullUser = %+v", FormatStruct(tl))

	fulluser := &mtproto.TL_userFull{}

	return fulluser, nil
}
