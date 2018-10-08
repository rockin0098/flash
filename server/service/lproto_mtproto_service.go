package service

import (
	"encoding/hex"

	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/session"

	. "github.com/rockin0098/flash/base/logger"
)

// . "github.com/rockin0098/flash/base/global"

func (s *LProtoService) MTProtoMessageProcess(sess *session.Session, raw *mtproto.RawMessage) (interface{}, error) {

	Log.Infof("entering.......sessid=%v, TransportType = %v, AuthKeyID = %v, QuickAckID = %v, Payload = \n%v\n",
		sess.SessionID(), raw.TransportType, raw.AuthKeyID, raw.QuickAckID, hex.EncodeToString(raw.Payload))

	if raw.AuthKeyID == 0 { // 未加密的消息, 握手协商消息
		reqmsg := &mtproto.UnencryptedMessage{}
		err := reqmsg.Decode(raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		res, err := s.MTProtoUnencryptedMessageProcess(sess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return nil, err
		}

		respmsg := &mtproto.UnencryptedMessage{
			TLObject: res.(mtproto.TLObject),
		}

		resPayload := respmsg.Encode()

		Log.Debugf("sessid = %v, resp payload = %v", sess.SessionID(), hex.EncodeToString(resPayload))

		return resPayload, nil

	} else { // 加密消息

		mtp := sess.MTProto()
		authID := mtp.State().AuthKeyID
		authKey := mtp.State().AuthKey

		Log.Debugf("client authKeyID = %v, authid = %v, authkey = %v", raw.AuthKeyID, authID, authKey)

		encryptedMessage := &mtproto.EncryptedMessage{
			AuthKeyID: raw.AuthKeyID,
		}
		err := encryptedMessage.Decode(authKey, raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return nil, err
		}

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
	case *mtproto.TL_req_DH_params:
		res, err = s.TL_req_DH_params_Process(sess, tl)
	case *mtproto.TL_set_client_DH_params:
		res, err = s.TL_set_client_DH_params_Process(sess, tl)
	default:
		Log.Debugf("havent implemented yet, type = %v", mtproto.TL_CLASS_NAME[tl.ClassID()])
	}

	return res, err
}

func (s *LProtoService) MTProtoEncryptedMessageProcess(sess *session.Session, msg *mtproto.EncryptedMessage) (interface{}, error) {

	return nil, nil
}
