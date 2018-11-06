package service

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/rockin0098/flash/base/global"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/model"
)

type TLService struct{}

var tlService = &TLService{}

func TLServiceInstance() *TLService {
	return tlService
}

func (s *TLService) TLMessageProcess(sess *Session, raw *mtproto.RawMessage) error {

	Log.Info("entering...")

	if raw.AuthKeyID == 0 { // 未加密的消息, 握手协商消息
		reqmsg := &mtproto.UnencryptedMessage{}
		err := reqmsg.Decode(raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return err
		}

		res, err := s.TLUnencryptedMessageProcess(sess, reqmsg)
		if err != nil {
			// Log.Error(err)
			Log.Warn(err)
			return err
		}

		ASSERT(res != nil)
		respmsg := &mtproto.UnencryptedMessage{
			TLObject: res.(mtproto.TLObject),
		}

		err = sess.WriteFull2(&mtproto.RawMessage{Payload: respmsg.Encode()})
		if err != nil {
			Log.Error(err)
			return err
		}

		return nil

	} else { // 加密消息

		Log.Debugf("client authKeyID = %v", raw.AuthKeyID)

		authid := raw.AuthKeyID
		mm := model.GetModelManager()
		akey := mm.GetAuthKeyValueByAuthID(authid)
		if akey == nil {
			return fmt.Errorf("authkey not found by authid=%v", authid)
		}

		reqmsg := &mtproto.EncryptedMessage{
			AuthKeyID: authid,
		}
		err := reqmsg.Decode(akey, raw.Payload[8:])
		if err != nil {
			Log.Error(err)
			return err
		}

		Log.Infof("Decode encrypted message, authid=%v, msgid=%v, salt=%v, seqno=%v, classType=%v",
			reqmsg.AuthKeyID, reqmsg.MessageID, reqmsg.Salt, reqmsg.SeqNo, reqmsg.TLObject)

		clientSessionID := reqmsg.ClientSessionID
		ss := SessionServiceInstance()
		csess, ok := ss.LoadClientSession(clientSessionID)
		if !ok {
			// 记录  session id
			csess = &ClientSession{
				Sess:             sess,
				ClientSessionID:  clientSessionID,
				AuthKeyID:        authid,
				Salt:             reqmsg.Salt,
				FirstMessageID:   reqmsg.MessageID,
				CloseDate:        time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout,
				CloseSessionDate: 0,
				ApiMessages:      make(chan *NetworkApiMessage, 1024),
				UniqueID:         rand.Int63(),
				ClientState:      kStateCreated,
				PendingMessages:  []*PendingMessage{},
				IsUpdates:        false,
				RpcMessages:      []*NetworkApiMessage{},
			}

			ss.NewClientSession(clientSessionID, csess)

		} else {
			csess.Sess = sess
		}

		if !csess.CheckBadServerSalt(authid, reqmsg.MessageID, reqmsg.SeqNo, reqmsg.Salt) {
			Log.Warnf("check bad server salt. authid = %v, class type = %T", authid, reqmsg.TLObject)

			if !ok || reqmsg.MessageID < csess.FirstMessageID {
				Log.Info("reply to client making new client session")

				newSessionCreated := &mtproto.TL_new_session_created{
					M_first_msg_id: reqmsg.MessageID,
					M_unique_id:    csess.UniqueID,
					M_server_salt:  csess.Salt,
				}

				csess.WriteFull(authid, mtproto.GenerateMessageID(), true, newSessionCreated)
			}

			return nil
		}

		_, isContainer := reqmsg.TLObject.(*mtproto.TL_msg_container)
		if !csess.CheckBadMsgNotification(authid, reqmsg.MessageID, reqmsg.SeqNo, isContainer) {
			Log.Warnf("check bad msg notification. authid = %v, class type = %T", authid, reqmsg.TLObject)
			return nil
		}

		err = s.TLEncryptedMessageProcess(csess, reqmsg)
		if err != nil {
			Log.Warn(err)
			return err
		}

		return nil
	}

	return nil
}

func (s *TLService) TLUnencryptedMessageProcess(sess *Session, msg *mtproto.UnencryptedMessage) (interface{}, error) {

	tlobj := msg.TLObject

	Log.Debugf("request - class type = %T, \ntlobj = %s", tlobj, tlobj)

	var res interface{}
	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_req_pq:
		res, err = s.TL_req_pq_Process(sess, msg)
	case *mtproto.TL_req_DH_params:
		res, err = s.TL_req_DH_params_Process(sess, msg)
	case *mtproto.TL_set_client_DH_params:
		res, err = s.TL_set_client_DH_params_Process(sess, msg)
	default:
		Log.Debugf("havent implemented yet, TLType = %T", tl)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tl)
	}

	Log.Debugf("response - class type = %T, \nresp tlobj = %s", tlobj, res)

	return res, err
}

