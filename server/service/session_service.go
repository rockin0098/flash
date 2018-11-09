package service

import (
	"container/list"
	"fmt"
	"net"
	"sync"
	"time"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/base/grmon"
	"github.com/rockin0098/meow/base/guid"
	"github.com/rockin0098/meow/base/tcpnet"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
)

const (
	kDefaultPingTimeout = 30
	kPingAddTimeout     = 15
)

const (
	kStateCreated = iota
	kStateOnline
	kStateOffline
)

func GenerateSessionID() string {
	return guid.GenerateStringUID()
}

type SessionManager interface {
	// server session
	Store(sessid string, sess *Session)
	Load(sessid string) (sess *Session, ok bool)
	Remove(sessid string)

	// client session
	StoreClientSession(clientSessID int64, sess *ClientSession)
	LoadClientSession(clientSessID int64) (sess *ClientSession, ok bool)
	RemoveClientSession(clientSessID int64)
}

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
type SessionManagerMemory struct {
	sessionStorage       *sync.Map
	clientSessionStorage *sync.Map
}

func (s *SessionManagerMemory) Store(sessid string, sess *Session) {
	s.sessionStorage.Store(sessid, sess)
}

func (s *SessionManagerMemory) Load(sessid string) (*Session, bool) {
	sobj, ok := s.sessionStorage.Load(sessid)
	if !ok {
		return nil, false
	}

	return sobj.(*Session), true
}

func (s *SessionManagerMemory) Remove(sessid string) {
	s.sessionStorage.Delete(sessid)
}

func (s *SessionManagerMemory) StoreClientSession(clientSessID int64, sess *ClientSession) {
	s.clientSessionStorage.Store(clientSessID, sess)
}

func (s *SessionManagerMemory) LoadClientSession(clientSessID int64) (*ClientSession, bool) {
	sobj, ok := s.clientSessionStorage.Load(clientSessID)
	if !ok {
		return nil, false
	}

	return sobj.(*ClientSession), true
}

func (s *SessionManagerMemory) RemoveClientSession(clientSessID int64) {
	s.clientSessionStorage.Delete(clientSessID)
}

type SessionService struct {
	sessionManager SessionManager
}

var sessionService = &SessionService{
	sessionManager: &SessionManagerMemory{
		sessionStorage:       &sync.Map{},
		clientSessionStorage: &sync.Map{},
	},
}

func SessionServiceInstance() *SessionService {
	return sessionService
}

func (s *SessionService) CreateSession(connid int64) *Session {

	sessid := GenerateSessionID()

	sess := &Session{
		SessionID: sessid,
		ConnID:    connid,
	}

	s.sessionManager.Store(sessid, sess)

	return sess
}

func (s *SessionService) RemoveSession(sessid string) {
	s.sessionManager.Remove(sessid)
}

func (s *SessionService) NewClientSession(clientSessionID int64, sess *ClientSession) {
	s.sessionManager.StoreClientSession(clientSessionID, sess)
	sess.ClientSessionStart() // 启动线程
}

func (s *SessionService) LoadClientSession(clientSessionID int64) (*ClientSession, bool) {
	return s.sessionManager.LoadClientSession(clientSessionID)
}

func (s *SessionService) RemoveClientSession(clientSessionID int64) {
	s.sessionManager.RemoveClientSession(clientSessionID)
}

type Session struct {
	SessionID       string
	ConnID          int64
	TcpContext      *tcpnet.TcpContext
	ClientSessionID int64
	MTProto         *mtproto.MTProto
}

func (s *Session) GetConnByID(connid int64) net.Conn {
	gm := guid.GetGUIDManager()
	return gm.Load(connid).(net.Conn)
}

func (s *Session) WriteFull2(raw *mtproto.RawMessage) error {

	defer CatchPanicWarning()

	b := raw.Payload
	if b == nil || len(b) == 0 {
		return fmt.Errorf("invalid data to write, data = %v", b)
	}

	mtp := s.MTProto
	data, err := mtp.Encode(raw)
	if err != nil {
		Log.Error(err)
		return err
	}

	tcpctx := s.TcpContext
	tcpctx.WriteFull(data)

	return nil
}

// client session

type ClientSession struct {
	Sess            *Session // server session
	ClientSessionID int64
	Layer           int32
	// client session data || runtime variables
	AuthKeyID        int64
	Salt             int64
	FirstMessageID   int64
	NextSeqNo        uint32
	CloseDate        int64
	CloseSessionDate int64
	ApiMessages      *list.List
	UniqueID         int64
	ClientState      int
	PendingMessages  []*PendingMessage
	IsUpdates        bool
	RpcMessages      []*NetworkApiMessage
}

type PendingMessage struct {
	MessageID int64
	Confirm   bool
	TL        mtproto.TLObject
}

