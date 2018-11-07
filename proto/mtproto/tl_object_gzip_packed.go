package mtproto

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

// gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TL_gzip_packed struct {
	M_classID     int32
	M_packed_data []byte
}

func (t *TL_gzip_packed) ClassID() int32 {
	return t.M_classID
}

func New_TL_gzip_packed() *TL_gzip_packed {
	return &TL_gzip_packed{
		M_classID: TL_CLASS_gzip_packed,
	}
}

func (t *TL_gzip_packed) String() string {
	return fmt.Sprintf("New_TL_gzip_packed")
}

func (t *TL_gzip_packed) Encode() []byte {
	ec := NewMTPEncodeBuffer(512)

	ec.Int(int32(TL_CLASS_gzip_packed))
	ec.Bytes(t.M_packed_data)

	return ec.GetBuffer()
}

func (t *TL_gzip_packed) Decode(b []byte) error {
	dc := NewMTPDecodeBuffer(b)

	t.M_packed_data = make([]byte, 0, 4096)

	var buf bytes.Buffer
	_, _ = buf.Write(dc.StringBytes())
	gz, _ := gzip.NewReader(&buf)

	b2 := make([]byte, 4096)
	for true {
		n, _ := gz.Read(b2)
		if n <= 0 {
			break
		}
		t.M_packed_data = append(t.M_packed_data, b2[0:n]...)
	}

	return dc.err
}
