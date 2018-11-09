package service

import (
	"container/list"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
)

// type TLMessageReplyTypeEnum int

// const (
// 	_ TLMessageReplyTypeEnum = iota
// 	TLMessageReplyNewMessageID
// 	TLMessageReplyZeroMessageID
// )

type MessageListWrapper struct {
	Messages []*mtproto.TL_message2
}

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

		err = sess.WriteFull2(&mtproto.RawMessage{Payload: respmsg.Encode()})
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

		clientSessionID := reqmsg.ClientSessionID
		ss := SessionServiceInstance()
		csess, ok := ss.LoadClientSession(clientSessionID)
		if !ok {
			// 记录  session id
			csess = &ClientSession{
				Sess:             sess,
				ClientSessionID:  clientSessionID,
				AuthKeyID:        authid,
				Salt:             reqmsg.Salt,
				FirstMessageID:   reqmsg.MessageID,
				CloseDate:        time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout,
				CloseSessionDate: 0,
				ApiMessages:      list.New(),
				UniqueID:         rand.Int63(),
				ClientState:      kStateCreated,
				PendingMessages:  []*PendingMessage{},
				IsUpdates:        false,
				RpcMessages:      []*NetworkApiMessage{},
			}

			ss.NewClientSession(clientSessionID, csess)

		} else {
			csess.Sess = sess
		}

		if !csess.CheckBadServerSalt(reqmsg.MessageID, reqmsg.SeqNo, reqmsg.Salt) {
			Log.Warnf("check bad server salt. authid = %v, class type = %T", authid, reqmsg.TLObject)

			if !ok || reqmsg.MessageID < csess.FirstMessageID {
				Log.Info("reply to client making new client session")

				newSessionCreated := &mtproto.TL_new_session_created{
					M_first_msg_id: reqmsg.MessageID,
					M_unique_id:    csess.UniqueID,
					M_server_salt:  csess.Salt,
				}

				csess.WriteFull(mtproto.GenerateMessageID(), true, newSessionCreated)
			}

			return nil
		}

		_, isContainer := reqmsg.TLObject.(*mtproto.TL_msg_container)
		if !csess.CheckBadMsgNotification(reqmsg.MessageID, reqmsg.SeqNo, isContainer) {
			Log.Warnf("check bad msg notification. authid = %v, class type = %T", authid, reqmsg.TLObject)
			return nil
		}

		var messages = &MessageListWrapper{[]*mtproto.TL_message2{}}
		s.TLExtractMessageProcess(csess, reqmsg.MessageID, reqmsg.SeqNo, reqmsg.TLObject, messages)
		Log.Infof("message wrapper = %+v", messages)

		err = s.TLMessageListWrapperProcess(csess, messages)
		if err != nil {
			Log.Warn(err)
			return err
		}

		return nil
	}

	return nil
}

