package session

import (
	"net"
	"sync"

	"github.com/rockin0098/flash/base/guid"
)

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
var sessionStore = &sync.Map{}

type Session struct {
	sessionID string
	conn      net.Conn
}

func NewSession(conn net.Conn) *Session {

	sess := &Session{
		sessionID: guid.GenerateSessionID(),
		conn:      conn,
	}

	sessionStore.Store(sess.sessionID, sess)

	return sess
}

func GetSession(sessid string) *Session {
	sobj, ok := sessionStore.Load(sessid)
	if !ok {
		return nil
	}

	return sobj.(*Session)
}

func (s *Session) SessionID() string {
	return s.sessionID
}
