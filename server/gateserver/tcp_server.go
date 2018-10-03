package gateserver

import (
	"net"

	"github.com/rockin0098/flash/base/grmon"
	"github.com/rockin0098/flash/mtproto"

	// . "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type TTcpServer struct {
	addr string
}

func NewTcpServer(addr string) *TTcpServer {
	return &TTcpServer{
		addr: addr,
	}
}

func (s *TTcpServer) Run() {
	Log.Infof("TCP server listening at %v", s.addr)
	lsn, err := net.Listen("tcp", s.addr)
	if err != nil {
		Log.Error(err)
		return
	}

	defer lsn.Close()
	for {
		conn, err := lsn.Accept()
		if err != nil {
			Log.Warn(err)
			continue
		}
		Log.Debugf("s:[%v] accept connection from %v", s.addr, conn.RemoteAddr().String())
		grm := grmon.GetGRMon()
		grm.Go("tcp_handler", func() { s.Handler(conn) })
	}
}

func (s *TTcpServer) Handler(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	proto := mtproto.NewMTProto(conn, conn)
	for {
		msg, err := proto.Read()
		if err != nil {
			Log.Warnf("s:[%v], remote: %v, connection error: %v", s.addr, remote, err)
			return
		}
		// Log.Debugf("s:[%v], remote : %v, receive data : %v", s.addr, remote, HexBuffer(buffer[:n]))
		Log.Debugf("s:[%v], remote : %v, receive msg : %v", s.addr, remote, msg)
	}
}
