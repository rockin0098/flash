package session

import "net"

// session 目前是单机版, 之后要改成分布式版本, 存储到 redis 或独立的session 服务
type Session struct {
	conn net.Conn
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}
