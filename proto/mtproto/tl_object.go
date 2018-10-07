package mtproto

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
