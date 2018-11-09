package tlservice

import (
	"time"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_contacts_getContacts
func (s *TLService) TL_contacts_getContacts_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_contacts_getContacts)

	userid := csess.GetUserID()

	mm := model.GetModelManager()
	contactList := mm.GetContactsByUserID(userid)

	var contacts mtproto.TLObject
	if len(contactList) > 0 {

		var clist []mtproto.TLObject
		var ids []int64
		for _, c := range contactList {
			ids = append(ids, c.ContactUserID)
			c2 := &mtproto.TL_contact{
				M_user_id: int32(c.ContactUserID),
				M_mutual:  mtproto.ToBool(c.Mutual == 1),
			}
			clist = append(clist, c2)
		}
		ids = append(ids, userid) // 加上自己
		users := mm.GetUsersByIDList(ids)
		var tlusers []mtproto.TLObject
		for _, u := range users {
			status := &mtproto.TL_userStatusOnline{
				M_expires: int32(time.Now().Unix()) + 1800,
			}
			tlu := User_to_TL_user(u, status)
			tlusers = append(tlusers, tlu)
		}

		contacts = &mtproto.TL_contacts_contacts{
			M_contacts:    clist,
			M_saved_count: 0,
			M_users:       tlusers,
		}
	} else {
		contacts = mtproto.New_TL_contacts_contacts()
	}

	return contacts, nil
}
