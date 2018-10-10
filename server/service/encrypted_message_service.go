package service

import (
	"time"

	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/session"
)

func (s *LProtoService) TL_ping_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	tlobj := msg.TLObject
	tl := tlobj.(*mtproto.TL_ping)

	pong := &mtproto.TL_pong{
		M_msg_id:  msg.MessageID,
		M_ping_id: tl.Get_ping_id(),
	}

	return pong, nil
}

func (s *LProtoService) TL_invokeWithLayer_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	tlobj := msg.TLObject
	tl := tlobj.(*mtproto.TL_invokeWithLayer)

	layer := tl.Get_layer()
	query := tl.Get_query()

	Log.Debugf("invokeWithLayer layer = %v, query = %T, \nquery = %s", layer, query, query)

	// must be initConnection
	initConn := query.(*mtproto.TL_initConnection)

	// initConnection := &mtproto.TL_initConnection{
	// 	M_api_id:           initConn.M_api_id,
	// 	M_device_model:     initConn.M_device_model,
	// 	M_system_version:   initConn.M_system_version,
	// 	M_app_version:      initConn.M_app_version,
	// 	M_system_lang_code: initConn.M_system_lang_code,
	// 	M_lang_pack:        initConn.M_lang_pack,
	// 	M_lang_code:        initConn.M_lang_code,
	// 	M_query:            query,
	// }

	msg2 := *msg
	msg2.TLObject = initConn

	return s.MTProtoEncryptedMessageProcess(cltSess, &msg2)
}

func (s *LProtoService) TL_initConnection_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	tlobj := msg.TLObject
	tl := tlobj.(*mtproto.TL_initConnection)

	Log.Infof("tl.Query = %v", tl.Get_query().String())

	// todo

	msg2 := *msg
	msg2.TLObject = tl.Get_query()

	return s.MTProtoEncryptedMessageProcess(cltSess, &msg2)
}

func (s *LProtoService) TL_help_getConfig_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	timenow := int32(time.Now().Unix())

	helpConfig := mtproto.TL_config{
		M_classID:                    mtproto.TL_CLASS_config,
		M_phonecalls_enabled:         nil,
		M_default_p2p_contacts:       nil,
		M_date:                       timenow,
		M_expires:                    timenow + EXPIRE_TIMEOUT,
		M_test_mode:                  nil,
		M_this_dc:                    2,
		M_dc_options:                 nil,
		M_chat_size_max:              200,
		M_megagroup_size_max:         30000,
		M_forwarded_count_max:        100,
		M_online_update_period_ms:    120000,
		M_offline_blur_timeout_ms:    5000,
		M_offline_idle_timeout_ms:    30000,
		M_online_cloud_timeout_ms:    300000,
		M_notify_cloud_delay_ms:      30000,
		M_notify_default_delay_ms:    1500,
		M_chat_big_size:              10,
		M_push_chat_period_ms:        60000,
		M_push_chat_limit:            2,
		M_saved_gifs_limit:           200,
		M_edit_time_limit:            172800,
		M_rating_e_decay:             2419200,
		M_stickers_recent_limit:      30,
		M_stickers_faved_limit:       5,
		M_channels_read_media_period: 0,
		M_tmp_sessions:               0,
		M_pinned_dialogs_count_max:   5,
		M_call_receive_timeout_ms:    20000,
		M_call_ring_timeout_ms:       90000,
		M_call_connect_timeout_ms:    30000,
		M_call_packet_timeout_ms:     10000,
		M_me_url_prefix:              "https://t.me/",
		M_suggested_lang_code:        "en",
		M_lang_pack_version:          31,
		M_disabled_features:          nil,
	}

	disabledFeature := &mtproto.TL_disabledFeature{
		M_classID:     mtproto.TL_CLASS_disabledFeature,
		M_feature:     "",
		M_description: "",
	}
	helpConfig.M_disabled_features = append(helpConfig.M_disabled_features, disabledFeature)

	dcOption := &mtproto.TL_dcOption{
		M_classID:    mtproto.TL_CLASS_dcOption,
		M_id:         2,
		M_ip_address: "10.30.0.28",
		M_port:       5222,
	}

	helpConfig.M_dc_options = append(helpConfig.M_dc_options, dcOption)

	return &helpConfig, nil
}

func (s *LProtoService) TL_msg_container_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	return nil, nil
}