type NetworkApiMessage struct {
	Date       int64
	QuickAckId int32 // 0: not use
	RpcRequest *mtproto.TL_message2
	State      int // TODO(@benqi): sync.AtomicInt32
	RpcMsgID   int64
	RpcResult  mtproto.TLObject
}

func (s *ClientSession) WriteFull(messageID int64, confirm bool, tl mtproto.TLObject) error {

	Log.Debugf("response -  client sessid = %v, authid = %v, class type = %T, \ntlobj = %v",
		s.ClientSessionID, s.AuthKeyID, tl, tl)

	b := s.EncodeMessage(s.AuthKeyID, messageID, confirm, tl)
	return s.Sess.WriteFull2(&mtproto.RawMessage{Payload: b})
}

func (s *ClientSession) SendPendingMessages() error {

	pmsgs := s.PendingMessages

	if len(pmsgs) == 0 {
		Log.Debug("no PendingMessages to send")
		return nil
	}

	if len(pmsgs) == 1 {
		return s.WriteFull(pmsgs[0].MessageID, pmsgs[0].Confirm, pmsgs[0].TL)
	} else {
		msgContainer := &mtproto.TL_msg_container{
			M_message2s: make([]*mtproto.TL_message2, 0, len(pmsgs)),
		}

		// var seqno int32
		for _, m := range pmsgs {
			msgId := m.MessageID
			if msgId == 0 {
				msgId = mtproto.GenerateMessageID()
			}
			message2 := &mtproto.TL_message2{
				M_msg_id: msgId,
				M_seqno:  s.generateMessageSeqNo(m.Confirm),
				M_bytes:  int32(len(m.TL.Encode())),
				M_body:   m.TL,
			}
			msgContainer.M_message2s = append(msgContainer.M_message2s, message2)
		}

		return s.WriteFull(0, false, msgContainer)
	}

	return nil
}

func (s *ClientSession) EncodeMessage(authKeyID int64, messageID int64, confirm bool, tl mtproto.TLObject) []byte {
	mm := model.GetModelManager()
	authKey := mm.GetAuthKeyValueByAuthID(authKeyID)
	if authKey == nil {
		Log.Error("authKeyID = %v not found", authKeyID)
		return nil
	}
	message := &mtproto.EncryptedMessage{
		Salt:            s.Salt,
		SeqNo:           s.generateMessageSeqNo(confirm),
		MessageID:       messageID,
		ClientSessionID: s.ClientSessionID,
		TLObject:        tl,
	}
	return message.Encode(authKeyID, authKey)
}

func (s *ClientSession) generateMessageSeqNo(increment bool) int32 {
	value := s.NextSeqNo
	if increment {
		s.NextSeqNo++
		return int32(value*2 + 1)
	} else {
		return int32(value * 2)
	}
}

func (s *ClientSession) ClientSessionStart() {
	Log.Info("entering...")

	grm := grmon.GetGRMon()
	grm.Go("client_session_run_message", func() {
		// for {
		// 	// apimsg := <-s.ApiMessages

		// }
	})

}

//// Check Server Salt
func (s *ClientSession) CheckBadServerSalt(msgId int64, seqNo int32, salt int64) bool {
	// Notice of Ignored Error Message
	//
	// Here, error_code can also take on the following values:
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//

	authid := s.AuthKeyID
	as := AuthServiceInstance()
	if !as.CheckBySalt(authid, salt) {
		s.Salt, _ = as.GetOrInsertSalt(authid)
		badServerSalt := &mtproto.TL_bad_server_salt{
			M_bad_msg_id:      msgId,
			M_error_code:      48,
			M_bad_msg_seqno:   seqNo,
			M_new_server_salt: s.Salt,
		}

		s.WriteFull(0, false, badServerSalt)
		return false
	}

	return true
}

