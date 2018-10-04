package mtproto

import (
	"encoding/binary"
	"math"
	"math/big"
)

type MTEncodeBuffer struct {
	buffer []byte
}

func NewMTEncodeBuffer(capability int) *MTEncodeBuffer {
	return &MTEncodeBuffer{buffer: make([]byte, 0, capability)}
}

func (e *MTEncodeBuffer) GetBuf() []byte {
	return e.buffer
}

func (e *MTEncodeBuffer) Int16(s int16) {
	e.buffer = append(e.buffer, 0, 0)
	binary.LittleEndian.PutUint16(e.buffer[len(e.buffer)-2:], uint16(s))
}

func (e *MTEncodeBuffer) UInt16(s uint16) {
	e.buffer = append(e.buffer, 0, 0)
	binary.LittleEndian.PutUint16(e.buffer[len(e.buffer)-2:], s)
}

func (e *MTEncodeBuffer) Int(s int32) {
	e.buffer = append(e.buffer, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buffer[len(e.buffer)-4:], uint32(s))
}

func (e *MTEncodeBuffer) UInt(s uint32) {
	e.buffer = append(e.buffer, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buffer[len(e.buffer)-4:], s)
}

func (e *MTEncodeBuffer) Long(s int64) {
	e.buffer = append(e.buffer, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buffer[len(e.buffer)-8:], uint64(s))
}

func (e *MTEncodeBuffer) Double(s float64) {
	e.buffer = append(e.buffer, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buffer[len(e.buffer)-8:], math.Float64bits(s))
}

func (e *MTEncodeBuffer) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *MTEncodeBuffer) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *MTEncodeBuffer) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buffer = append(e.buffer, res...)
}

func (e *MTEncodeBuffer) Bytes(s []byte) {
	e.buffer = append(e.buffer, s...)
}

func (e *MTEncodeBuffer) VectorInt(v []int32) {
	// x := make([]byte, 4+4+len(v)*4)
	// var c = int32(TLConstructor_CRC32_vector)
	// binary.LittleEndian.PutUint32(x, uint32(c))
	// binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	// i := 8
	// for _, v := range v {
	// 	binary.LittleEndian.PutUint32(x[i:], uint32(v))
	// 	i += 4
	// }
	// e.buffer = append(e.buffer, x...)
}

func (e *MTEncodeBuffer) VectorLong(v []int64) {
	// x := make([]byte, 4+4+len(v)*8)
	// var c = int32(TLConstructor_CRC32_vector)
	// binary.LittleEndian.PutUint32(x, uint32(c))
	// binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	// i := 8
	// for _, v := range v {
	// 	binary.LittleEndian.PutUint64(x[i:], uint64(v))
	// 	i += 8
	// }
	// e.buffer = append(e.buffer, x...)
}

func (e *MTEncodeBuffer) VectorString(v []string) {
	// x := make([]byte, 8)
	// var c = int32(TLConstructor_CRC32_vector)
	// binary.LittleEndian.PutUint32(x, uint32(c))
	// binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	// e.buffer = append(e.buffer, x...)
	// for _, v := range v {
	// 	e.String(v)
	// }
}

/*
func (e *MTEncodeBuffer) Vector(v []TLObject) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, CRC32_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buffer = append(e.buffer, x...)
	for _, v := range v {
		e.buffer = append(e.buffer, v.encode()...)
	}
}
*/
