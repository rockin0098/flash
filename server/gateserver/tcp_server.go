package gateserver

import (
	"net"

	"github.com/rockin0098/flash/base/grmon"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type TTcpServer struct {
	Addr string
}

func NewTcpServer(addr string) *TTcpServer {
	return &TTcpServer{
		Addr: addr,
	}
}

func (s *TTcpServer) Run() {
	Log.Infof("TCP server listening at %v", s.Addr)
	netListen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		Log.Error(err)
		return
	}

	defer netListen.Close()
	for {
		conn, err := netListen.Accept()
		if err != nil {
			Log.Warn(err)
			continue
		}
		Log.Debugf("s:[%v] accept connection from %v", s, conn.RemoteAddr().String())
		grm := grmon.GetGRMon()
		grm.Go("tcp_handler", func() { s.Handler(conn) })
	}
}

func (s *TTcpServer) Handler(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log.Errorf("s:[%v], remote: %v, connection error: %v", s, remote, err)
			return
		}
		Log.Debugf("s:[%v], remote : %v, receive data : %v", s, remote, HexBuffer(buffer[:n]))
	}
}