func (s *ClientSession) CheckBadMsgNotification(msgId int64, seqNo int32, isContainer bool) bool {
	// Notice of Ignored Error Message
	//
	// In certain cases, a server may notify a client that its incoming message was ignored for whatever reason.
	// Note that such a notification cannot be generated unless a message is correctly decoded by the server.
	//
	// bad_msg_notification#a7eff811 bad_msg_id:long bad_msg_seqno:int error_code:int = BadMsgNotification;
	// bad_server_salt#edab447b bad_msg_id:long bad_msg_seqno:int error_code:int new_server_salt:long = BadMsgNotification;
	//
	// Here, error_code can also take on the following values:
	//
	//  16: msg_id too low (most likely, client time is wrong;
	//      it would be worthwhile to synchronize it using msg_id notifications
	//      and re-send the original message with the “correct” msg_id or wrap
	//      it in a container with a new msg_id
	//      if the original message had waited too long on the client to be transmitted)
	//  17: msg_id too high (similar to the previous case,
	//      the client time has to be synchronized, and the message re-sent with the correct msg_id)
	//  18: incorrect two lower order msg_id bits (the server expects client message msg_id to be divisible by 4)
	//  19: container msg_id is the same as msg_id of a previously received message (this must never happen)
	//  20: message too old, and it cannot be verified whether the server has received a message with this msg_id or not
	//  32: msg_seqno too low (the server has already received a message with a lower msg_id
	//      but with either a higher or an equal and odd seqno)
	//  33: msg_seqno too high (similarly, there is a message with a higher msg_id
	//      but with either a lower or an equal and odd seqno)
	//  34: an even msg_seqno expected (irrelevant message), but odd received
	//  35: odd msg_seqno expected (relevant message), but even received
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//  64: invalid container.
	//
	// The intention is that error_code values are grouped (error_code >> 4):
	// for example, the codes 0x40 - 0x4f correspond to errors in container decomposition.
	//
	// Notifications of an ignored message do not require acknowledgment (i.e., are irrelevant).
	//
	// Important: if server_salt has changed on the server or if client time is incorrect,
	// any query will result in a notification in the above format.
	// The client must check that it has, in fact,
	// recently sent a message with the specified msg_id, and if that is the case,
	// update its time correction value (the difference between the client’s and the server’s clocks)
	// and the server salt based on msg_id and the server_salt notification,
	// so as to use these to (re)send future messages. In the meantime,
	// the original message (the one that caused the error message to be returned)
	// must also be re-sent with a better msg_id and/or server_salt.
	//
	// In addition, the client can update the server_salt value used to send messages to the server,
	// based on the values of RPC responses or containers carrying an RPC response,
	// provided that this RPC response is actually a match for the query sent recently.
	// (If there is doubt, it is best not to update since there is risk of a replay attack).
	//

	//=============================================================================================
	// TODO(@benqi): Time Synchronization, https://core.telegram.org/mtproto#time-synchronization
	//
	// Time Synchronization
	//
	// If client time diverges widely from server time,
	// a server may start ignoring client messages,
	// or vice versa, because of an invalid message identifier (which is closely related to creation time).
	// Under these circumstances,
	// the server will send the client a special message containing the correct time and
	// a certain 128-bit salt (either explicitly provided by the client in a special RPC synchronization request or
	// equal to the key of the latest message received from the client during the current session).
	// This message could be the first one in a container that includes other messages
	// (if the time discrepancy is significant but does not as yet result in the client’s messages being ignored).
	//
	// Having received such a message or a container holding it,
	// the client first performs a time synchronization (in effect,
	// simply storing the difference between the server’s time
	// and its own to be able to compute the “correct” time in the future)
	// and then verifies that the message identifiers for correctness.
	//
	// Where a correction has been neglected,
	// the client will have to generate a new session to assure the monotonicity of message identifiers.
	//

	var errorCode int32 = 0

	timeMessage := msgId / 4294967296.0
	date := time.Now().Unix()
	// glog.Info("date: ", date, ", timeMessage: ", timeMessage)

	if timeMessage+30 < date {
		errorCode = 16
	}
	if date > timeMessage+300 {
		errorCode = 17
	}

	//=================================================================================================
	// Check Message Identifier (msg_id)
	//
	// https://core.telegram.org/mtproto/description#message-identifier-msg-id
	// Message Identifier (msg_id)
	//
	// A (time-dependent) 64-bit number used uniquely to identify a message within a session.
	// Client message identifiers are divisible by 4,
	// server message identifiers modulo 4 yield 1 if the message is a response to a client message, and 3 otherwise.
	// Client message identifiers must increase monotonically (within a single session),
	// the same as server message identifiers, and must approximately equal unixtime*2^32.
	// This way, a message identifier points to the approximate moment in time the message was created.
	// A message is rejected over 300 seconds after it is created or 30 seconds
	// before it is created (this is needed to protect from replay attacks).
	// In this situation,
	// it must be re-sent with a different identifier (or placed in a container with a higher identifier).
	// The identifier of a message container must be strictly greater than those of its nested messages.
	//
	// Important: to counter replay-attacks the lower 32 bits of msg_id passed
	// by the client must not be empty and must present a fractional
	// part of the time point when the message was created.
	//
	if msgId%4 != 0 {
		errorCode = 18
	}

	// TODO(@benqi): other error code??
	if errorCode != 0 {
		badMsgNotification := &mtproto.TL_bad_msg_notification{
			M_bad_msg_id:    msgId,
			M_bad_msg_seqno: seqNo,
			M_error_code:    errorCode,
		}
		// glog.Info("badMsgNotification - ", badMsgNotification)
		// _ = badMsgNotification
		s.WriteFull(msgId, false, badMsgNotification)
		return false
	}

	return true
}
