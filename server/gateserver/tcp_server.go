package gateserver

import (
	"io"
	"net"

	"github.com/rockin0098/flash/base/grmon"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/process"
	"github.com/rockin0098/flash/server/session"

	// . "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type TcpServer struct {
	addr string
}

func NewTcpServer(addr string) *TcpServer {
	return &TcpServer{
		addr: addr,
	}
}

func (s *TcpServer) Run() {
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
		// ch := make(chan int, 0)
		// <-ch
		// for debugging =====> end
	}
}

func (s *TcpServer) ConnectionHandler(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	localAddr := conn.LocalAddr()

	cm := session.GetConnectionManager()
	connid := cm.Add(conn)
	sess := session.NewSession(connid)
	MAX_RESP_CHAN_LEN := 8192
	respChan := make(chan interface{}, MAX_RESP_CHAN_LEN)
	closeChan := make(chan interface{}, 1)
	bufReader := mtproto.NewBufferedReader(conn)
	initial := true

	connClose := func() {
		conn.Close()
		close(respChan)
		close(closeChan)
		cm.Remove(connid)
	}

	// 每个连接一个
	mtp := mtproto.NewMTProto(bufReader, conn, remoteAddr, localAddr, sess.SessionID(), respChan, closeChan)
	if mtp == nil {
		Log.Error("create mtproto failed, will close connection...")
		connClose()
		return
	}
	sess.SetMTProto(mtp)
	sess.CreateSessionWriter()

	grm := grmon.GetGRMon()
	grm.Go("tcp_read", func() {
		for {
			if !initial {
				b := make([]byte, 0)
				n, err := io.ReadFull(conn, b) // 阻塞
				if err != nil || n != 0 {
					Log.Warnf("s:[%v], remote: %v, sess = %v, connection error: %v", s.addr, remoteAddr, sess.SessionID(), err)
					connClose()
					return
				}

				err = mtp.Codec()()
				if err != nil {
					Log.Warnf("s:[%v], remote: %v, sess = %v,connection error: %v", s.addr, remoteAddr, sess.SessionID(), err)
					connClose()
					return
				}
			} else {
				initial = false
			}

			// 暂时没有控制并发数量
			grm.Go("tcp_worker", func() {
				err := process.GateProcess(mtp)
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
			var data interface{}
			var ok bool
			select {
			case data, ok = <-respChan:
				if !ok {
					Log.Warnf("resp channel closed, write routine will exit, sessid = %v", sess.SessionID())
					return
				}
			case <-closeChan:
				Log.Warnf("recieved from close channel, will close the connection, sessid = %v", sess.SessionID())
				conn.Close() // 此处只尝试关闭连接即可
				return
			}

			if data == nil { // 空数据则不处理
				Log.Warnf("tcp will write nil. sess = %v", sess.SessionID())
				continue
			}

			databytes := data.([]byte)
			nlen := len(databytes)
			left := nlen
			for left > 0 {
				sentdata := databytes[nlen-left:]
				// Log.Debugf("sentdata = %v", hex.EncodeToString(sentdata))
				n, err := conn.Write(sentdata)
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
