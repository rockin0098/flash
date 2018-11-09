package parser

import (
	"testing"

	. "github.com/rockin0098/meow/base/global"
	. "github.com/rockin0098/meow/base/logger"
)

func TestMain(m *testing.M) {
	Log.Info("test starting...")
	m.Run()
	Log.Info("test ending...")
}

func TestParseTL(t *testing.T) {
	layer := ParseTLLayer(TLContent)
	Log.Infof("layer = %v", FormatStruct(layer))
}

func TestTrimAnnotation(t *testing.T) {
	ss := trimAnnotation("asds // asdasdasd")
	Log.Info("ss = ", ss)
}
