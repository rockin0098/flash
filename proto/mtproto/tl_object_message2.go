package mtproto

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

func (t *TL_message2) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	ec.Int(int32(TL_CLASS_message2))
	// ec.Bytes(t.Get_nonce())
	// ec.Bytes(t.Get_server_nonce())
	// ec.String(t.Get_pq())
	// ec.VectorLong(t.Get_server_public_key_fingerprints())

	return ec.GetBuffer()
}

func (t *TL_message2) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	// t.M_nonce = dc.Bytes(16)
	// t.M_server_nonce = dc.Bytes(16)
	// t.M_pq = dc.String()
	// t.M_server_public_key_fingerprints = dc.VectorLong()

	return dc.err
}
