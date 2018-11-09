package tlservice

import (
	"time"

	"github.com/rockin0098/meow/base/crypto"
	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

const (
	kCodeType_None      = 0
	kCodeType_App       = 1
	kCodeType_Sms       = 2
	kCodeType_Call      = 3
	kCodeType_FlashCall = 4
)

const (
	kDBTypeNone   = 0
	kDBTypeCreate = 1
	kDBTypeLoad   = 2
	kDBTypeUpdate = 3
	kDBTypeDelete = 3
)

const (
	kCodeStateNone    = 0
	kCodeStateOk      = 1
	kCodeStateSent    = 2
	kCodeStateSignIn  = 3
	kCodeStateSignUp  = 4
	kCodeStateDeleted = -1
	kCodeStateTimeout = -2
)

// TL_auth_sendCode
func (s *TLService) TL_auth_sendCode_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_auth_sendCode)

	Log.Infof("TL_auth_sendCode = %+v", FormatStruct(tl))

	phone := tl.Get_phone_number()

	mm := model.GetModelManager()
	registered := mm.CheckPhoneExists(phone)

	Log.Infof("registered = %v", registered)

	// 调用短信服务发送短信, 并获取验证码
	//

	code := "66666"

	authPhoneTransaction := &model.AuthPhoneTransaction{
		AuthKeyID:        csess.AuthKeyID,
		PhoneNumber:      phone,
		Code:             code,
		CodeExpired:      int32(time.Now().Unix() + 15*60),
		TransactionHash:  crypto.GenerateStringNonce(16),
		SentCodeType:     kCodeType_App,
		FlashCallPattern: "",
		NextCodeType:     kCodeType_None,
		State:            kCodeStateSent,
		ApiID:            tl.Get_api_id(),
		ApiHash:          tl.Get_api_hash(),
		CreatedTime:      time.Now().Unix(),
	}

	err := mm.ModelAdd(authPhoneTransaction)
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	codeType := &mtproto.TL_auth_sentCodeTypeApp{
		M_length: 5,
	}
	authSentCode := &mtproto.TL_auth_sentCode{
		// M_phone_registered: mtproto.ToBool(registered),
		M_type:            codeType,
		M_phone_code_hash: authPhoneTransaction.TransactionHash,
		// M_next_type:       codeType,
		M_timeout: 60,
	}

	if registered {
		authSentCode.M_phone_registered = mtproto.ToBool(registered)
	}

	return authSentCode, nil
}
