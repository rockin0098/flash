package service

import (
	"encoding/hex"

	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/session"

	. "github.com/rockin0098/flash/base/logger"
)

// . "github.com/rockin0098/flash/base/global"

func (s *LProtoService) MTProtoMessageProcess(sess *session.Session, raw *mtproto.RawMessage) (interface{}, error) {

	Log.Infof("entering....... TransportType = %v, AuthKeyID = %v, QuickAckID = %v, Payload = \n%v\n",
		raw.TransportType, raw.AuthKeyID, raw.QuickAckID, hex.EncodeToString(raw.Payload))

	if raw.AuthKeyID == 0 { // 未加密的消息, 握手协商消息
		unencryptedMessage := &mtproto.UnencryptedMessage{}
		err := unencryptedMessage.Decode(raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		return s.MTProtoUnencryptedMessageProcess(sess, unencryptedMessage)

	} else { // 加密消息

		encryptedMessage := &mtproto.EncryptedMessage{}
		// err := encryptedMessage.Decode()

		return s.MTProtoEncryptedMessageProcess(sess, encryptedMessage)
	}

	return nil, nil
}

func (s *LProtoService) MTProtoUnencryptedMessageProcess(sess *session.Session, msg *mtproto.UnencryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("class type = %T", tlobj)

	var res interface{}
	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_req_pq:
		res, err = s.TL_req_pq_Process(sess, tl)
	default:
		Log.Debugf("havent implemented yet, type = %v", mtproto.TL_CLASS_NAME[tl.ClassID()])
	}

	return res, err
}

func (s *LProtoService) MTProtoEncryptedMessageProcess(sess *session.Session, msg *mtproto.EncryptedMessage) (interface{}, error) {

	return nil, nil
}
