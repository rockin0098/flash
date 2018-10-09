package lproto

import (
	"github.com/rockin0098/flash/proto/mtproto"
)

// 内部节点通信协议

type LContext struct {
	Request  *LProtoRequest
	Response *LProtoResponse
}

type LProtoRequest struct {
	LID              int64
	IsDirectResponse bool
	SessionID        string
	MTProtoRequest   *mtproto.RawMessage
}

type LProtoResponse struct {
	LID             int64
	Error           error
	MTProtoResponse interface{}
}
