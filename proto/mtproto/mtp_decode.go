package mtproto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"

	. "github.com/rockin0098/flash/base/logger"
)

type MTPDecodeBuffer struct {
	buffer []byte
	off    int
	size   int
	err    error
}

func NewMTPDecodeBuffer(b []byte) *MTPDecodeBuffer {
	return &MTPDecodeBuffer{buffer: b, off: 0, size: len(b), err: nil}
}

func (m *MTPDecodeBuffer) GetError() error {
	return m.err
}

func (m *MTPDecodeBuffer) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buffer[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *MTPDecodeBuffer) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buffer[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *MTPDecodeBuffer) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buffer[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *MTPDecodeBuffer) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buffer[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *MTPDecodeBuffer) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buffer[m.off:m.off+size])
	m.off += size
	return x
}

func (m *MTPDecodeBuffer) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.buffer[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.buffer[m.off]) | int(m.buffer[m.off+1])<<8 | int(m.buffer[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buffer[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding

	return x
}

func (m *MTPDecodeBuffer) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *MTPDecodeBuffer) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *MTPDecodeBuffer) VectorInt() []int32 {

	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(TL_CLASS_vector) {
		m.err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *MTPDecodeBuffer) VectorLong() []int64 {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(TL_CLASS_vector) {
		m.err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *MTPDecodeBuffer) VectorString() []string {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(TL_CLASS_vector) {
		m.err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *MTPDecodeBuffer) Bool() bool {
	constructor := m.Int()
	if m.err != nil {
		return false
	}
	switch constructor {
	case int32(TL_CLASS_boolTrue):
		return true
	case int32(TL_CLASS_boolFalse):
		return false
	}
	return false
}

// func (m *MTPDecodeBuffer) TLObject() TLObject {
// 	// constructor := m.Int()
// 	// if m.err != nil {
// 	// 	return nil
// 	// }
// 	// switch constructor {
// 	// case int32(TL_CLASS_boolTrue):
// 	// 	return true
// 	// case int32(TL_CLASS_boolFalse):
// 	// 	return false
// 	// }
// 	// return false

// 	return nil
// }

/*
func (m *MTPDecodeBuffer) Vector() []TLObject {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(TLConstructor_CRC32_vector) {
		m.err = fmt.Errorf("DecodeVector: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TLObject, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}
*/

func (m *MTPDecodeBuffer) TLObjectClassID() int32 {
	classID := m.Int()
	if m.err != nil {
		return 0
	}

	className := TL_CLASS_NAME[classID]
	Log.Infof("TLObject, classID: %x, className: %v", uint32(classID), className)

	return classID
}

func (m *MTPDecodeBuffer) TLObject() TLObject {
	classID := m.Int()
	if m.err != nil {
		return nil
	}

	className := TL_CLASS_NAME[classID]
	Log.Debugf("TLObject - classID: %x, className: %v", uint32(classID), className)

	tlo := NewTLObjectByClassID(classID)
	if tlo == nil {
		Log.Errorf("Can't find registed classID: %v", classID)
		return nil
	}

	tlo.Decode(m.buffer[m.off:])

	return tlo
}
