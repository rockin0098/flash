package service

import (
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/rockin0098/flash/base/global"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"
)

func (s *TLService) TLRpcMessageProcess(csess *ClientSession, msgId int64, seqNo int32, object mtproto.TLObject) error {

	tlobj := object
	c := csess

	Log.Debugf("request - client sessid = %v,  class type = %T, \ntlobj = %v",
		csess.ClientSessionID, tlobj, tlobj)

	ASSERT(c.AuthKeyID != 0)

	if s.TLRpcUpdatesType(object) {
		c.IsUpdates = true
	}

	// doing rpc

	var err error
	var res mtproto.TLObject

	// rpc invoke
	switch tlobj.(type) {
	case *mtproto.TL_help_getConfig:
		res, err = s.TL_help_getConfig_Process(csess, tlobj)
	case *mtproto.TL_auth_logOut:
		res, err = s.TL_auth_logOut_Process(csess, tlobj)
	case *mtproto.TL_langpack_getLangPack:
		res, err = s.TL_langpack_getLangPack_Process(csess, tlobj)
	case *mtproto.TL_help_getNearestDc:
		res, err = s.TL_help_getNearestDc_Process(csess, tlobj)
	case *mtproto.TL_auth_checkPhone:
		res, err = s.TL_auth_checkPhone_Process(csess, tlobj)

	default:
		Log.Debugf("havent implemented yet, TLType = %T", tlobj)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tlobj)
	}

	if err != nil {
		Log.Errorf("object = %T, err = %v", tlobj, err)
		return err
	}

	reply := &mtproto.TL_rpc_result{
		M_req_msg_id: msgId,
		M_result:     res,
	}

	err = c.WriteFull(c.AuthKeyID, mtproto.GenerateMessageID(), true, reply)

	return err
}

func (s *TLService) TL_help_getConfig_Process(csess *ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	timenow := int32(time.Now().Unix())

	helpConfig := &mtproto.TL_config{
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

	Log.Infof("helpConfig = %v", FormatStruct(helpConfig))

	return helpConfig, nil
}

func (s *TLService) TL_auth_logOut_Process(csess *ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := msg.TLObject
	// tl := tlobj.(*mtproto.TL_auth_logOut)

	tltrue := mtproto.ToBool(true)

	return tltrue, nil
}

// TL_langpack_getLangPack
func (s *TLService) TL_langpack_getLangPack_Process(csess *ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
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

// TL_help_getNearestDc
func (s *TLService) TL_help_getNearestDc_Process(csess *ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	dc := &mtproto.TL_nearestDc{
		M_country:    "US",
		M_this_dc:    2,
		M_nearest_dc: 2,
	}

	return dc, nil
}

// TL_auth_checkPhone
func (s *TLService) TL_auth_checkPhone_Process(csess *ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_checkPhone)

	// to do : check phone number format
	phone := tl.Get_phone_number()

	mm := model.GetModelManager()
	registered := mm.CheckPhoneExists(phone)

	checkedPhone := &mtproto.TL_auth_checkedPhone{
		M_phone_registered: mtproto.ToBool(registered),
	}

	return checkedPhone, nil
}
