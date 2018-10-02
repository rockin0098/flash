package gate_server

import (
	"time"

	"github.com/rockin0098/flash/base/grmon"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

const (
	SERVER_NAME   = "GateServer"
	SERVER_CONFIG = "config/gate.json"
)

func GateEntry() {

	UseMaxCpu()

	Log.Infof("#############################################")
	Log.Infof("############ " + SERVER_NAME + " ############")
	Log.Infof("#############################################")
	// Log.Infof("BUILD_VERSION = %v", version.BUILD_VERSION)
	// Log.Infof("APPLICATION_VERSION = %v", version.APPLICATION_VERSION)
	// Log.Infof("NUMBER_VERSION = %v", version.NUMBER_VERSION)

	tcpserver := NewTcpServer(":5222")
	grm := grmon.GetGRMon()
	grm.Go("TcpServer", tcpserver.Run)

	Log.Infof("keep main func here ... ...")
	for {
		time.Sleep(time.Duration(3) * time.Second)
	}
}
