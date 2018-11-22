package tlservice

import (
	"time"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_contacts_importContacts
func (s *TLService) TL_contacts_importContacts_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_contacts_importContacts)

	Log.Infof("%T = %+v", tl, tl)

	newContacts := tl.Get_contacts()
	if len(newContacts) == 0 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		Log.Warn(err, ": contacts empty")
		return nil, err
	}

	var importedContacts mtproto.TLObject
	var imported []mtproto.TLObject
	var users []mtproto.TLObject

	for _, nc := range newContacts {

		tlnc := nc.(*mtproto.TL_inputPhoneContact)
		phone := tlnc.Get_phone()
		u := s.Dao.UserDao.GetUserByPhoneNumber(phone)
		if u == nil {
			Log.Warnf("phone have not registered yet, phone = %v", phone)
			continue
		}

		myid := csess.GetUserID()
		mys := s.Dao.UserDao.GetUsersByIDList([]int32{myid})
		ASSERT(len(mys) == 1)
		my := mys[0]

		if !s.Dao.UserContactDao.IsMyContact(myid, u.ID) {
			now := int32(time.Now().Unix())
			mycontact := &model.UserContact{
				OwnerUserID:      myid,
				ContactUserID:    u.ID,
				ContactPhone:     u.Phone,
				ContactFirstName: u.FirstName,
				ContactLastName:  u.LastName,
				Mutual:           1,
				Date2:            now,
			}

			ucontact := &model.UserContact{
				OwnerUserID:      u.ID,
				ContactUserID:    myid,
				ContactPhone:     my.Phone,
				ContactFirstName: my.FirstName,
				ContactLastName:  my.LastName,
				Mutual:           1,
				Date2:            now,
			}

			err := s.Dao.ModelAdd(mycontact)
			if err != nil {
				Log.Error(err)
				return nil, err
			}
			err = s.Dao.ModelAdd(ucontact)
			if err != nil {
				Log.Error(err)
				return nil, err
			}

			imp := &mtproto.TL_importedContact{
				M_user_id:   int32(u.ID),
				M_client_id: tlnc.M_client_id,
			}

			imported = append(imported, imp)

			status := &mtproto.TL_userStatusOnline{
				M_expires: int32(time.Now().Unix()) + 1800,
			}
			tluser := model.User_to_TL_user(u, status)
			users = append(users, tluser)
		}
	}

	importedContacts = &mtproto.TL_contacts_importedContacts{
		M_imported: imported,
		M_users:    users,
	}

	return importedContacts, nil
}
