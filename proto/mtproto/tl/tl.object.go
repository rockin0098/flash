package tl

type TLObject interface {
	Encode()
}

func NewTLObject() TLObject {
	return nil
}

type TLObjectClassID struct{}

func (t *TLObjectClassID) Encode() {
}
