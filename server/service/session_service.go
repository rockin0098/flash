package service

import (
	"sync"

	"github.com/rockin0098/flash/base/guid"
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

type SessionService struct {
	sessionManager SessionManager
}

var sessionService = &SessionService{}

func SessionServiceInstance() *SessionService {
	return sessionService
}

type Session struct {
	connID string
}

func (s *SessionService) CreateSession(connid string) *Session {

	return nil
}
