package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_help_getNearestDc
func (s *TLService) TL_help_getNearestDc_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	dc := &mtproto.TL_nearestDc{
		M_country:    "US",
		M_this_dc:    2,
		M_nearest_dc: 2,
	}

	return dc, nil
}
