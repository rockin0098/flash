package mtproto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/rockin0098/meow/base/crypto"
)

func GenerateMessageID() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	messageID := ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
	for {
		//rpc_response
		if (messageID % 4) != 1 {
			messageID += 1
		} else {
			break
		}

		/****************************
		 * // rpc_request
		 * if (messageID % 4) != 3 {
		 * 	messageID += 1
		 * } else {
		 * 	break
		 * }
		 */
	}

	return messageID
}

/*
	uint32_t x = incoming ? 8 : 0;
	static uint8_t sha[68];

	SHA256_Init(&sha256Ctx);
	SHA256_Update(&sha256Ctx, messageKey, 16);
	SHA256_Update(&sha256Ctx, authKey + x, 36);
	SHA256_Final(sha, &sha256Ctx);

	SHA256_Init(&sha256Ctx);
	SHA256_Update(&sha256Ctx, authKey + 40 + x, 36);
	SHA256_Update(&sha256Ctx, messageKey, 16);
	SHA256_Final(sha + 32, &sha256Ctx);

	memcpy(result, sha, 8);
	memcpy(result + 8, sha + 32 + 8, 16);
	memcpy(result + 8 + 16, sha + 24, 8);

	memcpy(result + 32, sha + 32, 8);
	memcpy(result + 32 + 8, sha + 8, 16);
	memcpy(result + 32 + 8 + 16, sha + 32 + 24, 8);

*/
func generateMessageKey(msgKey, authKey []byte, incoming bool) (aesKey, aesIV []byte) {
	var x = 0
	if incoming {
		x = 8
	}

	switch MTPROTO_VERSION {
	case "2.0":
		t_a := make([]byte, 0, 52)
		t_a = append(t_a, msgKey[:16]...)
		t_a = append(t_a, authKey[x:x+36]...)
		sha256_a := crypto.Sha256Digest(t_a)

		t_b := make([]byte, 0, 52)
		t_b = append(t_b, authKey[40+x:40+x+36]...)
		t_b = append(t_b, msgKey[:16]...)
		sha256_b := crypto.Sha256Digest(t_b)

		aesKey = make([]byte, 0, 32)
		aesKey = append(aesKey, sha256_a[:8]...)
		aesKey = append(aesKey, sha256_b[8:8+16]...)
		aesKey = append(aesKey, sha256_a[24:24+8]...)

		aesIV = make([]byte, 0, 32)
		aesIV = append(aesIV, sha256_b[:8]...)
		aesIV = append(aesIV, sha256_a[8:8+16]...)
		aesIV = append(aesIV, sha256_b[24:24+8]...)

	default:
		aesKey = make([]byte, 0, 32)
		aesIV = make([]byte, 0, 32)
		t_a := make([]byte, 0, 48)
		t_b := make([]byte, 0, 48)
		t_c := make([]byte, 0, 48)
		t_d := make([]byte, 0, 48)

		t_a = append(t_a, msgKey...)
		t_a = append(t_a, authKey[x:x+32]...)

		t_b = append(t_b, authKey[32+x:32+x+16]...)
		t_b = append(t_b, msgKey...)
		t_b = append(t_b, authKey[48+x:48+x+16]...)

		t_c = append(t_c, authKey[64+x:64+x+32]...)
		t_c = append(t_c, msgKey...)

		t_d = append(t_d, msgKey...)
		t_d = append(t_d, authKey[96+x:96+x+32]...)

		sha1_a := crypto.Sha1Digest(t_a)
		sha1_b := crypto.Sha1Digest(t_b)
		sha1_c := crypto.Sha1Digest(t_c)
		sha1_d := crypto.Sha1Digest(t_d)

		aesKey = append(aesKey, sha1_a[0:8]...)
		aesKey = append(aesKey, sha1_b[8:8+12]...)
		aesKey = append(aesKey, sha1_c[4:4+12]...)

		aesIV = append(aesIV, sha1_a[8:8+12]...)
		aesIV = append(aesIV, sha1_b[0:8]...)
		aesIV = append(aesIV, sha1_c[16:16+4]...)
		aesIV = append(aesIV, sha1_d[0:8]...)
	}

	return
}

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
	// AuthKeyID int64 // must be 0
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
	AuthKeyID       int64
	NeedAck         bool
	MsgKey          []byte
	Salt            int64
	ClientSessionID int64 // client session id
	MessageID       int64
	SeqNo           int32
	TLObject        TLObject
}

