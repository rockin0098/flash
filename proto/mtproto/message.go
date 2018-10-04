package mtproto

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/rockin0098/flash/base/crypto"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto/tl"
)

type MTProtoMessage interface {
	Encode() ([]byte, error)
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

func (m *RawMessage) Encode() ([]byte, error) {

	return nil, nil
}

func (m *RawMessage) Decode(b []byte) error {
	m.Payload = b
	return nil
}

type UnencryptedMessage struct {
	NeedAck   bool
	MessageID int64
	TLObject  tl.TLObject
}

func (m *UnencryptedMessage) Encode() ([]byte, error) {

	return nil, nil
}

func (m *UnencryptedMessage) Decode(b []byte) error {

	Log.Infof("UnencryptedMessage decode = %v", hex.EncodeToString(b))

	dc := NewMTDecodeBuffer(b)
	// m.authKeyId = dc.Long()
	m.MessageID = dc.Long()

	Log.Info("messageId:", m.MessageID)
	Log.Infof("dc: buffer:%v, off:%v, size:%v, err:%v", hex.EncodeToString(dc.buffer), dc.off, dc.size, dc.err)
	// mod := m.messageID & 3
	// if mod != 1 && mod != 3 {
	// 	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	messageLen := dc.Int()
	if messageLen < 4 {
		return fmt.Errorf("message len(%d) < 4", messageLen)
	}
	// glog.Info("messageLen:", m.messageId)

	if int(messageLen) != dc.size-12 {
		return fmt.Errorf("message len: %d (need %d)", messageLen, dc.size-12)
	}

	m.TLObject = dc.TLObject()
	// if m.TLObject == nil {
	// 	return fmt.Errorf("decode object is nil")
	// }

	// proto.Message()
	// glog.Info("Recved object: ", m.Object.(proto.Message).String())
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
	TLObject  tl.TLObject
}

func (m *EncryptedMessage) Encode() ([]byte, error) {

	return nil, nil
}

func (m *EncryptedMessage) Decode(authKey []byte, b []byte) error {

	return nil
}

// encrypted stream
type AesCTR128Stream struct {
	reader  *bufio.Reader
	writer  io.Writer
	encrypt *crypto.AesCTR128Encrypt
	decrypt *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(reader *bufio.Reader, writer io.Writer, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		reader:  reader,
		writer:  writer,
		decrypt: d,
		encrypt: e,
	}
}

func (this *AesCTR128Stream) Read(p []byte) (int, error) {
	Log.Debug("entering...")
	n, err := this.reader.Read(p)
	if err == nil {
		this.decrypt.Encrypt(p[:n])
		return n, nil
	}
	return n, err
}

func (this *AesCTR128Stream) Write(p []byte) (int, error) {
	this.encrypt.Encrypt(p[:])
	return this.writer.Write(p)
}
