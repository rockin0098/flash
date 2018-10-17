package service

import (
	"fmt"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"
)

type TLService struct{}

var tlService = &TLService{}

func TLServiceInstance() *TLService {
	return tlService
}

func (s *TLService) TLMessageProcess(sess *Session, raw *mtproto.RawMessage) error {

	Log.Info("entering...")

	if raw.AuthKeyID == 0 { // 未加密的消息, 握手协商消息
		reqmsg := &mtproto.UnencryptedMessage{}
		err := reqmsg.Decode(raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return err
		}

		res, err := s.TLUnencryptedMessageProcess(sess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return err
		}

		ASSERT(res != nil)
		respmsg := &mtproto.UnencryptedMessage{
			TLObject: res.(mtproto.TLObject),
		}

		err = sess.WriteFull(&mtproto.RawMessage{Payload: respmsg.Encode()})
		if err != nil {
			Log.Error(err)
			return err
		}

		return nil

	} else { // 加密消息

		Log.Debugf("client authKeyID = %v", raw.AuthKeyID)

		authid := raw.AuthKeyID
		mm := model.GetModelManager()
		akey := mm.GetAuthKeyValueByAuthID(authid)
		if akey == nil {
			return fmt.Errorf("authkey not found by authid=%v", authid)
		}

		reqmsg := &mtproto.EncryptedMessage{
			AuthKeyID: authid,
		}
		err := reqmsg.Decode(akey, raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return err
		}

		Log.Infof("Decode encrypted message, authid=%v, msgid=%v, salt=%v, seqno=%v, classType=%v",
			reqmsg.AuthKeyID, reqmsg.MessageID, reqmsg.Salt, reqmsg.SeqNo, reqmsg.TLObject)

		// 记录 client session id
		sess.ClientSessionID = reqmsg.ClientSessionID

		badresp, ok := sess.CheckBadServerSalt(authid, reqmsg.MessageID, reqmsg.SeqNo, reqmsg.Salt)
		if !ok {
			Log.Warnf("check bad server salt. authid = %v, class type = %T", authid, reqmsg.TLObject)
			payload := sess.EncodeMessage(authid, 0, false, badresp)
			return sess.WriteFull(&mtproto.RawMessage{Payload: payload})
		}

		res, err := s.TLEncryptedMessageProcess(sess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return err
		}

		var resPayload []byte = nil
		if res != nil {
			resPayload = sess.EncodeMessage(authid, reqmsg.MessageID, false, res.(mtproto.TLObject))
			return sess.WriteFull(&mtproto.RawMessage{Payload: resPayload})
		}

		return nil
	}

	return nil
}

func (s *TLService) TLUnencryptedMessageProcess(sess *Session, msg *mtproto.UnencryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("class type = %T, \ntlobj = %s", tlobj, tlobj)

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
		Log.Debugf("havent implemented yet, TLType = %T", tl)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tl)
	}

	Log.Debugf("class type = %T, \nresp tlobj = %s", tlobj, res)

	return res, err
}

func (s *TLService) TLEncryptedMessageProcess(sess *Session, msg *mtproto.EncryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("sessid = %v, client sessid = %v, authid = %v, class type = %T, \ntlobj = %v",
		sess.SessionID, sess.ClientSessionID, msg.AuthKeyID, tlobj, tlobj)

	var res interface{}
	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_ping:
		res, err = s.TL_ping_Process(sess, msg)
	case *mtproto.TL_invokeWithLayer:
		res, err = s.TL_invokeWithLayer_Process(sess, msg)
	case *mtproto.TL_initConnection:
		res, err = s.TL_initConnection_Process(sess, msg)
	case *mtproto.TL_help_getConfig:
		res, err = s.TL_help_getConfig_Process(sess, msg)
	case *mtproto.TL_msg_container:
		res, err = s.TL_msg_container_Process(sess, msg)
	case *mtproto.TL_message2:
		res, err = s.TL_message2_Process(sess, msg)
	case *mtproto.TL_auth_logOut:
		res, err = s.TL_auth_logOut_Process(sess, msg)
	default:
		Log.Debugf("havent implemented yet, TLType = %T", tl)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tl)
	}

	Log.Debugf("client sessid = %v, authid = %v, class type = %T, \nresp tlobj = %v",
		sess.SessionID, msg.AuthKeyID, res, res)

	return res, err
}
