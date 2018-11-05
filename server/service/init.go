package service

import "github.com/rockin0098/flash/base/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.Log
}
