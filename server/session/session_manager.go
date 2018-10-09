package session

import (
	"sync"

	"github.com/rockin0098/flash/base/grmon"
	"github.com/rockin0098/flash/base/guid"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/base/mq"
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

var sessionManager SessionManager = &SessionManagerMemory{
	sessionStorage: &sync.Map{},
}

func GetSessionManager() SessionManager {
	return sessionManager
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

// session
type Session struct {
	rw        *sync.RWMutex
	sessionID string
	connID    string
	mtproto   *mtproto.MTProto
	mq        mq.MQ
}

func NewSession(connID string) *Session {

	sess := &Session{
		rw:        &sync.RWMutex{},
		sessionID: GenerateSessionID(),
		connID:    connID,
		mq:        mq.NewMQ(),
	}

	sessionManager.Store(sess.sessionID, sess)

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

	return s.sessionID
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

func (s *Session) MQ() mq.MQ {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.mq
}

func (s *Session) CreateSessionWriter() {
	// mq routine
	s.mq.CreateChannel(s.sessionID) // 唯一channel
	grm := grmon.GetGRMon()
	grm.GoLoop("sess_response_loop", func() {
		data := s.mq.Get(s.sessionID)
		mtp := s.mtproto
		err := mtp.Write(data)
		if err != nil {
			Log.Warn(err)
			mtp.Close()
		}
	})
}

func (s *Session) Write(data interface{}) {
	s.mq.Put(s.sessionID, data)
}
