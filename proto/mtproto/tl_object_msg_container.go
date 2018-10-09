package mtproto

//msg_container#73f1f8dc messages:vector<message> = MessageContainer; // parsed manually
type TL_msg_container struct {
	M_classID  int32
	M_message2 []*TL_message2
}

func (t *TL_msg_container) ClassID() int32 {
	return t.M_classID
}

func (t *TL_msg_container) Set_message2(M_message2 []*TL_message2) {
	t.M_message2 = M_message2
}

func (t *TL_msg_container) Get_message2() []*TL_message2 {
	return t.M_message2
}

func New_TL_msg_container() *TL_msg_container {
	return &TL_msg_container{
		M_classID: TL_CLASS_msg_container,
	}
}

func (t *TL_msg_container) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	ec.Int(int32(TL_CLASS_msg_container))
	// ec.Bytes(t.Get_nonce())
	// ec.Bytes(t.Get_server_nonce())
	// ec.String(t.Get_pq())
	// ec.VectorLong(t.Get_server_public_key_fingerprints())

	return ec.GetBuffer()
}

func (t *TL_msg_container) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	// t.M_nonce = dc.Bytes(16)
	// t.M_server_nonce = dc.Bytes(16)
	// t.M_pq = dc.String()
	// t.M_server_public_key_fingerprints = dc.VectorLong()

	return dc.err
}
