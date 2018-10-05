package mtproto

type TLObject interface {
	Encode() []byte
	Decode(b []byte)
}

func NewTLObjectByClassID(classID int32) TLObject {
	m, ok := tlObjectClassMap[classID]
	if !ok {
		return nil
	}
	return m()
}