func (s *TLService) TLExtractMessageProcess(csess *ClientSession, msgid int64, seqNo int32, object mtproto.TLObject, messages *MessageListWrapper) {

	switch object.(type) {
	case *mtproto.TL_msg_container:
		msgContainer, _ := object.(*mtproto.TL_msg_container)

		// A container does not require an acknowledgment
		if seqNo%2 != 0 {
			Log.Error("A container does not require an acknowledgment.")
			return
		}

		for _, m := range msgContainer.M_message2s {
			if m.M_body == nil {
				continue
			}

			// Check msgId
			//
			// A container is always generated after its entire contents; therefore,
			// its sequence number is greater than or equal to the sequence numbers of the messages contained in it.
			//
			if m.M_seqno > seqNo {
				Log.Errorf("sequence number is greater than or equal to the sequence numbers of the messages contained in it: %d", seqNo)
				// TODO(@benqi): close client and add to banned??
				continue
			}

			// may not carry other simple containers
			if _, ok := m.M_body.(*mtproto.TL_msg_container); ok {
				Log.Error("may not carry other simple containers")
				// TODO(@benqi): close client and add to banned??
				continue
			}

			s.TLExtractMessageProcess(csess, m.M_msg_id, m.M_seqno, m.M_body, messages)
		}
	case *mtproto.TL_gzip_packed:
		gzipPacked, _ := object.(*mtproto.TL_gzip_packed)
		Log.Info("processGzipPacked - request data: ", gzipPacked)

		dc := mtproto.NewMTPDecodeBuffer(gzipPacked.M_packed_data)
		o := dc.TLObject()
		if o == nil {
			Log.Errorf("Decode query error: %s", hex.EncodeToString(gzipPacked.M_packed_data))
			return
		}
		s.TLExtractMessageProcess(csess, msgid, seqNo, o, messages)
		// case *mtproto.TL_msg_copy:
		// 	// not use in client
		// 	Log.Error("android client not use msg_copy: ", object)
	case *mtproto.TL_invokeWithLayer:
		Log.Infof("TL_invokeWithLayer...  client sessid = %v", csess.ClientSessionID)

		tlobj := object
		tl := tlobj.(*mtproto.TL_invokeWithLayer)

		layer := tl.Get_layer()
		query := tl.Get_query()
		csess.Layer = layer

		Log.Debugf("invokeWithLayer layer = %v, query = %T, \nquery = %s", layer, query, query)

		// must be initConnection
		initConn := query.(*mtproto.TL_initConnection)

		messages.Messages = append(messages.Messages, &mtproto.TL_message2{
			M_msg_id: msgid,
			M_seqno:  seqNo,
			M_body:   initConn,
		})
	case *mtproto.TL_invokeAfterMsg:
		// Log.Error("not implemented %T", object)
		messages.Messages = append(messages.Messages, &mtproto.TL_message2{
			M_msg_id: msgid,
			M_seqno:  seqNo,
			M_body:   object,
		})
	case *mtproto.TL_invokeAfterMsgs:
		// Log.Error("not implemented %T", object)
		messages.Messages = append(messages.Messages, &mtproto.TL_message2{
			M_msg_id: msgid,
			M_seqno:  seqNo,
			M_body:   object,
		})
	case *mtproto.TL_invokeWithoutUpdates:
		// Log.Error("not implemented %T", object)
		messages.Messages = append(messages.Messages, &mtproto.TL_message2{
			M_msg_id: msgid,
			M_seqno:  seqNo,
			M_body:   object,
		})
	default:
		messages.Messages = append(messages.Messages, &mtproto.TL_message2{
			M_msg_id: msgid,
			M_seqno:  seqNo,
			M_body:   object,
		})
	}

	return
}

