package session

import (
	"sync"

	"github.com/rockin0098/flash/base/guid"
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
	// client session
	StoreClient(clientSessID int64, sess *Session)
	LoadClient(clientSessID int64) (*Session, bool)
	RemoveClient(clientSessID int64)
}

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
type SessionManagerMemory struct {
	serverSessionStorage *sync.Map
	clientSessionStorage *sync.Map
}

var sessionManager SessionManager = &SessionManagerMemory{
	serverSessionStorage: &sync.Map{},
	clientSessionStorage: &sync.Map{},
}

func GetSessionManager() SessionManager {
	return sessionManager
}

func (s *SessionManagerMemory) Store(sessid string, sess *Session) {
	s.serverSessionStorage.Store(sessid, sess)
}

func (s *SessionManagerMemory) Load(sessid string) (*Session, bool) {
	sobj, ok := s.serverSessionStorage.Load(sessid)
	if !ok {
		return nil, false
	}

	return sobj.(*Session), true
}

func (s *SessionManagerMemory) Remove(sessid string) {
	s.serverSessionStorage.Delete(sessid)
}

func (s *SessionManagerMemory) StoreClient(clientSessID int64, sess *Session) {
	s.clientSessionStorage.Store(clientSessID, sess)
}

func (s *SessionManagerMemory) LoadClient(clientSessID int64) (*Session, bool) {
	sobj, ok := s.clientSessionStorage.Load(clientSessID)
	if !ok {
		return nil, false
	}

	return sobj.(*Session), true
}

func (s *SessionManagerMemory) RemoveClient(clientSessID int64) {
	s.clientSessionStorage.Delete(clientSessID)
}

// session
type Session struct {
	rw              *sync.RWMutex
	serverSessionID string
	clientSessionID int64
	connID          string
	mtproto         *mtproto.MTProto
}

func NewSession(connID string) *Session {

	sess := &Session{
		rw:              &sync.RWMutex{},
		serverSessionID: GenerateSessionID(),
		connID:          connID,
	}

	sessionManager.Store(sess.serverSessionID, sess)

	return sess
}

func GetSession(sessid string) *Session {
	sess, ok := sessionManager.Load(sessid)
	if !ok {
		return nil
	}

	return sess
}

func (s *Session) SessionID() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.serverSessionID
}

func (s *Session) ClientSessionID() int64 {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.clientSessionID
}

func (s *Session) ConnectionID() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.connID
}

func (s *Session) SetMTProto(mtp *mtproto.MTProto) {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.mtproto = mtp
}

func (s *Session) MTProto() *mtproto.MTProto {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.mtproto
}
