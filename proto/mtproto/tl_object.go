package mtproto

const (
	TL_CLASS_message2      int32 = 1538843921
	TL_CLASS_msg_container int32 = 1945237724
)

func init() {
	tlObjectClassMap[TL_CLASS_message2] = func() TLObject { return New_TL_message2() }
	tlObjectClassMap[TL_CLASS_msg_container] = func() TLObject { return New_TL_msg_container() }
}

type TLObject interface {
	ClassID() int32
	Encode() []byte
	Decode(b []byte) error
}

func NewTLObjectByClassID(classID int32) TLObject {
	m, ok := tlObjectClassMap[classID]
	if !ok {
		return nil
	}
	return m()
}
