package mtproto

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/rockin0098/flash/base/crypto"
	. "github.com/rockin0098/flash/base/logger"
)

type MTProtoMessage interface {
	Encode() []byte
	Decode(b []byte) error
}

type RawMessage struct {
	TransportType int
	AuthKeyID     int64 // 由原始数据解压获得
	QuickAckID    int32 // EncryptedMessage，则可能存在

	// 原始数据
	Payload []byte
}

func NewRawMessage(transportType int, authKeyID int64, quickAckID int32) *RawMessage {
	return &RawMessage{
		TransportType: transportType,
		AuthKeyID:     authKeyID,
		QuickAckID:    quickAckID,
	}
}

func (m *RawMessage) Encode() []byte {
	return m.Payload
}

func (m *RawMessage) Decode(b []byte) error {
	m.Payload = b
	return nil
}

type UnencryptedMessage struct {
	NeedAck   bool
	MessageID int64
	TLObject  TLObject
}

func (m *UnencryptedMessage) Encode() []byte {

	x := NewMTPEncodeBuffer(512)
	x.Long(0)
	m.MessageID = GenerateMessageID()
	x.Long(m.MessageID)

	if m.TLObject == nil {
		x.Int(0)
	} else {
		b := m.TLObject.Encode()
		x.Int(int32(len(b)))
		x.Bytes(b)
	}

	return x.buffer
}

func (m *UnencryptedMessage) Decode(b []byte) error {

	Log.Infof("UnencryptedMessage decode = %v", hex.EncodeToString(b))

	dc := NewMTPDecodeBuffer(b)
	// m.authKeyId = dc.Long()
	m.MessageID = dc.Long()

	messageLen := dc.Int()
	if messageLen < 4 {
		return fmt.Errorf("message len(%d) < 4", messageLen)
	}

	if int(messageLen) != dc.size-12 {
		return fmt.Errorf("message len: %d (need %d)", messageLen, dc.size-12)
	}

	m.TLObject = dc.TLObject()

	return dc.err
}

type EncryptedMessage struct {
	AuthKeyID int64
	NeedAck   bool
	MsgKey    []byte
	Salt      int64
	SessionID int64
	MessageID int64
	SeqNo     int32
	TLObject  TLObject
}

func (m *EncryptedMessage) Encode() []byte {

	return nil
}

func (m *EncryptedMessage) Decode(authKey []byte, b []byte) error {

	return nil
}

// encrypted stream
type AesCTR128Stream struct {
	reader    io.Reader
	writer    io.Writer
	writeChan chan interface{}
	encrypt   *crypto.AesCTR128Encrypt
	decrypt   *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(reader io.Reader, writer io.Writer, writeChan chan interface{}, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		reader:    reader,
		writer:    writer,
		writeChan: writeChan,
		decrypt:   d,
		encrypt:   e,
	}
}

func (this *AesCTR128Stream) Read(p []byte) (int, error) {
	// Log.Debug("entering...")
	n, err := this.reader.Read(p)
	if err == nil {
		this.decrypt.Encrypt(p[:n])
		return n, nil
	}
	return n, err
}

func (this *AesCTR128Stream) Write(p []byte) {
	this.encrypt.Encrypt(p[:])

	this.writeChan <- p
}
