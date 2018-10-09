package service

import (
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

	Log.Debugf("invokeWithLayer layer = %v, query = %T", layer, query)

	// must be initConnection
	initConn := query.(*mtproto.TL_initConnection)

	initConnection := &mtproto.TL_initConnection{
		M_api_id:           initConn.M_api_id,
		M_device_model:     initConn.M_device_model,
		M_system_version:   initConn.M_system_version,
		M_app_version:      initConn.M_app_version,
		M_system_lang_code: initConn.M_system_lang_code,
		M_lang_pack:        initConn.M_lang_pack,
		M_lang_code:        initConn.M_lang_code,
		M_query:            query,
	}

	return initConnection, nil
}

func (s *LProtoService) TL_initConnection_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	tlobj := msg.TLObject
	tl := tlobj.(*mtproto.TL_initConnection)

	// todo

	return tl, nil
}

func (s *LProtoService) TL_help_getConfig_Process(cltSess *session.ClientSession, msg *mtproto.EncryptedMessage) (interface{}, error) {
	Log.Infof("entering... client sessid = %v", cltSess.SessionID())

	// tlobj := msg.TLObject
	// tl := tlobj.(*mtproto.TL_help_getConfig)

	// layer := tl.Get_layer()
	// query := tl.Get_query()

	// Log.Debugf("invokeWithLayer layer = %v, query = %T", layer, query)

	// // must be initConnection
	// initConn := query.(*mtproto.TL_initConnection)

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

	return nil, nil
}
