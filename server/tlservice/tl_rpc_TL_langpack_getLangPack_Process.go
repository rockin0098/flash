package tlservice

import (
	"io/ioutil"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

// TL_langpack_getLangPack
func (s *TLService) TL_langpack_getLangPack_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_langpack_getLangPack)

	var langpack LangPack

	content, err := ioutil.ReadFile("./config/lang.pack.en.json")
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	err = UnserializeFromJson(string(content), &langpack)
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	diff := mtproto.New_TL_langPackDifference()
	diff.Set_lang_code(tl.M_lang_code)
	diff.Set_version(langpack.Version)
	diff.Set_from_version(0)

	diffss := make([]*mtproto.TL_langPackString, 0)
	for _, ss := range langpack.Strings {
		diffss = append(diffss, &mtproto.TL_langPackString{
			M_key:   ss.Key,
			M_value: ss.Value,
		})
	}

	for _, ss := range langpack.StringPluralizeds {
		diffss = append(diffss, &mtproto.TL_langPackString{
			M_key:   ss.Key,
			M_value: ss.Value,
		})
	}

	return diff, nil
}
