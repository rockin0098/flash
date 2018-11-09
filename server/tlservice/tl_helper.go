package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
)

func User_to_TL_user(user *model.User, status mtproto.TLObject) *mtproto.TL_user {
	tluser := &mtproto.TL_user{
		M_self:           mtproto.ToBool(true),
		M_contact:        mtproto.ToBool(true),
		M_mutual_contact: mtproto.ToBool(true),
		M_id:             int32(user.ID),
		M_access_hash:    user.AccessHash,
		M_first_name:     user.FirstName,
		M_last_name:      user.LastName,
		M_username:       user.UserName,
		M_phone:          user.Phone,
		M_status:         status,
	}

	return tluser
}
