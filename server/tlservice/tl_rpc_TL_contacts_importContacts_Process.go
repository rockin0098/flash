package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_contacts_importContacts
func (s *TLService) TL_contacts_importContacts_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_account_updateStatus)

	// Log.Infof("TL_account_updateStatus = %+v", FormatStruct(tl))

	return mtproto.ToBool(true), nil
}
