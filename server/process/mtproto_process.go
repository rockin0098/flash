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

	ls := service.LProtoServiceInstance()
	ctx, err := ls.MakeLProtoContext(&lproto.LProtoRequest{
		LID:            guid.GenerateUID(),
		SessionID:      mtp.SessionID(),
		MTProtoRequest: raw,
	})
	if err != nil {
		Log.Error(err)
		return err
	}

	err = mtp.Write(ctx.Response.MTProtoResponse)

	return err
}
