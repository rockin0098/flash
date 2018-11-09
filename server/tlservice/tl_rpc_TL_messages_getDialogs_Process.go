package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_messages_getDialogs
func (s *TLService) TL_messages_getDialogs_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_messages_getDialogs)

	messageDialogs := &mtproto.TL_messages_dialogs{}

	return messageDialogs, nil
}
