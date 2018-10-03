package lproto

import (
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/service"
)

// 内部节点通信协议

type LContext struct {
	Request  *LProtoRequest
	Response *LProtoResponse
}

type LProtoRequest struct {
	LID            int64
	MTProtoRequest *mtproto.RawMessage
}

type LProtoResponse struct {
	LID             int64
	Error           error
	MTProtoResponse interface{}
}

func MakeLProtoContext(request *LProtoRequest) (*LContext, error) {

	ctx := &LContext{
		Request: request,
	}

	// 此处可以发送到其他节点处理,当前单机版本 直接调用 mtproto service 处理
	mts := service.MTProtoServiceInstance()
	mtresponse, err := mts.MessageProcess(request.MTProtoRequest)
	if err != nil {
		return nil, err
	}

	ctx.Response = &LProtoResponse{
		LID:             request.LID,
		Error:           err,
		MTProtoResponse: mtresponse,
	}

	return ctx, nil
}