func (s *TLService) TLEncryptedMessageProcess(csess *ClientSession, msg *mtproto.EncryptedMessage) error {

	tlobj := msg.TLObject

	Log.Debugf("request - client sessid = %v, authid = %v, class type = %T, \ntlobj = %v",
		csess.ClientSessionID, msg.AuthKeyID, tlobj, tlobj)

	var err error

	switch tl := tlobj.(type) {
	case *mtproto.TL_ping:
		err = s.TL_ping_Process(csess, msg)
	case *mtproto.TL_invokeWithLayer:
		err = s.TL_invokeWithLayer_Process(csess, msg)
	case *mtproto.TL_initConnection:
		err = s.TL_initConnection_Process(csess, msg)
	case *mtproto.TL_help_getConfig:
		err = s.TL_help_getConfig_Process(csess, msg)
	case *mtproto.TL_msg_container:
		err = s.TL_msg_container_Process(csess, msg)
	case *mtproto.TL_message2:
		err = s.TL_message2_Process(csess, msg)
	case *mtproto.TL_auth_logOut:
		err = s.TL_auth_logOut_Process(csess, msg)
	case *mtproto.TL_langpack_getLangPack:
		err = s.TL_langpack_getLangPack_Process(csess, msg)
	case *mtproto.TL_help_getNearestDc:
		err = s.TL_help_getNearestDc_Process(csess, msg)
	case *mtproto.TL_ping_delay_disconnect:
		err = s.TL_ping_delay_disconnect_Process(csess, msg)
	case *mtproto.TL_auth_checkPhone:
		err = s.TL_auth_checkPhone_Process(csess, msg)
	case *mtproto.TL_msgs_state_req:
		err = s.TL_msgs_state_req_Process(csess, msg)
	case *mtproto.TL_msgs_ack:
		err = s.TL_msgs_ack_Process(csess, msg)
	default:
		Log.Debugf("havent implemented yet, TLType = %T", tl)
		err = fmt.Errorf("havent implemented yet, TLType = %T", tl)
	}

	return err
}

// 消息类型判断
func (s *TLService) TLRpcUpdatesType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_account_registerDevice,
		*mtproto.TL_account_unregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return false

	case *mtproto.TL_upload_saveFilePart,
		*mtproto.TL_upload_saveBigFilePart:

		// upload connection
		return false

	case *mtproto.TL_upload_getFile,
		*mtproto.TL_upload_getWebFile,
		*mtproto.TL_upload_getCdnFile,
		*mtproto.TL_upload_reuploadCdnFile,
		*mtproto.TL_upload_getCdnFileHashes:

		// download
		return false

	case *mtproto.TL_help_getConfig:
		// TODO(@benqi): 可能为TEMP，判断TEMP的规则：
		//  从android client代码看，此连接仅收到help.getConfig消息
	}

	return true
}

func (s *TLService) TLRpcPushType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_account_registerDevice,
		*mtproto.TL_account_unregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return true
	}
	return false
}

func (s *TLService) TLRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_auth_checkedPhone,
		*mtproto.TL_auth_sendCode,
		*mtproto.TL_auth_signIn,
		*mtproto.TL_auth_signUp,
		*mtproto.TL_auth_exportedAuthorization,
		*mtproto.TL_auth_exportAuthorization,
		*mtproto.TL_auth_importAuthorization,
		*mtproto.TL_auth_cancelCode,
		*mtproto.TL_auth_resendCode,
		*mtproto.TL_auth_requestPasswordRecovery,
		*mtproto.TL_auth_checkPassword,
		*mtproto.TL_auth_recoverPassword:

		return true

	case *mtproto.TL_help_getConfig,
		*mtproto.TL_help_getCdnConfig:

		return true

	case *mtproto.TL_langpack_getLanguages,
		*mtproto.TL_langpack_getDifference,
		*mtproto.TL_langpack_getLangPack,
		*mtproto.TL_langpack_getStrings:

		return true

	case *mtproto.TL_account_getPassword,
		*mtproto.TL_account_deleteAccount,
		*mtproto.TL_account_updatePasswordSettings,
		*mtproto.TL_account_getPasswordSettings:

		return true
	}

	return false
}

func (s *TLService) TLRpcUploadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_upload_saveFilePart,
		*mtproto.TL_upload_saveBigFilePart,
		*mtproto.TL_upload_reuploadCdnFile:
		return true
	}
	return false
}

func (s *TLService) TLRpcDownloadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TL_upload_getFile,
		*mtproto.TL_upload_getWebFile,
		*mtproto.TL_upload_getCdnFile,
		*mtproto.TL_upload_getCdnFileHashes:
		return true
	}
	return false
}
