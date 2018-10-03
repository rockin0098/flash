package gateserver

import (
	"bufio"
	"net"

	"github.com/rockin0098/flash/base/grmon"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/process"

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
		grm.Go("tcp_handler", func() { s.ConnectionHandler(conn) })
	}
}

func (s *TTcpServer) ConnectionHandler(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	localAddr := conn.LocalAddr()

	for {
		mtp := mtproto.NewMTProto(bufio.NewReader(conn), conn, remoteAddr, localAddr)
		err := mtp.Read()
		if err != nil {
			Log.Warnf("s:[%v], remote: %v, connection error: %v", s.addr, remoteAddr, err)
			conn.Close()
			return
		}

		err = process.MTProtoProcess(mtp)
		if err != nil {
			Log.Error(err)
			conn.Close()
			return
		}
	}
}
