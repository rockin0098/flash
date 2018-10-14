package agentserver

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/rockin0098/flash/base/datasource"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

const (
	SERVER_NAME    = "AgentServer"
	SERVER_CONFIG  = "config/agent.json"
	MTPROTO_CONFIG = "config/config.json"
)

type GateServerConfig struct {
	NodeID     string
	DataSource []datasource.DataSource
}

var gateServerConfig = &GateServerConfig{}

func AgentEntry() {

	UseMaxCpu()

	Log.Infof("#############################################")
	Log.Infof("################ " + SERVER_NAME + " #################")
	Log.Infof("#############################################")
	// Log.Infof("BUILD_VERSION = %v", version.BUILD_VERSION)
	// Log.Infof("APPLICATION_VERSION = %v", version.APPLICATION_VERSION)
	// Log.Infof("NUMBER_VERSION = %v", version.NUMBER_VERSION)

	Log.Info("server init...")
	ServerInit()

	Log.Info("tcp server init...")
	// grm := grmon.GetGRMon()
	// grm.Go("TcpServe", NewTcpServer(":5222").Run)

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
	)
}
