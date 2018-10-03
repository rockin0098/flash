package gateserver

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
	Log.Infof("################ " + SERVER_NAME + " #################")
	Log.Infof("#############################################")
	// Log.Infof("BUILD_VERSION = %v", version.BUILD_VERSION)
	// Log.Infof("APPLICATION_VERSION = %v", version.APPLICATION_VERSION)
	// Log.Infof("NUMBER_VERSION = %v", version.NUMBER_VERSION)

	tcpserver := NewTcpServer(":5222")
	grm := grmon.GetGRMon()
	grm.Go("TcpServer", tcpserver.Run)

	// tcpserver2 := NewTcpServer(":8800")
	// grm.Go("TcpServer2", tcpserver2.Run)

	// tcpserver3 := NewTcpServer(":12345")
	// grm.Go("TcpServer3", tcpserver3.Run)

	Log.Infof("keep main func here ... ...")
	for {
		time.Sleep(time.Duration(3) * time.Second)
	}
}
