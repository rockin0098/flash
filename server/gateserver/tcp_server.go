package gateserver

import (
	"bufio"
	"net"

	"github.com/rockin0098/flash/base/grmon"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/process"
	"github.com/rockin0098/flash/server/session"

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

	sess := session.NewSession(conn)
	MAX_RESP_CHAN_LEN := 4096
	respChan := make(chan interface{}, MAX_RESP_CHAN_LEN)

	grm := grmon.GetGRMon()
	grm.Go("tcp_read", func() {
		for {
			mtp := mtproto.NewMTProto(bufio.NewReader(conn), conn, remoteAddr, localAddr, sess.SessionID(), respChan)
			err := mtp.Read()
			if err != nil {
				Log.Warnf("s:[%v], remote: %v, connection error: %v", s.addr, remoteAddr, err)
				conn.Close()
				return
			}

			// 暂时没有控制并发数量
			grm.Go("tcp_worker", func() {
				err = process.GateProcess(mtp)
				if err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
			})
		}
	})

	grm.Go("tcp_write", func() {
		for {
			data := <-respChan
			if data == nil { // 空数据则不处理
				Log.Warnf("tcp write nil. sess = %+v", sess)
				continue
			}

			databytes := data.([]byte)
			nlen := len(databytes)
			left := nlen
			for left > 0 {
				n, err := conn.Write(databytes)
				if err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				left = left - n
			}
		}
	})
}
