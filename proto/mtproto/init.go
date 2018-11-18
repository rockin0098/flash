package mtproto

import "github.com/rockin0098/meow/base/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.Log
}
