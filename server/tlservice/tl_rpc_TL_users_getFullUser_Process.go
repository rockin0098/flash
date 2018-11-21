package tlservice

import (
	"fmt"
	"time"

	"github.com/rockin0098/meow/server/model"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_users_getFullUser
func (s *TLService) TL_users_getFullUser_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_users_getFullUser)

	id := tl.Get_id()

	var userid int64
	// var accesshash int64

	switch id.ClassID() {
	case mtproto.TL_CLASS_inputUserSelf:
		userid = csess.GetUserID()
	case mtproto.TL_CLASS_inputUser:
		userid = int64(id.(*mtproto.TL_inputUser).Get_user_id())
		// accesshash = id.(*mtproto.TL_inputUser).Get_access_hash()
	default:
		panic(fmt.Sprintf("class type = %T", id))
	}

	Log.Infof("userid = %v", userid)

	fulluser := mtproto.New_TL_userFull()
	fulluser.Set_phone_calls_available(mtproto.ToBool(true))
	fulluser.Set_phone_calls_private(mtproto.ToBool(false))
	fulluser.Set_about("")
	fulluser.Set_common_chats_count(0)

	user := s.Dao.UserDao.GetUserByID(userid)
	state := mtproto.New_TL_userStatusOnline()
	tluser := model.User_to_TL_user(user, state)

	link := &mtproto.TL_contacts_link{
		M_my_link:      mtproto.New_TL_contactLinkContact(),
		M_foreign_link: mtproto.New_TL_contactLinkContact(),
		M_user:         tluser,
	}

	fulluser.Set_link(link)
	fulluser.Set_user(tluser)
	fulluser.Set_notify_settings(
		&mtproto.TL_peerNotifySettings{
			M_mute_until: 1,
			M_sound:      "default",
		},
	)

	sizeData := &mtproto.TL_photoSize{
		M_type: "0",
		M_w:    100,
		M_h:    100,
		M_size: 1024,
		M_location: &mtproto.TL_fileLocation{
			M_volume_id: 0,
			M_local_id:  0,
			M_secret:    0,
			M_dc_id:     2,
		},
	}

	photo := &mtproto.TL_photo{
		M_id:           0,
		M_has_stickers: mtproto.ToBool(false),
		M_access_hash:  0,
		M_date:         int32(time.Now().Unix()),
		M_sizes:        []mtproto.TLObject{sizeData},
	}

	fulluser.Set_profile_photo(photo)

	return fulluser, nil
}
