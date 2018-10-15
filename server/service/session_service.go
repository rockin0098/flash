package service

import (
	"net"
	"sync"

	"github.com/rockin0098/flash/base/guid"
	"github.com/rockin0098/flash/base/tcpnet"
	"github.com/rockin0098/flash/proto/mtproto"
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

func (s *Session) WriteFull(b []byte) error {
	mtp := s.MTProto
	err := mtp.Encode(b)
	return err
}