func (m *EncryptedMessage) Encode(authKeyID int64, authKey []byte) []byte {

	objData := m.TLObject.Encode()
	var additional_size = (32 + len(objData)) % 16
	if additional_size != 0 {
		additional_size = 16 - additional_size
	}
	if MTPROTO_VERSION == "2.0" && additional_size < 12 {
		additional_size += 16
	}

	x := NewMTPEncodeBuffer(32 + len(objData) + additional_size)
	// x.Long(authKeyId)
	// msgKey := make([]byte, 16)
	// x.Bytes(msgKey)
	x.Long(m.Salt)
	x.Long(m.ClientSessionID)
	if m.MessageID == 0 {
		m.MessageID = GenerateMessageID()
	}
	x.Long(m.MessageID)
	x.Int(m.SeqNo)
	x.Int(int32(len(objData)))
	x.Bytes(objData)
	x.Bytes(crypto.GenerateNonce(additional_size))

	encryptedData, _ := m.encrypt(authKey, x.buffer)
	x2 := NewMTPEncodeBuffer(56 + len(objData) + additional_size)
	x2.Long(authKeyID)
	x2.Bytes(m.MsgKey)
	x2.Bytes(encryptedData)

	Log.Infof("encrypted message encoding, authid=%v, MessageID=%v, classType=%T ", authKeyID, m.MessageID, m.TLObject)

	return x2.buffer
}

func (m *EncryptedMessage) Decode(authKey []byte, b []byte) error {

	msgKey := b[:16]
	// aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	// x, err := doAES256IGEdecrypt(b[16:], aesKey, aesIV)

	x, err := m.decrypt(msgKey, authKey, b[16:])
	if err != nil {
		return err
	}

	Log.Debugf("decrypted message = \n%v", hex.EncodeToString(x))

	dc := NewMTPDecodeBuffer(x)

	m.Salt = dc.Long()            // salt
	m.ClientSessionID = dc.Long() // session_id
	m.MessageID = dc.Long()

	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	//	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	m.SeqNo = dc.Int()
	messageLen := dc.Int()
	if int(messageLen) > dc.size-32 {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messagxeLen, dc.size-32)
	}

	// Log.Infof("salt: %d, sessionId: %d, messageId: %d, seqNo: %d, messageLen: %d, buffer = \n%v", m.Salt, m.ClientSessionID, m.MessageID, m.SeqNo, messageLen, hex.Dump(dc.buffer[dc.off:]))
	Log.Infof("salt: %d, sessionId: %d, messageId: %d, seqNo: %d, messageLen: %d", m.Salt, m.ClientSessionID, m.MessageID, m.SeqNo, messageLen)
	m.TLObject = dc.TLObject()
	if m.TLObject == nil {
		return fmt.Errorf("Decode object is nil")
	}

	// glog.Info("Recved object: ", m.Object.String())

	return nil
}

func (m *EncryptedMessage) decrypt(msgKey, authKey, data []byte) ([]byte, error) {

	var dataLen = uint32(len(data))
	// 创建aesKey, aesIV
	aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	d := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := d.Decrypt(data)
	if err != nil {
		Log.Error("descrypted data error: ", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	messageLen := binary.LittleEndian.Uint32(x[28:])
	// glog.Info("descrypt - messageLen = ", messageLen)
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("descrypted data error: Wrong message length %d", messageLen)
		Log.Error(err)
		return nil, err
	}

	sha256MsgKey := make([]byte, 96)
	switch MTPROTO_VERSION {
	case "2.0":
		t_d := make([]byte, 0, 32+dataLen)
		t_d = append(t_d, authKey[88:88+32]...)
		t_d = append(t_d, x[:dataLen]...)
		copy(sha256MsgKey, crypto.Sha256Digest(t_d))
	default:
		copy(sha256MsgKey[4:], crypto.Sha1Digest(x))
	}

	if !bytes.Equal(sha256MsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("descrypted data error: (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(data[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			m.AuthKeyID,
			hex.EncodeToString(authKey[88:88+32]),
			hex.EncodeToString(sha256MsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		Log.Error(err)
		return nil, err
	}

	return x, nil
}

func (m *EncryptedMessage) encrypt(authKey []byte, data []byte) ([]byte, error) {
	message_key := make([]byte, 32)
	switch MTPROTO_VERSION {
	case "2.0":
		t_d := make([]byte, 0, 32+len(data))
		t_d = append(t_d, authKey[88+8:88+8+32]...)
		t_d = append(t_d, data...)
		copy(message_key, crypto.Sha256Digest(t_d))
	default:
		copy(message_key[4:], crypto.Sha1Digest(data))
	}

	// copy(message_key[8:], )
	// memcpy(p_data + 8, message_key + 8, 16);

	aesKey, aesIV := generateMessageKey(message_key[8:8+16], authKey, true)
	e := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(data)
	if err != nil {
		Log.Error("Encrypt data error: ", err)
		return nil, err
	}

	m.MsgKey = message_key[8 : 8+16]
	return x, nil
}

/////////////////// encrypted message end ////////////////

// encrypted stream
type AesCTR128Stream struct {
	reader    io.Reader
	writer    io.Writer
	writeChan chan interface{}
	encrypt   *crypto.AesCTR128Encrypt
	decrypt   *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(reader io.Reader, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		reader:  reader,
		decrypt: d,
		encrypt: e,
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

func (this *AesCTR128Stream) Encrypt(p []byte) {
	this.encrypt.Encrypt(p[:])
}
