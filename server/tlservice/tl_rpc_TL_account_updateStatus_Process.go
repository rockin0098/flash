package tlservice

import (
	"time"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_account_updateStatus
func (s *TLService) TL_account_updateStatus_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_account_updateStatus)

	Log.Infof("%T = %+v", tl, tl)

	offline := mtproto.FromBool(tl.Get_offline())
	if offline {
		// statusOffline := &mtproto.TL_userStatusOffline{
		// 	M_was_online: int32(time.Now().Unix()),
		// }
	} else {
		now := time.Now().Unix()
		up := &model.UserPresence{
			UserID:            csess.GetUserID(),
			LastSeenAt:        now,
			LastSeenAuthKeyId: 0,
			LastSeenIp:        "",
		}

		mm := model.GetModelManager()
		err := mm.ModelAdd(up)
		if err != nil {
			Log.Error(err)
			return nil, err
		}
	}

	// to do 状态推送至所有联系人
	// TL_update_short { TL_update_user_status }

	return mtproto.ToBool(true), nil
}
