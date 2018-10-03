package process

import (
	"github.com/rockin0098/flash/base/guid"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/lproto"
	"github.com/rockin0098/flash/proto/mtproto"
)

func MTProtoProcess(mtp *mtproto.MTProto) error {

	raw := mtp.Message().(*mtproto.RawMessage)

	ctx, err := lproto.MakeLProtoContext(&lproto.LProtoRequest{
		LID:            guid.GenerateUID(),
		MTProtoRequest: raw,
	})
	if err != nil {
		Log.Error(err)
		return err
	}

	err = mtp.Write(ctx.Response.MTProtoResponse)

	return err
}
