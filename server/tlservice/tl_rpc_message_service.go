package tlservice

import (
	"fmt"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

func (s *TLService) TLRpcMessageProcess(csess *service.ClientSession, msgid int64, seqNo int32, object mtproto.TLObject) error {

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

	// help
	case *mtproto.TL_help_getConfig:
		res, err = s.TL_help_getConfig_Process(csess, tlobj)
	case *mtproto.TL_help_getNearestDc:
		res, err = s.TL_help_getNearestDc_Process(csess, tlobj)

	// auth
	case *mtproto.TL_auth_logOut:
		res, err = s.TL_auth_logOut_Process(csess, tlobj)
	case *mtproto.TL_auth_checkPhone:
		res, err = s.TL_auth_checkPhone_Process(csess, tlobj)
	case *mtproto.TL_auth_sendCode:
		res, err = s.TL_auth_sendCode_Process(csess, tlobj)
	case *mtproto.TL_auth_resendCode:
		res, err = s.TL_auth_resendCode_Process(csess, tlobj)
	case *mtproto.TL_auth_signIn:
		res, err = s.TL_auth_signIn_Process(csess, tlobj)

	// account
	case *mtproto.TL_account_updateStatus:
		res, err = s.TL_account_updateStatus_Process(csess, tlobj)

	// contacts
	case *mtproto.TL_contacts_importContacts:
		res, err = s.TL_contacts_importContacts_Process(csess, tlobj)
	case *mtproto.TL_contacts_getContacts:
		res, err = s.TL_contacts_getContacts_Process(csess, tlobj)

	// langpack
	case *mtproto.TL_langpack_getLangPack:
		res, err = s.TL_langpack_getLangPack_Process(csess, tlobj)

	// messages
	case *mtproto.TL_messages_getDialogs:
		res, err = s.TL_messages_getDialogs_Process(csess, tlobj)
	case *mtproto.TL_messages_getPinnedDialogs:
		res, err = s.TL_messages_getPinnedDialogs_Process(csess, tlobj)

	// users
	case *mtproto.TL_users_getFullUser:
		res, err = s.TL_users_getFullUser_Process(csess, tlobj)
	default:
		Log.Debugf("havent implemented yet, TLType = %T", tlobj)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tlobj)
	}

	var reply mtproto.TLObject

	if err != nil {
		Log.Errorf("object = %T, err = %v", tlobj, err)
		rpcerr, ok := err.(*mtproto.TL_rpc_error) // rpc error
		if ok && rpcerr.Get_error_code() != int32(mtproto.TLRpcErrorCodes_NOTRETURN_CLIENT) {
			// err = c.WriteFull(mtproto.GenerateMessageID(), true, rpcerr)
			reply = rpcerr
			return err
		} else {
			return err
		}
	} else {
		reply = &mtproto.TL_rpc_result{
			M_req_msg_id: msgid,
			M_result:     res,
		}
	}

	err = c.WriteFull(mtproto.GenerateMessageID(), true, reply)

	return err
}
