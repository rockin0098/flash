package session

import (
	"sync"
	"time"

	"github.com/rockin0098/flash/proto/mtproto"
)

type ClientSessionManager interface {
	// client session
	Store(sessid int64, sess *ClientSession)
	Load(sessid int64) (sess *ClientSession, ok bool)
	Remove(sessid int64)
}

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
type ClientSessionManagerMemory struct {
	sessionStorage *sync.Map
}

var clientSessionManager ClientSessionManager = &ClientSessionManagerMemory{
	sessionStorage: &sync.Map{},
}

func GetClientSessionManager() ClientSessionManager {
	return clientSessionManager
}

func (s *ClientSessionManagerMemory) Store(sessid int64, sess *ClientSession) {
	s.sessionStorage.Store(sessid, sess)
}

func (s *ClientSessionManagerMemory) Load(sessid int64) (*ClientSession, bool) {
	sobj, ok := s.sessionStorage.Load(sessid)
	if !ok {
		return nil, false
	}

	return sobj.(*ClientSession), true
}

func (s *ClientSessionManagerMemory) Remove(sessid int64) {
	s.sessionStorage.Delete(sessid)
}

// session
type ClientSession struct {
	rw              *sync.RWMutex
	Layer           int32
	sessionID       int64
	authID          int64
	salt            int64
	firstMessageID  int64
	nextSeqNo       uint32
	serverSessionID string
}

func NewClientSession(clientSessionID int64, authid int64, salt int64, firstMessageID int64, serverSessionID string) *ClientSession {

	cltsess := &ClientSession{
		rw:             &sync.RWMutex{},
		sessionID:      clientSessionID,
		authID:         authid,
		salt:           salt,
		firstMessageID: firstMessageID,

		// save server sess
		serverSessionID: serverSessionID,
	}

	auth := &Auth{
		rw:             &sync.RWMutex{},
		sessionID:      clientSessionID,
		firstMessageID: firstMessageID,
	}

	clientSessionManager.Store(clientSessionID, cltsess)
	authManager.Store(authid, auth)

	return cltsess
}

func GetClientSession(clientSessionID int64) *ClientSession {
	sess, ok := clientSessionManager.Load(clientSessionID)
	if !ok {
		return nil
	}

	return sess
}

func (s *ClientSession) SessionID() int64 {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.sessionID
}

func (s *ClientSession) AuthKeyID() int64 {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.authID
}

func (s *ClientSession) ServerSessionID() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.serverSessionID
}

func (s *ClientSession) Write(data interface{}) {
	sess := GetSession(s.serverSessionID)
	sess.Write(data)
}

//// Check Server Salt
func (s *ClientSession) CheckBadServerSalt(authid int64, msgId int64, seqNo int32, salt int64) (*mtproto.TL_bad_server_salt, bool) {
	// Notice of Ignored Error Message
	//
	// Here, error_code can also take on the following values:
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//

	am := GetAuthManager()
	if !am.CheckBySalt(authid, salt) {
		s.salt, _ = am.GetOrInsertSalt(authid)
		badServerSalt := &mtproto.TL_bad_server_salt{
			M_bad_msg_id:      msgId,
			M_error_code:      48,
			M_bad_msg_seqno:   seqNo,
			M_new_server_salt: s.salt,
		}
		// c.sendToClient(connID, md, 0, false, badServerSalt.To_BadMsgNotification())
		return badServerSalt, false
	}

	return nil, true
}

func (s *ClientSession) CheckBadMsgNotification(authid int64, msgId int64, seqNo int32, isContainer bool) (*mtproto.TL_bad_msg_notification, bool) {
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
		// c.sendToClient(connID, md, 0, false, badMsgNotification.To_BadMsgNotification())
		return badMsgNotification, false
	}

	return nil, true
}

func (s *ClientSession) EncodeMessage(authKeyID int64, authKey []byte, messageID int64, confirm bool, tl mtproto.TLObject) []byte {
	message := &mtproto.EncryptedMessage{
		Salt:            s.salt,
		SeqNo:           s.generateMessageSeqNo(confirm),
		MessageID:       messageID,
		ClientSessionID: s.sessionID,
		TLObject:        tl,
	}
	return message.Encode(authKeyID, authKey)
}

func (s *ClientSession) generateMessageSeqNo(increment bool) int32 {
	value := s.nextSeqNo
	if increment {
		s.nextSeqNo++
		return int32(value*2 + 1)
	} else {
		return int32(value * 2)
	}
}
