package mtproto

import (
	"fmt"

	"github.com/rockin0098/flash/mtproto/tl"
)

type TUnencryptedMessage struct {
	NeedAck   bool
	MessageID int64
	TLObject  tl.TLObject
}

func (m *TUnencryptedMessage) Encode() ([]byte, error) {

	return nil, nil
}

func (m *TUnencryptedMessage) Decode(b []byte) error {

	dc := NewMTDecodeBuffer(b)
	// m.authKeyId = dc.Long()
	m.MessageID = dc.Long()

	// glog.Info("messageId:", m.messageId)
	// mod := m.messageId & 3
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

	// m.Object = dc.Object()
	// if m.Object == nil {
	// 	return fmt.Errorf("decode object is nil")
	// }

	// proto.Message()
	// glog.Info("Recved object: ", m.Object.(proto.Message).String())
	return dc.err
}

type TEncryptedMessage struct {
	authKeyID int64
	NeedAck   bool

	msgKey    []byte
	Salt      int64
	SessionID int64
	MessageID int64
	SeqNo     int32
	TLObject  tl.TLObject
}

func (m *TEncryptedMessage) Encode() ([]byte, error) {

	return nil, nil
}

func (m *TEncryptedMessage) Decode(authKey []byte, b []byte) error {

	return nil
}
