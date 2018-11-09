package mtproto

import (
	"fmt"
	"io/ioutil"

	. "github.com/rockin0098/meow/base/global"
	. "github.com/rockin0098/meow/base/logger"
)

const (
	TL_CLASS_message2      int32 = 1538843921
	TL_CLASS_msg_container int32 = 1945237724
	TL_CLASS_gzip_packed   int32 = 812830625
	TL_CLASS_rpc_result    int32 = -212046591
)

func init() {
	tlObjectClassMap[TL_CLASS_message2] = func() TLObject { return New_TL_message2() }
	tlObjectClassMap[TL_CLASS_msg_container] = func() TLObject { return New_TL_msg_container() }
	tlObjectClassMap[TL_CLASS_gzip_packed] = func() TLObject { return New_TL_gzip_packed() }
	tlObjectClassMap[TL_CLASS_rpc_result] = func() TLObject { return New_TL_rpc_result() }
}

type TLObject interface {
	ClassID() int32
	String() string
	Encode() []byte
	Decode(b []byte) error
}

func NewTLObjectByClassID(classID int32) TLObject {
	m, ok := tlObjectClassMap[classID]
	if !ok {
		return nil
	}
	return m()
}

func ToBool(b bool) TLObject {
	if b {
		return New_TL_boolTrue()
	} else {
		return New_TL_boolFalse()
	}
}

func ToBool2(b bool) TLObject {
	if b {
		return New_TL_true()
	} else {
		return nil
	}
}

func FromBool(b TLObject) bool {
	_, isTrue := b.(*TL_boolTrue)
	if isTrue {
		return true
	}

	_, isFalse := b.(*TL_boolTrue)
	if isFalse {
		return false
	}

	t := fmt.Sprintf("invalid TL type = %T", b)
	panic(t)
}

func LoadConfig(file string) *TLConfig {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		Log.Error(err)
		return nil
	}

	pc := PurifyConfig(string(content))

	err = UnserializeFromJson(pc, Config)
	if err != nil {
		Log.Error(err)
		return nil
	}

	Log.Debugf("mtproto config = %v", FormatStruct(Config))

	return Config
}

var Config = &TLConfig{}

// dcOption#5d8c6cc
type TLDcOption struct {
	M_ipv6       bool
	M_media_only bool
	M_tcpo_only  bool
	M_cdn        bool
	M_static     bool
	M_id         int32
	M_ip_address string
	M_port       int32
}

// disabledFeature#ae636f24
type TLDisabledFeature struct {
	M_feature     string
	M_description string
}

// config#9c840964
type TLConfig struct {
	M_phonecalls_enabled         bool
	M_default_p2p_contacts       bool
	M_date                       int32
	M_expires                    int32
	M_test_mode                  bool
	M_this_dc                    int32
	M_dc_options                 []TLDcOption
	M_chat_size_max              int32
	M_megagroup_size_max         int32
	M_forwarded_count_max        int32
	M_online_update_period_ms    int32
	M_offline_blur_timeout_ms    int32
	M_offline_idle_timeout_ms    int32
	M_online_cloud_timeout_ms    int32
	M_notify_cloud_delay_ms      int32
	M_notify_default_delay_ms    int32
	M_chat_big_size              int32
	M_push_chat_period_ms        int32
	M_push_chat_limit            int32
	M_saved_gifs_limit           int32
	M_edit_time_limit            int32
	M_rating_e_decay             int32
	M_stickers_recent_limit      int32
	M_stickers_faved_limit       int32
	M_channels_read_media_period int32
	M_tmp_sessions               int32
	M_pinned_dialogs_count_max   int32
	M_call_receive_timeout_ms    int32
	M_call_ring_timeout_ms       int32
	M_call_connect_timeout_ms    int32
	M_call_packet_timeout_ms     int32
	M_me_url_prefix              string
	M_suggested_lang_code        string
	M_lang_pack_version          int32
	M_disabled_features          []TL_disabledFeature
}
