package service

import (
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	. "github.com/rockin0098/meow/base/global"
)

func TestLangpack(t *testing.T) {

	var langs LangPack

	_, err := toml.DecodeFile("/Users/dev/workspace/golang/gopath/src/github.com/rockin0098/meow/config/lang_pack_en.toml", &langs)
	if err != nil {
		Log.Error(err)
		return
	}

	Log.Infof("%v", FormatStruct(langs))

	ioutil.WriteFile("/Users/dev/workspace/golang/gopath/src/github.com/rockin0098/meow/config/lang.pack.en.json", []byte(FormatStruct(langs)), 0666)

}
