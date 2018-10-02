package gate_server

import (
	"net"

	"github.com/rockin0098/flash/base/grmon"

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
		Log.Debugf("accept connection from %v", conn.RemoteAddr().String())
		grm := grmon.GetGRMon()
		grm.Go("tcp_handler", s.Handler, conn)
	}
}

func (s *TTcpServer) Handler(connobj interface{}) {
	conn := connobj.(net.Conn)
	remote := conn.RemoteAddr().String()
	buffer := make([]byte, 0, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log.Errorf("remote: %v, connection error: %v", remote, err)
			return
		}
		Log.Debug("remote : %v, receive data string: %v", remote, string(buffer[:n]))
	}
}
