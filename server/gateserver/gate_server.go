package gateserver

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/rockin0098/meow/base/datasource"
	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/base/grmon"
	. "github.com/rockin0098/meow/base/logger"
	"github.com/rockin0098/meow/base/tcpnet"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
	"github.com/rockin0098/meow/server/tlservice"
)

const (
	SERVER_NAME    = "GateServer"
	SERVER_CONFIG  = "config/gate.json"
	MTPROTO_CONFIG = "config/config.json"
)

type GateServerConfig struct {
	NodeID     string
	DataSource []datasource.DataSource
}

var gateServerConfig = &GateServerConfig{}

func GateEntry() {

	UseMaxCpu()

	Log.Infof("#############################################")
	Log.Infof("################ " + SERVER_NAME + " #################")
	Log.Infof("#############################################")
	// Log.Infof("BUILD_VERSION = %v", version.BUILD_VERSION)
	// Log.Infof("APPLICATION_VERSION = %v", version.APPLICATION_VERSION)
	// Log.Infof("NUMBER_VERSION = %v", version.NUMBER_VERSION)

	Log.Info("server init...")
	ServerInit()

	Log.Info("server starting...")
	ServerStart()

	Log.Infof("keep main func here ... ...")
	for {
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func LoadConfig() *GateServerConfig {
	cfile := SERVER_CONFIG
	if len(os.Args) == 2 {
		cfile = os.Args[1]
	}

	content, err := ioutil.ReadFile(cfile)
	if err != nil {
		Log.Error(err)
		os.Exit(-1)
	}

	pc := PurifyConfig(string(content))

	err = UnserializeFromJson(pc, gateServerConfig)
	if err != nil {
		Log.Error(err)
		os.Exit(-1)
	}

	return gateServerConfig
}

func ServerInit() {
	serverConfig := LoadConfig()
	mtproto.LoadConfig(MTPROTO_CONFIG)

	datasource.DataSourceInit(
		serverConfig.DataSource,
		new(model.AuthKey),
		new(model.User),
		new(model.AuthPhoneTransaction),
		new(model.AuthUser),
		new(model.UserPassword),
		new(model.UserPresence),
		new(model.UserContact),
		new(model.UserDialog),
		new(model.Message),
		new(model.Chat),
		new(model.PhotoData),
		new(model.UserPtsUpdates),
		new(model.UserQtsUpdates),
		new(model.MessageBox),
	)
}

func ServerStart() {
	newServerStart(":5222")
}

func newServerStart(addr string) {
	server := tcpnet.NewTcpServer(addr)
	server.MaxConnection = 100
	server.MaxWorker = 100
	server.MaxWriteChanLen = 8192
	server.OnAccept = OnAccept
	server.OnData = OnData
	server.OnWork = OnWork
	server.OnClose = OnClose

	grm := grmon.GetGRMon()
	grm.Go("TcpServer_"+addr, func() { server.Run() })
}

func OnAccept(ctx *tcpnet.TcpContext) error {

	ss := service.SessionServiceInstance()
	sess := ss.CreateSession(ctx.ConnID)
	ctx.Set(service.SESSION, sess)
	sess.TcpContext = ctx
	sess.MTProto = mtproto.NewMTProto(sess.SessionID)

	return nil
}

func OnData(ctx *tcpnet.TcpContext) (interface{}, error) {

	conn := ctx.Conn
	sess := ctx.MustGet(service.SESSION).(*service.Session)
	mtp := sess.MTProto
	var err error

	first := make([]byte, 0)
	n, err := conn.Read(first)
	if err != nil {
		Log.Warnf("read fist failed, n = %v, err = %v", n, err)
		return nil, err
	}

	// debug begin
	// bts := make([]byte, 10*4096)
	// n2, err := conn.Read(bts)
	// if err != nil {
	// 	Log.Errorf("debug read failed, err = %v", err)
	// 	return nil, err
	// }

	// content := bts[:n2]
	// Log.Infof("n2 = %v, packet content = %v", n2, hex.EncodeToString(content))

	// buffer := bytes.NewBuffer(content)

	// if ctx.PacketNum == 0 { // 第一次要选择codec
	// 	err = mtp.SelectCodec(buffer)
	// } else {
	// 	ASSERT(mtp.Codec != nil)
	// 	err = mtp.Codec(buffer)
	// }

	// debug end

	if ctx.PacketNum == 0 { // 第一次要选择codec
		err = mtp.SelectCodec(conn)
	} else {
		ASSERT(mtp.Codec != nil)
		err = mtp.Codec(conn)
	}

	if err != nil {
		Log.Error(err)
		return nil, err
	}

	return mtp.Message, nil
}

func OnWork(ctx *tcpnet.TcpContext, m interface{}) error {

	sess := ctx.MustGet(service.SESSION).(*service.Session)
	msg := m.(*mtproto.RawMessage)

	tls := tlservice.TLServiceInstance()
	err := tls.TLMessageProcess(sess, msg)
	if err != nil {
		Log.Error(err)
		return err
	}

	return nil
}

func OnClose(ctx *tcpnet.TcpContext) error {

	sess := ctx.MustGet(service.SESSION).(*service.Session)
	Log.Infof("connid = %v, sessid = %v will be closed", ctx.ConnID, sess.SessionID)

	ss := service.SessionServiceInstance()
	ss.RemoveSession(sess.SessionID)

	return nil
}
