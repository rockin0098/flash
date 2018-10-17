package service

import (
	"fmt"
	"net"
	"sync"

	. "github.com/rockin0098/flash/base/global"
	"github.com/rockin0098/flash/base/guid"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/base/tcpnet"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"
)

func GenerateSessionID() string {
	return guid.GenerateStringUID()
}

type SessionManager interface {
	// server session
	Store(sessid string, sess *Session)
	Load(sessid string) (sess *Session, ok bool)
	Remove(sessid string)
}

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
type SessionManagerMemory struct {
	sessionStorage *sync.Map
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

type SessionService struct {
	sessionManager SessionManager
}

var sessionService = &SessionService{
	sessionManager: &SessionManagerMemory{
		sessionStorage: &sync.Map{},
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

type Session struct {
	SessionID       string
	ConnID          int64
	TcpContext      *tcpnet.TcpContext
	ClientSessionID int64
	MTProto         *mtproto.MTProto

	// runtime variables
	AuthKeyID      int64
	Salt           int64
	FirstMessageID int64
	NextSeqNo      uint32
}

func (s *Session) GetConnByID(connid int64) net.Conn {
	gm := guid.GetGUIDManager()
	return gm.Load(connid).(net.Conn)
}

func (s *Session) WriteFull(authKeyID int64, messageID int64, confirm bool, tl mtproto.TLObject) error {

	Log.Debugf("response - sessid = %v, client sessid = %v, authid = %v, class type = %T, \ntlobj = %v",
		s.SessionID, s.ClientSessionID, authKeyID, tl, tl)

	b := s.EncodeMessage(authKeyID, messageID, confirm, tl)
	return s.WriteFull2(&mtproto.RawMessage{Payload: b})
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

func (s *Session) EncodeMessage(authKeyID int64, messageID int64, confirm bool, tl mtproto.TLObject) []byte {
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

func (s *Session) generateMessageSeqNo(increment bool) int32 {
	value := s.NextSeqNo
	if increment {
		s.NextSeqNo++
		return int32(value*2 + 1)
	} else {
		return int32(value * 2)
	}
}

//// Check Server Salt
func (s *Session) CheckBadServerSalt(authid int64, msgId int64, seqNo int32, salt int64) (*mtproto.TL_bad_server_salt, bool) {
	// Notice of Ignored Error Message
	//
	// Here, error_code can also take on the following values:
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//

	as := AuthServiceInstance()
	if !as.CheckBySalt(authid, salt) {
		salt, _ = as.GetOrInsertSalt(authid)
		badServerSalt := &mtproto.TL_bad_server_salt{
			M_bad_msg_id:      msgId,
			M_error_code:      48,
			M_bad_msg_seqno:   seqNo,
			M_new_server_salt: salt,
		}

		// c.sendToClient(connID, md, 0, false, badServerSalt.To_BadMsgNotification())
		return badServerSalt, false
	}

	return nil, true
}
