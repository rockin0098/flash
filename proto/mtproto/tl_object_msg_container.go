package mtproto

import (
	"fmt"
)

//msg_container#73f1f8dc messages:vector<message> = MessageContainer; // parsed manually
type TL_msg_container struct {
	M_classID   int32
	M_message2s []*TL_message2
}

func (t *TL_msg_container) ClassID() int32 {
	return t.M_classID
}

func New_TL_msg_container() *TL_msg_container {
	return &TL_msg_container{
		M_classID: TL_CLASS_msg_container,
	}
}

func (t *TL_msg_container) String() string {
	return fmt.Sprintf("TL_msg_container")
}

func (t *TL_msg_container) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	ec.Int(int32(TL_CLASS_msg_container))
	ec.Int(int32(len(t.M_message2s)))
	for _, m := range t.M_message2s {
		ec.Bytes(m.Encode())
	}

	return ec.GetBuffer()
}

func (t *TL_msg_container) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	len := dc.Int()
	Log.Info("TLMsgContainer: messages len: ", len)
	spos := dc.off
	for i := 0; i < int(len); i++ {

		message2 := &TL_message2{}
		err := message2.Decode(b[spos:])
		if err != nil {
			Log.Error("Decode message2 error: ", err)
			return err
		}

		Log.Infof("TLMsgContainer: messages[%d]: %v, body = %v", i, message2, message2.M_body)

		spos = spos + 8 + 4 + 4 + int(message2.M_bytes)

		t.M_message2s = append(t.M_message2s, message2)
	}
	return dc.err
}
