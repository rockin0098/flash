package service

import (
	"encoding/hex"

	"github.com/rockin0098/flash/server/session"

	"github.com/rockin0098/flash/proto/lproto"

	// . "github.com/rockin0098/flash/base/global"

	. "github.com/rockin0098/flash/base/logger"
)

type LProtoService struct{}

var lprotoService = &LProtoService{}

func LProtoServiceInstance() *LProtoService {
	return lprotoService
}

func (s *LProtoService) MakeLProtoContext(request *lproto.LProtoRequest) (*lproto.LContext, error) {

	ctx := &lproto.LContext{
		Request: request,
	}

	// 此处可以发送到其他节点处理,当前单机版本 直接调用 lproto service 处理
	lresp, err := s.MessageProcess(request)
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	ctx.Response = lresp.(*lproto.LProtoResponse)

	return ctx, nil
}

func (s *LProtoService) MessageProcess(lrequest *lproto.LProtoRequest) (interface{}, error) {

	sessid := lrequest.SessionID
	raw := lrequest.MTProtoRequest
	sess := session.GetSession(sessid)

	Log.Infof("request - sessid=%v, TransportType = %v, AuthKeyID = %v, QuickAckID = %v, \n Req Payload = %v\n",
		sess.SessionID(), raw.TransportType, raw.AuthKeyID, raw.QuickAckID, hex.EncodeToString(raw.Payload))

	mtpresp, err := s.MTProtoMessageProcess(sess, raw)

	Log.Infof("response - sessid = %v, TransportType = %v, AuthKeyID = %v, QuickAckID = %v, \n Req Payload = %v\n",
		sess.SessionID(), raw.TransportType, raw.AuthKeyID, raw.QuickAckID, hex.EncodeToString(mtpresp.([]byte)))

	lresp := &lproto.LProtoResponse{
		LID:             lrequest.LID,
		Error:           err,
		MTProtoResponse: mtpresp,
	}

	return lresp, nil
}
