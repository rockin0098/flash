package service

import (
	"encoding/hex"

	// . "github.com/rockin0098/flash/base/global"

	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto"
)

type MTProtoService struct{}

var mtprotoService = &MTProtoService{}

func MTProtoServiceInstance() *MTProtoService {
	return mtprotoService
}

func (s *MTProtoService) MessageProcess(raw *mtproto.RawMessage) (interface{}, error) {

	Log.Infof("entering....... TransportType = %v, AuthKeyID = %v, QuickAckID = %v, Payload = \n%v\n",
		raw.TransportType, raw.AuthKeyID, raw.QuickAckID, hex.EncodeToString(raw.Payload))

	if raw.AuthKeyID == 0 { // 未加密的消息, 握手消息
		// unencryptedMessage := &mtproto.UnencryptedMessage{}
		// err := unencryptedMessage.Decode(raw.Payload)
		// if err != nil {
		// 	Log.Error(err)
		// 	return nil, err
		// }

	} else {

	}

	return nil, nil
}
