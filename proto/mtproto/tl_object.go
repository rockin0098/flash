package mtproto

type TLObject interface {
	Encode() []byte
	Decode(b []byte)
}

func NewTLObject() TLObject {
	return nil
}

func NewTLObjectByClassID(classID int32) TLObject {
	m, ok := tlObjectClassMap[classID]
	if !ok {
		return nil
	}
	return m()
}

// for testing
type TLObjectClassID struct {
	ClassID int32
}

func (t *TLObjectClassID) Encode() []byte {
	return nil
}

func (t *TLObjectClassID) Decode(b []byte) {}
