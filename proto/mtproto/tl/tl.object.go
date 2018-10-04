package tl

type TLObject interface {
	Encode() []byte
	Decode(b []byte)
}

func NewTLObject() TLObject {
	return nil
}

type TLObjectClassID struct{}

func (t *TLObjectClassID) Encode() {
	test := TL_CLASS_req_pq
	_ = test
}

func NewTLObjectByClassID(classID int32) TLObject {
	m, ok := tlObjectClassMap[classID]
	if !ok {
		return nil
	}
	return m()
}
