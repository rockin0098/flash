package gateserver

import (
	"bufio"
	"bytes"
	"encoding/hex"
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

		// for debugging, only allow one connection
		ch := make(chan int, 0)
		<-ch
		// for debugging =====> end
	}
}

func (s *TTcpServer) ConnectionHandler(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	localAddr := conn.LocalAddr()

	cm := session.GetConnectionManager()
	connid := cm.Add(conn)
	sess := session.NewSession(connid)
	MAX_RESP_CHAN_LEN := 8192
	respChan := make(chan interface{}, MAX_RESP_CHAN_LEN)
	bufReader := bufio.NewReader(conn)

	connClose := func() {
		conn.Close()
		cm.Remove(connid)
	}

	grm := grmon.GetGRMon()
	grm.Go("tcp_read", func() {
		for {
			// for debugging
			// s.test_read(conn)
			// for debbuging ===> end
			var buf *bytes.Buffer

			{
				b := make([]byte, 1024)
				n, err := conn.Read(b)
				if err != nil {
					Log.Error(err)
					connClose()
					return
				}

				Log.Infof("test_read n : %v", n)
				Log.Infof("test_read b : %v", hex.EncodeToString(b[:n]))

				buf = bytes.NewBuffer(b[:n])
			}

			bufReader = bufio.NewReader(buf)
			mtp := mtproto.NewMTProto(bufReader, conn, remoteAddr, localAddr, sess.SessionID(), respChan)
			err := mtp.Read()
			if err != nil {
				Log.Warnf("s:[%v], remote: %v, connection error: %v", s.addr, remoteAddr, err)
				connClose()
				return
			}

			// 暂时没有控制并发数量
			grm.Go("tcp_worker", func() {
				err = process.GateProcess(mtp)
				if err != nil {
					Log.Error(err)
					connClose()
					return
				}
			})
		}
	})

	grm.Go("tcp_write", func() {
		for {
			data := <-respChan
			if data == nil { // 空数据则不处理
				Log.Warnf("tcp will write nil. sess = %+v", sess)
				continue
			}

			databytes := data.([]byte)
			nlen := len(databytes)
			left := nlen
			for left > 0 {
				n, err := conn.Write(databytes)
				if err != nil {
					Log.Error(err)
					connClose()
					return
				}
				left = left - n
			}
		}
	})
}

func (s *TTcpServer) test_read(conn net.Conn) []byte {
	b := make([]byte, 512)
	n, err := conn.Read(b)
	if err != nil {
		Log.Error(err)
		return nil
	}

	Log.Infof("test_read n : %v", n)
	Log.Infof("test_read b : %v", hex.EncodeToString(b[:n]))

	return b
}
