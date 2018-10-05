package mtproto

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
	"net"

	"github.com/rockin0098/flash/base/crypto"
	// . "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

// 协议版本
const (
	MTPROTO_ABRIDGED_VERSION     = 1 // 删节版本
	MTPROTO_INTERMEDIATE_VERSION = 2 // 中间版本
	MTPROTO_FULL_VERSION         = 3 // 完整版本
	MTPROTO_APP_VERSION          = 4 // Android等当前客户端使用版本
)

// Transport类型，不支持UDP
const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
	TRANSPORT_UDP  = 3 // UDP
)

const (
	// Tcp Transport
	MTPROTO_ABRIDGED_FLAG     = 0xef
	MTPROTO_INTERMEDIATE_FLAG = 0xeeeeeeee

	// Http Transport
	HTTP_HEAD_FLAG   = 0x44414548
	HTTP_POST_FLAG   = 0x54534f50
	HTTP_GET_FLAG    = 0x20544547
	HTTP_OPTION_FLAG = 0x4954504f

	VAL2_FLAG = 0x00000000
)

type MTProto struct {
	reader     *bufio.Reader
	writer     io.Writer
	respChan   chan interface{}
	message    MTProtoMessage
	e          *crypto.AesCTR128Encrypt
	d          *crypto.AesCTR128Encrypt
	remoteAddr net.Addr
	localAddr  net.Addr
	sessionID  string
	// AuthKey []byte
}

func NewMTProto(reader *bufio.Reader, writer io.Writer, remoteAddr net.Addr, localAddr net.Addr, sessid string, respChan chan interface{}) *MTProto {
	return &MTProto{
		reader:     reader,
		writer:     writer,
		respChan:   respChan,
		remoteAddr: remoteAddr,
		localAddr:  localAddr,
		sessionID:  sessid,
	}
}

func (s *MTProto) RemoteAddr() net.Addr {
	return s.remoteAddr
}

func (s *MTProto) LocalAddr() net.Addr {
	return s.localAddr
}

func (s *MTProto) Message() MTProtoMessage {
	return s.message
}

func (s *MTProto) SessionID() string {
	return s.sessionID
}