func (s *TLService) TLMessageListWrapperProcess(csess *ClientSession, msglist *MessageListWrapper) error {
	var (
		hasRpcRequest bool
		hasHttpWait   bool
		ok            bool
	)

	messages := msglist.Messages
	c := csess

	for _, message := range messages {
		Log.Infof("message - %T", message)

		if message.M_body == nil {
			continue
		}

		var err error

		switch message.M_body.(type) {
		case *mtproto.TL_rpc_drop_answer: // 所有链接都有可能
			rpcDropAnswer, _ := message.M_body.(*mtproto.TL_rpc_drop_answer)
			err = s.TL_rpc_drop_answer_Process(csess, message.M_msg_id, message.M_seqno, rpcDropAnswer)
		case *mtproto.TL_get_future_salts: // GENERIC
			getFutureSalts, _ := message.M_body.(*mtproto.TL_get_future_salts)
			err = s.TL_get_future_salts_Process(csess, message.M_msg_id, message.M_seqno, getFutureSalts)
		case *mtproto.TL_http_wait: // 未知
			// c.onHttpWait(connID, md, message.MsgId, message.Seqno, message.Object)
			hasHttpWait = true
			c.IsUpdates = true
		case *mtproto.TL_ping: // android未用
			ping, _ := message.M_body.(*mtproto.TL_ping)
			err = s.TL_ping_Process(csess, message.M_msg_id, message.M_seqno, ping)
		case *mtproto.TL_ping_delay_disconnect: // PUSH和GENERIC
			ping, _ := message.M_body.(*mtproto.TL_ping_delay_disconnect)
			err = s.TL_ping_delay_disconnect_Process(csess, message.M_msg_id, message.M_seqno, ping)
		case *mtproto.TL_destroy_session: // GENERIC
			destroySession, _ := message.M_body.(*mtproto.TL_destroy_session)
			err = s.TL_destroy_session_Process(csess, message.M_msg_id, message.M_seqno, destroySession)
		case *mtproto.TL_destroy_auth_key:
			destroyAuthKey, _ := message.M_body.(*mtproto.TL_destroy_auth_key)
			err = s.TL_destroy_auth_key_Process(csess, message.M_msg_id, message.M_seqno, destroyAuthKey)
		case *mtproto.TL_msgs_ack: // 所有链接都有可能
			msgsAck, _ := message.M_body.(*mtproto.TL_msgs_ack)
			err = s.TL_msgs_ack_Process(csess, message.M_msg_id, message.M_seqno, msgsAck)
			// TODO(@benqi): check c.isUpdates
		case *mtproto.TL_msgs_state_req: // android未用
			err = s.TL_msgs_state_req_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_msgs_state_info: // android未用
			err = s.TL_msgs_state_info_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_msgs_all_info: // android未用
			err = s.TL_msgs_all_info_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_msg_resend_req: // 都有可能
			err = s.TL_msg_resend_req_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_msg_detailed_info: // 都有可能
			err = s.TL_msg_detailed_info_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_msg_new_detailed_info: // 都有可能
			err = s.TL_msg_new_detailed_info_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_contest_saveDeveloperInfo: // 未知
			err = s.TL_contest_saveDeveloperInfo_Process(csess, message.M_msg_id, message.M_seqno, message.M_body)
		case *mtproto.TL_invokeAfterMsg:
			invokeAfterMsg, _ := message.M_body.(*mtproto.TL_invokeAfterMsg)
			err = s.TLRpcMessageProcess(csess, message.M_msg_id, message.M_seqno, invokeAfterMsg.Get_query())
			if err == nil && !hasRpcRequest {
				hasRpcRequest = ok
			}
		case *mtproto.TL_invokeAfterMsgs: // 未知
			invokeAfterMsgs, _ := message.M_body.(*mtproto.TL_invokeAfterMsg)
			err = s.TLRpcMessageProcess(csess, message.M_msg_id, message.M_seqno, invokeAfterMsgs.Get_query())
			if err == nil && !hasRpcRequest {
				hasRpcRequest = ok
			}
		case *mtproto.TL_initConnection:
			initConnection, _ := message.M_body.(*mtproto.TL_initConnection)
			err = s.TLRpcMessageProcess(csess, message.M_msg_id, message.M_seqno, initConnection.Get_query())
			if err == nil && !hasRpcRequest {
				hasRpcRequest = ok
			}
		case *mtproto.TL_invokeWithoutUpdates:
			invokeWithoutUpdates, _ := message.M_body.(*mtproto.TL_invokeWithoutUpdates)
			err = s.TLRpcMessageProcess(csess, message.M_msg_id, message.M_seqno, invokeWithoutUpdates.Get_query())
			if err == nil && !hasRpcRequest {
				hasRpcRequest = ok
			}
		default:
			err = s.TLRpcMessageProcess(csess, message.M_msg_id, message.M_seqno, message.M_body)
			if err == nil && !hasRpcRequest {
				hasRpcRequest = ok
			}
		}

		if err != nil {
			Log.Errorf("object type = %T, err = %v", message.M_body, err)
			return err
		}
	}

	// send pending
	c.SendPendingMessages()
	c.PendingMessages = []*PendingMessage{}

	// set user online
	// if c.IsUpdates {
	// 	c.manager.setUserOnline(c.sessionId, connID)
	// }

	// if connID.connType == mtproto.TRANSPORT_TCP {
	// 	if c.isUpdates {
	// 		c.manager.updatesSession.SubscribeUpdates(c, connID)
	// 		// c.manager.setUserOnline(c.sessionId, connID)
	// 		// c.manager.updatesSession.SubscribeUpdates(c, connID)
	// 	}
	// 	c.sendPendingMessagesToClient(connID, md, c.pendingMessages)
	// 	c.pendingMessages = []*pendingMessage{}
	// } else {
	// 	if !hasRpcRequest {
	// 		if len(c.pendingMessages) > 0 {
	// 			c.sendPendingMessagesToClient(connID, md, c.pendingMessages)
	// 			c.pendingMessages = []*pendingMessage{}
	// 		} else {
	// 			c.manager.updatesSession.SubscribeUpdates(c, connID)
	// 			//if !hasHttpWait {
	// 			//	// TODO(@benqi): close http
	// 			//} else {
	// 			//	// c.manager.setUserOnline(c.sessionId, connID)
	// 			//	c.manager.updatesSession.SubscribeUpdates(c, connID)
	// 			//}
	// 		}
	// 	} else {
	// 		// wait
	// 	}
	// }

	_ = hasHttpWait

	// if len(c.rpcMessages) > 0 {
	// 	c.manager.rpcQueue.Push(&rpcApiMessages{connID: connID, md: md, sessionId: c.sessionId, rpcMessages: c.rpcMessages})
	// 	c.rpcMessages = []*networkApiMessage{}
	// }

	return nil
}

