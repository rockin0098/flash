package service

import (
	"time"

	// . "github.com/rockin0098/flash/base/global"
	"github.com/rockin0098/flash/proto/mtproto"
)

const (
	EXPIRE_TIMEOUT = 3600
)

const (
	kNetworkMessageStateNone             = 0 // created
	kNetworkMessageStateReceived         = 1 // received from client
	kNetworkMessageStateRunning          = 2 // invoke api
	kNetworkMessageStateWaitReplyTimeout = 3 // invoke timeout
	kNetworkMessageStateInvoked          = 4 // invoke ok, send to client
	kNetworkMessageStatePushSync         = 5 // invoke ok, send to client
	kNetworkMessageStateAck              = 6 // received client ack
	kNetworkMessageStateWaitAckTimeout   = 7 // wait ack timeout
	kNetworkMessageStateError            = 8 // invalid error
	kNetworkMessageStateEnd              = 9 // end state
)

func (s *TLService) makePendingMessage(messageId int64, confirm bool, tl mtproto.TLObject) *PendingMessage {
	return &PendingMessage{messageId, confirm, tl}
}

func (s *TLService) TL_rpc_drop_answer_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_rpc_drop_answer) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	var answer mtproto.TLObject
	c := csess

	var found = false
	for e := c.ApiMessages.Front(); e != nil; e = e.Next() {
		v, _ := e.Value.(*NetworkApiMessage)
		if v.RpcRequest.M_msg_id == request.M_req_msg_id {
			if v.State == kNetworkMessageStateReceived {

				aw := &mtproto.TL_rpc_answer_dropped{
					M_msg_id: request.M_req_msg_id,
					// M_seq_no  int32
					// M_bytes   int32
				}

				answer = aw
			} else if v.State == kNetworkMessageStateInvoked {
				aw := &mtproto.TL_rpc_answer_dropped_running{}
				answer = aw
			} else {
				aw := &mtproto.TL_rpc_answer_unknown{}
				answer = aw
			}
			found = true
			break
		}
	}

	if !found {
		aw := &mtproto.TL_rpc_answer_dropped_running{}
		answer = aw
	}

	c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, true, &mtproto.TL_rpc_result{M_req_msg_id: msgId, M_result: answer}))

	return nil
}

func (s *TLService) TL_get_future_salts_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_get_future_salts) error {

	c := csess
	as := AuthServiceInstance()
	salts, _ := as.GetOrInsertSaltList(c.AuthKeyID, int(request.M_num))
	futureSalts := &mtproto.TL_future_salts{
		M_req_msg_id: msgId,
		M_now:        int32(time.Now().Unix()),
	}

	for _, sa := range salts {
		futureSalts.M_salts = append(futureSalts.M_salts, sa)
	}

	c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, futureSalts))

	return nil
}

func (s *TLService) TL_ping_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_ping) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	c := csess

	pong := &mtproto.TL_pong{
		M_msg_id:  msgId,
		M_ping_id: request.Get_ping_id(),
	}

	c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, pong))
	c.CloseDate = time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout

	return nil
}

// TL_ping_delay_disconnect
func (s *TLService) TL_ping_delay_disconnect_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_ping_delay_disconnect) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	c := csess

	pong := &mtproto.TL_pong{
		M_msg_id:  msgId,
		M_ping_id: request.Get_ping_id(),
	}

	// csess.WriteFull(msg.AuthKeyID, mtproto.GenerateMessageID(), false, pong)
	c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, pong))
	c.CloseDate = time.Now().Unix() + int64(request.Get_disconnect_delay()) + kPingAddTimeout
	return nil
}

// TL_destroy_session
func (s *TLService) TL_destroy_session_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_destroy_session) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	c := csess

	if request.M_session_id == c.ClientSessionID {
		// The result of this being applied to the current session is undefined.
		Log.Error("the result of this being applied to the current session is undefined.")

		// TODO(@benqi): handle error???
		return nil
	}

	ss := SessionServiceInstance()
	if _, ok := ss.LoadClientSession(request.M_session_id); ok {
		destroySessionOk := &mtproto.TL_destroy_session_ok{
			M_session_id: request.M_session_id,
		}
		c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, destroySessionOk))

		ss.RemoveClientSession(request.M_session_id)

	} else {
		destroySessionNone := &mtproto.TL_destroy_session_none{
			M_session_id: request.M_session_id,
		}
		c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, destroySessionNone))
	}

	return nil
}

