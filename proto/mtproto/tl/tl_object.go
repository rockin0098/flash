package tl

type TLObject interface {
	Encode()
}

func NewTLObject() TLObject {
	return nil
}
