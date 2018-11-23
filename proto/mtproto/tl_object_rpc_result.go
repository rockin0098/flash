package mtproto

import (
	"fmt"
	// . "github.com/rockin0098/meow/base/logger"
)

// rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TL_rpc_result struct {
	M_classID    int32
	M_req_msg_id int64
	M_result     TLObject
}

func (t *TL_rpc_result) ClassID() int32 {
	return t.M_classID
}

func New_TL_rpc_result() *TL_rpc_result {
	return &TL_rpc_result{
		M_classID: TL_CLASS_rpc_result,
	}
}

func (t *TL_rpc_result) String() string {
	return fmt.Sprintf("TL_rpc_result")
}

func (t *TL_rpc_result) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	ec.Int(int32(TL_CLASS_rpc_result))
	ec.Long(t.M_req_msg_id)
	ec.Bytes(t.M_result.Encode())

	return ec.GetBuffer()
}

func (t *TL_rpc_result) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	t.M_req_msg_id = dc.Long()
	t.M_result = dc.TLObject()

	return dc.err
}