// TL_destroy_auth_key
func (s *TLService) TL_destroy_auth_key_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_destroy_auth_key) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	c := csess

	destroyOK := &mtproto.TL_destroy_auth_key_ok{}
	c.PendingMessages = append(c.PendingMessages, s.makePendingMessage(0, false, destroyOK))

	return nil
}

func (s *TLService) TL_msgs_ack_Process(csess *ClientSession, msgId int64, seqNo int32, request *mtproto.TL_msgs_ack) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	c := csess

	for _, id := range request.Get_msg_ids() {
		for e := c.ApiMessages.Front(); e != nil; e = e.Next() {
			v, _ := e.Value.(*NetworkApiMessage)
			if v.RpcMsgID == id {
				v.State = kNetworkMessageStateAck
				Log.Info("onMsgsAck - networkSyncMessage change kNetworkMessageStateAck")
			}
		}
	}

	return nil
}

func (s *TLService) TL_msgs_state_req_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_msgs_state_info_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_msgs_all_info_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_msg_resend_req_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_msg_detailed_info_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_msg_new_detailed_info_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

func (s *TLService) TL_contest_saveDeveloperInfo_Process(csess *ClientSession, msgId int64, seqNo int32, request mtproto.TLObject) error {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	return nil
}

// func (s *TLService) TL_initConnection_Process(csess *ClientSession, msg *mtproto.EncryptedMessage) error {
// 	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

// 	tlobj := msg.TLObject
// 	tl := tlobj.(*mtproto.TL_initConnection)

// 	Log.Infof("tl.Query = %v", tl.Get_query().String())

// 	// todo
// 	msg2 := *msg
// 	msg2.TLObject = tl.Get_query()

// 	return s.TLEncryptedMessageProcess(csess, &msg2)
// }

// func (s *TLService) TL_msg_container_Process(csess *ClientSession, msg *mtproto.EncryptedMessage) error {
// 	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

// 	seqno := msg.SeqNo
// 	if seqno%2 != 0 {
// 		Log.Error("A container does not require an acknowledgment.")
// 		return nil
// 	}

// 	Log.Infof("msg.TLObject = %v", msg.TLObject.(*mtproto.TL_msg_container))
// 	message2s := msg.TLObject.(*mtproto.TL_msg_container).M_message2s
// 	for idx, m := range message2s {
// 		Log.Infof("idx = %v, m = %T, m = %v", idx, m, m)

// 		msg2 := *msg
// 		msg2.TLObject = m

// 		err := s.TLEncryptedMessageProcess(csess, &msg2)
// 		if err != nil {
// 			Log.Error(err)
// 			continue
// 		}
// 	}

// 	return nil
// }

// //message2 msg_id:long seqno:int bytes:int body:Object = Message; // parsed manually
// func (s *TLService) TL_message2_Process(csess *ClientSession, msg *mtproto.EncryptedMessage) error {
// 	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

// 	Log.Infof("msg.TLObject = %v", msg.TLObject.(*mtproto.TL_message2))
// 	tlmsg2 := msg.TLObject.(*mtproto.TL_message2)

// 	Log.Infof("TL_message2.TLObject = %T, msg_id=%v, seqno=%v, bytes=%v",
// 		tlmsg2.M_body, tlmsg2.M_msg_id, tlmsg2.M_seqno, tlmsg2.M_bytes)

// 	msg2 := *msg
// 	msg2.TLObject = tlmsg2.M_body

// 	err := s.TLEncryptedMessageProcess(csess, &msg2)
// 	if err != nil {
// 		Log.Error(err)
// 		return err
// 	}

// 	return nil
// }

// // TL_msgs_state_req
// func (s *TLService) TL_msgs_state_req_Process(csess *ClientSession, msg *mtproto.EncryptedMessage) error {
// 	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

// 	return nil
// }
