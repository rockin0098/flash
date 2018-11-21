package tlservice

import (
	"time"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

const (
	seqUpdatesNgenId        = "seq_updates_ngen_"
	ptsUpdatesNgenId        = "pts_updates_ngen_"
	qtsUpdatesNgenId        = "qts_updates_ngen_"
	channelPtsUpdatesNgenId = "channel_pts_updates_ngen_"
)

// TL_updates_getState
func (s *TLService) TL_updates_getState_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	pts := int32(1)

	userid := csess.GetUserID()
	upts := s.Dao.UserPtsUpdatesDao.GetUserPtsUpdatesByID(userid)
	if upts != nil {
		pts = upts.Pts
	}

	state := &mtproto.TL_updates_state{
		M_pts:          pts,
		M_qts:          0,
		M_seq:          -1,
		M_date:         int32(time.Now().Unix()),
		M_unread_count: 0,
	}

	return state, nil
}
