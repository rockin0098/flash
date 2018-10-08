package service

import (
	"encoding/hex"
	"fmt"

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

		// mtp := sess.MTProto()
		// authID := mtp.State().AuthKeyID
		// authKey := mtp.State().AuthKey

		// Log.Debugf("client authKeyID = %v, authid = %v, authkey = %v", raw.AuthKeyID, authID, authKey)
		Log.Debugf("client authKeyID = %v", raw.AuthKeyID)

		ms := ModelServiceInstance()
		authKey := ms.GetAuthKeyByAuthID(raw.AuthKeyID)
		if authKey == nil {
			return nil, fmt.Errorf("authkey not found by authid=%v", raw.AuthKeyID)
		}
		akey, err := hex.DecodeString(authKey.Body)
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		reqmsg := &mtproto.EncryptedMessage{
			AuthKeyID: raw.AuthKeyID,
		}
		err = reqmsg.Decode(akey, raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		// 记录 client session id
		sm := session.GetSessionManager()
		_, ok := sm.LoadClient(reqmsg.ClientSessionID)
		if !ok {
			sm.StoreClient(reqmsg.ClientSessionID, sess)
		}

		res, err := s.MTProtoEncryptedMessageProcess(sess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return nil, err
		}

		respmsg := &mtproto.EncryptedMessage{
			TLObject: res.(mtproto.TLObject),
		}

		resPayload := respmsg.Encode(raw.AuthKeyID, akey)
		Log.Debugf("sessid = %v, resp payload = %v", sess.SessionID(), hex.EncodeToString(resPayload))

		return resPayload, nil
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
		res, err = s.TL_req_pq_Process(sess, msg)
	case *mtproto.TL_req_DH_params:
		res, err = s.TL_req_DH_params_Process(sess, msg)
	case *mtproto.TL_set_client_DH_params:
		res, err = s.TL_set_client_DH_params_Process(sess, msg)
	default:
		Log.Debugf("havent implemented yet, type = %T", tl)
	}

	return res, err
}

func (s *LProtoService) MTProtoEncryptedMessageProcess(sess *session.Session, msg *mtproto.EncryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("class type = %T", tlobj)

	var res interface{}
	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_ping:
		res, err = s.TL_ping_Process(sess, msg)
	default:
		Log.Debugf("havent implemented yet, type = %T", tl)
	}

	return res, err
}

// verify encrypted messages
func (s *LProtoService) verify1() {

	// if !sess.CheckBadServerSalt(sessionMsg.connID, sessionMsg.md, message.MessageId, message.SeqNo, message.Salt) {
	// 	glog.Infof("salt invalid - {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.md)
	// 	// glog.Error("salt invalid..")
	// 	return
	// }

	// _, isContainer := message.Object.(*mtproto.TLMsgContainer)
	// if !sess.CheckBadMsgNotification(sessionMsg.connID, sessionMsg.md, message.MessageId, message.SeqNo, isContainer) {
	// 	glog.Infof("bad msg invalid - {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.md)
	// 	// glog.Error("bad msg invalid..")
	// 	return
	// }

	// /*
	// 	//=============================================================================================
	// 	// Check Message Sequence Number (msg_seqno)
	// 	//
	// 	// https://core.telegram.org/mtproto/description#message-sequence-number-msg-seqno
	// 	// Message Sequence Number (msg_seqno)
	// 	//
	// 	// A 32-bit number equal to twice the number of “content-related” messages
	// 	// (those requiring acknowledgment, and in particular those that are not containers)
	// 	// created by the sender prior to this message and subsequently incremented
	// 	// by one if the current message is a content-related message.
	// 	// A container is always generated after its entire contents; therefore,
	// 	// its sequence number is greater than or equal to the sequence numbers of the messages contained in it.
	// 	//

	// 	if message.SeqNo < sess.lastSeqNo {
	// 		err = fmt.Errorf("sequence number is greater than or equal to the sequence numbers of the messages contained in it: %d", message.SeqNo)
	// 		glog.Error(err)

	// 		// TODO(@benqi): ignore this message or close client conn??
	// 		return
	// 	}
	// 	sess.lastSeqNo = message.SeqNo

	// 	sess.onMessageData(sessionMsg.md, message.MessageId, message.SeqNo, message.Object)
	// */

	// var messages = &messageListWrapper{[]*mtproto.TLMessage2{}}
	// extractClientMessage(message.MessageId, message.SeqNo, message.Object, messages, func(layer int32) {
	// 	s.Layer = layer

	// 	// TODO(@benqi): clear session_manager
	// })
}

func (s *LProtoService) verify2() {

}
