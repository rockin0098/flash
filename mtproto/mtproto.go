package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type MTProto struct {
	rwc     io.ReadWriteCloser
	AuthKey []byte
}

func NewMTProto(rwc io.ReadWriteCloser) *MTProto {
	return &MTProto{
		rwc: rwc,
	}
}

func (s *MTProto) RemoteAddr() net.Addr {
	return s.rwc.(net.Conn).RemoteAddr()
}

func (s *MTProto) LocalAddr() net.Addr {
	return s.rwc.(net.Conn).LocalAddr()
}

func (s *MTProto) Receive() (interface{}, error) {

	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(s.rwc, b)
	if err != nil {
		return nil, err
	}

	Log.Infof("n: %v, first_byte: %v", n, hex.EncodeToString(b[:1]))
	needAck := bool(b[0]>>7 == 1)

	b[0] = b[0] & 0x7f

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		Log.Info("size1: ", size)
	} else {
		Log.Info("first_byte2: ", hex.EncodeToString(b[:1]))
		b := make([]byte, 3)
		n, err = io.ReadFull(s.rwc, b)
		if err != nil {
			return nil, err
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
		Log.Info("size2: ", size)
	}

	buf := make([]byte, size)
	left := size
	for left > 0 {
		n, err = io.ReadFull(s.rwc, buf[size-left:])
		if err != nil {
			return nil, err
		}
		Log.Info("ReadFull2: ", HexBuffer(buf))
		left -= n
	}

	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		Log.Errorf("Server response error: %v", int32(binary.LittleEndian.Uint32(buf)))
		return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!! %v", s.RemoteAddr())
	}

	authKeyID := int64(binary.LittleEndian.Uint64(buf))
	if authKeyID == 0 {
		Log.Debug("authKeyID=0, unencrypted message...")
		var message = &TUnencryptedMessage{}
		message.NeedAck = needAck
		err = message.Decode(buf[8:])
		if err != nil {
			return nil, err
		}

		Log.Debugf("unencrypted message = %v", FormatStruct(message))

		return message, nil
	} else {
		Log.Infof("Recv authKeyID not 0, authKeyID = %v", authKeyID)

		// TODO(@benqi): 检查m.State状态，authKeyID不为0时codec状态必须是CODEC_AUTH_KEY_OK或CODEC_resPQ
		// if m.State != CODEC_CONNECTED && m.State != CODEC_AUTH_KEY_OK && m.State != CODEC_resPQ && m.State != CODEC_dh_gen_ok {
		// 	// 连接需要断开
		// 	return nil, fmt.Errorf("Invalid state, is CODEC_AUTH_KEY_OK or CODEC_resPQ, but is %d", m.State)
		// }
		// if m.AuthKeyId == 0 {
		// 	key := m.GetAuthKey(authKeyID)
		// 	if key == nil {
		// 		return nil, fmt.Errorf("Can't find authKey by authKeyID %d", authKeyID)
		// 	}
		// 	m.AuthKeyId = authKeyID
		// 	m.AuthKey = key
		// 	glog.Info("Found key, keyId: ", authKeyID, ", key: ", hex.EncodeToString(key))
		// } else if m.AuthKeyId != authKeyID {
		// 	return nil, fmt.Errorf("Key error, cacheKey is ", m.AuthKeyId, ", recved keyId is ", authKeyID)
		// }

		var message = &TEncryptedMessage{}
		err = message.Decode(s.AuthKey, buf[8:])
		if err != nil {
			Log.Error("Decode encrypted message error: ", err)
			return nil, err
		}

		// if m.State != CODEC_AUTH_KEY_OK {
		// 	// m.SessionId = message.SessionId
		// 	m.State = CODEC_AUTH_KEY_OK
		// }

		return message, nil
	}
}

func (s *MTProto) Send(data interface{}) error {

	return nil
}
