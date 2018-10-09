package process

import (
	"github.com/rockin0098/flash/base/guid"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/lproto"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/service"
)

func GateProcess(mtp *mtproto.MTProto) error {

	raw := mtp.Message().(*mtproto.RawMessage)

	request := &lproto.LProtoRequest{
		LID:              guid.GenerateUID(),
		IsDirectResponse: true,
		SessionID:        mtp.SessionID(),
		MTProtoRequest:   raw,
	}

	ls := service.LProtoServiceInstance()
	ctx, err := ls.MakeLProtoContext(request)
	if err != nil {
		Log.Error(err)
		return err
	}

	if request.IsDirectResponse {
		return nil
	} else { // 需要 gate 返回
		// 返回消息
		mtpResp := ctx.Response.MTProtoResponse
		if mtpResp == nil { // 如果返回消息为空, 则表明处理出错, 关闭连接
			mtp.Close()
		} else {
			rmsg := mtproto.NewRawMessage(raw.TransportType, raw.AuthKeyID, raw.QuickAckID)
			rmsg.Decode(ctx.Response.MTProtoResponse.([]byte))
			err = mtp.Write(rmsg)
		}

		return err
	}

	return nil
}
