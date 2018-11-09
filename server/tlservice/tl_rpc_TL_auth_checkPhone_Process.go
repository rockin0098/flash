package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_auth_checkPhone
func (s *TLService) TL_auth_checkPhone_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_checkPhone)

	// to do : check phone number format
	phone := tl.Get_phone_number()

	mm := model.GetModelManager()
	registered := mm.CheckPhoneExists(phone)

	checkedPhone := &mtproto.TL_auth_checkedPhone{
		M_phone_registered: mtproto.ToBool(registered),
	}

	return checkedPhone, nil
}
