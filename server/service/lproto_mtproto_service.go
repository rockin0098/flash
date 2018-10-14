package service

import (
	"errors"
	"fmt"

	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"
	"github.com/rockin0098/flash/server/session"

	. "github.com/rockin0098/flash/base/logger"
)

// . "github.com/rockin0098/flash/base/global"

func (s *LProtoService) MTProtoMessageProcess(sess *session.Session, raw *mtproto.RawMessage) (interface{}, error) {

	Log.Info("entering...")

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

		return resPayload, nil

	} else { // 加密消息

		Log.Debugf("client authKeyID = %v", raw.AuthKeyID)

		authid := raw.AuthKeyID
		mm := model.GetModelManager()
		akey := mm.GetAuthKeyValueByAuthID(authid)
		if akey == nil {
			return nil, fmt.Errorf("authkey not found by authid=%v", authid)
		}

		reqmsg := &mtproto.EncryptedMessage{
			AuthKeyID: authid,
		}
		err := reqmsg.Decode(akey, raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return nil, err
		}

		Log.Infof("Decode encrypted message, authid=%v, msgid=%v, salt=%v, seqno=%v, classType=%v", reqmsg.AuthKeyID, reqmsg.MessageID, reqmsg.Salt, reqmsg.SeqNo, reqmsg.TLObject)

		// 记录 client session id
		csm := session.GetClientSessionManager()
		cltSess, ok := csm.Load(reqmsg.ClientSessionID)
		if !ok {
			cltSess = session.NewClientSession(reqmsg.ClientSessionID, authid, reqmsg.Salt, reqmsg.MessageID, sess.SessionID())
			csm.Store(reqmsg.ClientSessionID, cltSess)
		}

		badresp, ok := cltSess.CheckBadServerSalt(authid, reqmsg.MessageID, reqmsg.SeqNo, reqmsg.Salt)
		if !ok {
			Log.Warnf("check bad server salt. authid = %v, class type = %T", authid, reqmsg.TLObject)
			payload := cltSess.EncodeMessage(authid, akey, 0, false, badresp)
			return payload, errors.New("check bad server salt failed")
		}

		res, err := s.MTProtoEncryptedMessageProcess(cltSess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return nil, err
		}

		var resPayload []byte = nil
		if res != nil {
			resPayload = cltSess.EncodeMessage(authid, akey, reqmsg.MessageID, false, res.(mtproto.TLObject))
		}

		return resPayload, nil
	}

	return nil, nil
}

func (s *LProtoService) MTProtoUnencryptedMessageProcess(sess *session.Session, msg *mtproto.UnencryptedMessage) (interface{}, error) {

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
		Log.Debugf("havent implemented yet, type = %T", tl)
	}

	Log.Debugf("class type = %T, \nresp tlobj = %s", tlobj, res)

	return res, err
}

func (s *LProtoService) MTProtoEncryptedMessageProcess(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("client sessid = %v, authid = %v, class type = %T, \ntlobj = %v",
		cltSess.SessionID(), cltSess.AuthKeyID(), tlobj, tlobj)

	var res interface{}
	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_ping:
		res, err = s.TL_ping_Process(cltSess, msg)
	case *mtproto.TL_invokeWithLayer:
		res, err = s.TL_invokeWithLayer_Process(cltSess, msg)
	case *mtproto.TL_initConnection:
		res, err = s.TL_initConnection_Process(cltSess, msg)
	case *mtproto.TL_help_getConfig:
		res, err = s.TL_help_getConfig_Process(cltSess, msg)
	case *mtproto.TL_msg_container:
		res, err = s.TL_msg_container_Process(cltSess, msg)
	case *mtproto.TL_message2:
		res, err = s.TL_message2_Process(cltSess, msg)
	case *mtproto.TL_auth_logOut:
		res, err = s.TL_auth_logOut_Process(cltSess, msg)
	default:
		Log.Debugf("havent implemented yet, type = %T", tl)
	}

	Log.Debugf("client sessid = %v, authid = %v, class type = %T, \nresp tlobj = %v",
		cltSess.SessionID(), cltSess.AuthKeyID(), res, res)

	return res, err
}
