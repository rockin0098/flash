package mtproto

import (
	"fmt"
)

//message msg_id:long seqno:int bytes:int body:Object = Message; // parsed manually
type TL_message2 struct {
	M_classID int32
	M_msg_id  int64
	M_seqno   int32
	M_bytes   int32
	M_body    TLObject
}

func (t *TL_message2) ClassID() int32 {
	return t.M_classID
}

func New_TL_message2() *TL_message2 {
	return &TL_message2{
		M_classID: TL_CLASS_message2,
	}
}

func (t *TL_message2) String() string {
	return fmt.Sprintf("TL_message2")
}

func (t *TL_message2) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	// ec.Int(int32(TL_CLASS_message2))
	ec.Long(t.M_msg_id)
	ec.Int(t.M_seqno)
	ec.Int(t.M_bytes)
	ec.Bytes(t.M_body.Encode())

	return ec.GetBuffer()
}

func (t *TL_message2) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	t.M_msg_id = dc.Long()
	t.M_seqno = dc.Int()
	t.M_bytes = dc.Int()
	t.M_body = dc.TLObject()

	return dc.err
}