/////////////////////////////////////////////////////////////////////////////////////
/**
  Android client code:

	RAND_bytes(bytes, 64);
	uint32_t val = (bytes[3] << 24) | (bytes[2] << 16) | (bytes[1] << 8) | (bytes[0]);
	uint32_t val2 = (bytes[7] << 24) | (bytes[6] << 16) | (bytes[5] << 8) | (bytes[4]);
	if (bytes[0] != 0xef &&
		val != 0x44414548 &&
		val != 0x54534f50 &&
		val != 0x20544547 &&
		val != 0x4954504f &&
		val != 0xeeeeeeee &&
		val2 != 0x00000000) {
		bytes[56] = bytes[57] = bytes[58] = bytes[59] = 0xef;
		break;
	}
*/
func (s *MTProto) Read() error {

	br := s.reader
	// var b_0_1 = make([]byte, 1)
	b_0_1, err := br.Peek(1)
	if err != nil {
		Log.Errorf("peek b_0_1 error: %v", err)
		return err
	}
	Log.Debug("b_0_1 : ", hex.EncodeToString(b_0_1))

	if b_0_1[0] == MTPROTO_ABRIDGED_FLAG {
		Log.Debug("mtproto abridged version.")
		br.Discard(1)
		return s.ReadMTProtoAbriaged()
	}

	// not abridged version, we'll lookup codec!
	b_1_3, err := br.Peek(4)
	if err != nil {
		Log.Errorf("peek b_1_3 error: %v", err)
		return err
	}
	Log.Debug("1.b_1_3 : ", hex.EncodeToString(b_1_3))

	b_1_3 = b_1_3[1:4]

	Log.Debug("2.b_1_3 : ", hex.EncodeToString(b_1_3))

	// first uint32
	val := (uint32(b_1_3[2]) << 24) | (uint32(b_1_3[1]) << 16) | (uint32(b_1_3[0]) << 8) | (uint32(b_0_1[0]))
	Log.Debugf("first uint32 : %08x", val)
	if val == HTTP_HEAD_FLAG || val == HTTP_POST_FLAG || val == HTTP_GET_FLAG || val == HTTP_OPTION_FLAG {
		// http 协议
		Log.Info("mtproto http.")
		return s.ReadMTProtoHttp()
	}

	// an intermediate version
	if val == MTPROTO_INTERMEDIATE_FLAG {
		//glog.Warning("MTProtoProxyCodec - mtproto intermediate version, impl in the future!!")
		//return nil, errors.New("mtproto intermediate version not impl!!")
		Log.Info("mtproto intermediate version.")
		br.Discard(4)
		return s.ReadMTProtoIntermidiate()
	}

	// recv 4~64 bytes
	// var b_4_60 = make([]byte, 60)
	b_4_60, err := br.Peek(64)
	if err != nil {
		Log.Errorf("peek b_4_60 error: %v", err)
		return err
	}
	b_4_60 = b_4_60[4:64]
	val2 := (uint32(b_4_60[3]) << 24) | (uint32(b_4_60[2]) << 16) | (uint32(b_4_60[1]) << 8) | (uint32(b_4_60[0]))
	if val2 == VAL2_FLAG {
		Log.Info("mtproto full version.")
		return s.ReadMTProtoFull()
	}

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = b_4_60[51-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		Log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}
	s.e = e

	d, err := crypto.NewAesCTR128Encrypt(b_4_60[4:36], b_4_60[36:52])
	if err != nil {
		Log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}
	s.d = d

	d.Encrypt(b_0_1)
	d.Encrypt(b_1_3)
	d.Encrypt(b_4_60)

	if b_4_60[52] != 0xef && b_4_60[53] != 0xef && b_4_60[54] != 0xef && b_4_60[55] != 0xef {
		Log.Errorf("first 56~59 byte != 0xef")
		return errors.New("mtproto buf[56:60]'s byte != 0xef!!")
	}

	Log.Info("first_bytes_64: ", hex.EncodeToString(b_0_1), hex.EncodeToString(b_1_3), hex.EncodeToString(b_4_60))
	br.Discard(64)

	return s.ReadMTProtoApp()
}

func (s *MTProto) ReadMTProtoAbriaged() error {
	Log.Debug("entering...")

	return nil
}

func (s *MTProto) ReadMTProtoIntermidiate() error {
	Log.Debug("entering...")

	return nil
}

func (s *MTProto) ReadMTProtoFull() error {
	Log.Debug("entering...")

	return nil
}

func (s *MTProto) ReadMTProtoHttp() error {
	Log.Debug("entering...")

	return nil
}

func (s *MTProto) ReadMTProtoApp() error {
	Log.Debug("entering...")

	var size int
	var n int
	var err error

	stream := NewAesCTR128Stream(s.reader, s.writer, s.d, s.e)
	b := make([]byte, 1)
	n, err = io.ReadFull(stream, b)
	if err != nil {
		Log.Error(err)
		return err
	}

	Log.Info("first_byte: ", hex.EncodeToString(b[:1]))
	needAck := bool(b[0]>>7 == 1)
	_ = needAck

	b[0] = b[0] & 0x7f
	Log.Info("first_byte2: ", hex.EncodeToString(b[:1]))

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		Log.Info("size1: ", size)
		if size == 0 {
			return nil
		}
	} else {
		Log.Info("first_byte2: ", hex.EncodeToString(b[:1]))
		b2 := make([]byte, 3)
		n, err = io.ReadFull(stream, b2)
		if err != nil {
			return err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
		Log.Info("size2: ", size)
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(stream, buf[size-left:])
		if err != nil {
			Log.Error("ReadFull2 error: ", err)
			return err
		}
		left -= n
	}

	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		Log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil
	}

	authKeyID := int64(binary.LittleEndian.Uint64(buf))
	message := NewRawMessage(TRANSPORT_TCP, authKeyID, 0)
	message.Decode(buf)

	Log.Debugf("message.Payload = %v", hex.EncodeToString(message.Payload))
	s.message = message

	return nil
}

func (s *MTProto) Write(data interface{}) error {

	s.respChan <- data

	return nil
}
