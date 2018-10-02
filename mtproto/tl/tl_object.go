package tl

type TLBuffer struct {
	data []byte
}

func NewTLBuffer(capability int) *TLBuffer {
	return &TLBuffer{data: make([]byte, 0, capability)}
}

type TLObject interface {
	Encode()
}

func NewTLObject() TLObject {
	return nil
}
