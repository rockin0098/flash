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
		M_phonecalls_enabled:         mtproto.ToBool2(mtproto.Config.M_phonecalls_enabled),
		M_default_p2p_contacts:       mtproto.ToBool2(mtproto.Config.M_default_p2p_contacts),
		M_date:                       timenow,
		M_expires:                    timenow + EXPIRE_TIMEOUT,
		M_test_mode:                  mtproto.ToBool(mtproto.Config.M_test_mode),
		M_this_dc:                    mtproto.Config.M_this_dc,
		M_dc_options:                 nil,
		M_chat_size_max:              mtproto.Config.M_chat_size_max,
		M_megagroup_size_max:         mtproto.Config.M_megagroup_size_max,
		M_forwarded_count_max:        mtproto.Config.M_forwarded_count_max,
		M_online_update_period_ms:    mtproto.Config.M_online_update_period_ms,
		M_offline_blur_timeout_ms:    mtproto.Config.M_offline_blur_timeout_ms,
		M_offline_idle_timeout_ms:    mtproto.Config.M_offline_idle_timeout_ms,
		M_online_cloud_timeout_ms:    mtproto.Config.M_online_cloud_timeout_ms,
		M_notify_cloud_delay_ms:      mtproto.Config.M_notify_cloud_delay_ms,
		M_notify_default_delay_ms:    mtproto.Config.M_notify_default_delay_ms,
		M_chat_big_size:              mtproto.Config.M_chat_big_size,
		M_push_chat_period_ms:        mtproto.Config.M_push_chat_period_ms,
		M_push_chat_limit:            mtproto.Config.M_push_chat_limit,
		M_saved_gifs_limit:           mtproto.Config.M_saved_gifs_limit,
		M_edit_time_limit:            mtproto.Config.M_edit_time_limit,
		M_rating_e_decay:             mtproto.Config.M_rating_e_decay,
		M_stickers_recent_limit:      mtproto.Config.M_stickers_recent_limit,
		M_stickers_faved_limit:       mtproto.Config.M_stickers_faved_limit,
		M_channels_read_media_period: mtproto.Config.M_channels_read_media_period,
		M_tmp_sessions:               mtproto.Config.M_tmp_sessions,
		M_pinned_dialogs_count_max:   mtproto.Config.M_pinned_dialogs_count_max,
		M_call_receive_timeout_ms:    mtproto.Config.M_call_receive_timeout_ms,
		M_call_ring_timeout_ms:       mtproto.Config.M_call_ring_timeout_ms,
		M_call_connect_timeout_ms:    mtproto.Config.M_call_connect_timeout_ms,
		M_call_packet_timeout_ms:     mtproto.Config.M_call_packet_timeout_ms,
		M_me_url_prefix:              mtproto.Config.M_me_url_prefix,
		M_suggested_lang_code:        mtproto.Config.M_suggested_lang_code,
		M_lang_pack_version:          mtproto.Config.M_lang_pack_version,
		M_disabled_features:          nil,
	}

	for _, df := range mtproto.Config.M_disabled_features {
		disabledFeature := &mtproto.TL_disabledFeature{
			M_classID:     mtproto.TL_CLASS_disabledFeature,
			M_feature:     df.M_feature,
			M_description: df.M_description,
		}
		helpConfig.M_disabled_features = append(helpConfig.M_disabled_features, disabledFeature)
	}

	for _, do := range mtproto.Config.M_dc_options {
		dcOption := &mtproto.TL_dcOption{
			M_classID:    mtproto.TL_CLASS_dcOption,
			M_ipv6:       mtproto.ToBool2(do.M_ipv6),
			M_media_only: mtproto.ToBool2(do.M_media_only),
			M_tcpo_only:  mtproto.ToBool2(do.M_tcpo_only),
			M_cdn:        mtproto.ToBool2(do.M_cdn),
			M_static:     mtproto.ToBool2(do.M_static),
			M_id:         do.M_id,
			M_ip_address: do.M_ip_address,
			M_port:       do.M_port,
		}

		helpConfig.M_dc_options = append(helpConfig.M_dc_options, dcOption)
	}

	return &helpConfig, nil
}

func (s *LProtoService) TL_msg_container_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	return nil, nil
}
