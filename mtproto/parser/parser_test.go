package parser

import (
	"testing"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

func TestMain(m *testing.M) {
	Log.Info("test starting...")
	m.Run()
	Log.Info("test ending...")
}

func TestParseTL(t *testing.T) {
	schema := ParseTL(TLContent)
	Log.Infof("schema = %v", FormatStruct(schema))
}

func TestTrimAnnotation(t *testing.T) {
	ss := trimAnnotation("asds // asdasdasd")
	Log.Info("ss = ", ss)
}