func (s *TLService) TLUnencryptedMessageProcess(sess *Session, msg *mtproto.UnencryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("request - class type = %T, \ntlobj = %s", tlobj, tlobj)

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

	Log.Debugf("response - class type = %T, \nresp tlobj = %s", tlobj, res)

	return res, err
}

// func (s *TLService) TLMessageReplyType(tl mtproto.TLObject) TLMessageReplyTypeEnum {

// 	switch tl.(type) {
// 	case *mtproto.TL_ping,
// 		*mtproto.TL_ping_delay_disconnect:

// 		return TLMessageReplyNewMessageID

// 	}

// 	return TLMessageReplyZeroMessageID
// }

// 消息类型判断
func (s *TLService) TLRpcUpdatesType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_account_registerDevice,
		*mtproto.TL_account_unregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return false

	case *mtproto.TL_upload_saveFilePart,
		*mtproto.TL_upload_saveBigFilePart:

		// upload connection
		return false

	case *mtproto.TL_upload_getFile,
		*mtproto.TL_upload_getWebFile,
		*mtproto.TL_upload_getCdnFile,
		*mtproto.TL_upload_reuploadCdnFile,
		*mtproto.TL_upload_getCdnFileHashes:

		// download
		return false

	case *mtproto.TL_help_getConfig:
		// TODO(@benqi): 可能为TEMP，判断TEMP的规则：
		//  从android client代码看，此连接仅收到help.getConfig消息
	}

	return true
}

func (s *TLService) TLRpcPushType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_account_registerDevice,
		*mtproto.TL_account_unregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return true
	}
	return false
}

func (s *TLService) TLRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_auth_checkedPhone,
		*mtproto.TL_auth_sendCode,
		*mtproto.TL_auth_signIn,
		*mtproto.TL_auth_signUp,
		*mtproto.TL_auth_exportedAuthorization,
		*mtproto.TL_auth_exportAuthorization,
		*mtproto.TL_auth_importAuthorization,
		*mtproto.TL_auth_cancelCode,
		*mtproto.TL_auth_resendCode,
		*mtproto.TL_auth_requestPasswordRecovery,
		*mtproto.TL_auth_checkPassword,
		*mtproto.TL_auth_recoverPassword:

		return true

	case *mtproto.TL_help_getConfig,
		*mtproto.TL_help_getCdnConfig:

		return true

	case *mtproto.TL_langpack_getLanguages,
		*mtproto.TL_langpack_getDifference,
		*mtproto.TL_langpack_getLangPack,
		*mtproto.TL_langpack_getStrings:

		return true

	case *mtproto.TL_account_getPassword,
		*mtproto.TL_account_deleteAccount,
		*mtproto.TL_account_updatePasswordSettings,
		*mtproto.TL_account_getPasswordSettings:

		return true
	}

	return false
}

func (s *TLService) TLRpcUploadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_upload_saveFilePart,
		*mtproto.TL_upload_saveBigFilePart,
		*mtproto.TL_upload_reuploadCdnFile:
		return true
	}
	return false
}

func (s *TLService) TLRpcDownloadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_upload_getFile,
		*mtproto.TL_upload_getWebFile,
		*mtproto.TL_upload_getCdnFile,
		*mtproto.TL_upload_getCdnFileHashes:
		return true
	}
	return false
}
