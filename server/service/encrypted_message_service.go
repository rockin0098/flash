package service

import (
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/session"
)

func (s *LProtoService) TL_ping_Process(sess *session.Session, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... sessid = %v", sess.SessionID())

	tlobj := msg.TLObject
	tl := tlobj.(*mtproto.TL_ping)

	pong := &mtproto.TL_pong{
		M_msg_id:  msg.MessageID,
		M_ping_id: tl.Get_ping_id(),
	}

	return pong, nil
}
