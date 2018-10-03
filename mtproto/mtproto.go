package mtproto

import (
	"io"
	"net"
	// . "github.com/rockin0098/flash/base/global"
	// . "github.com/rockin0098/flash/base/logger"
)

// 协议版本
const (
	MTPROTO_ABRIDGED_VERSION     = 1 // 删节版本
	MTPROTO_INTERMEDIATE_VERSION = 2 // 中间版本
	MTPROTO_FULL_VERSION         = 3 // 完整版本
	MTPROTO_APP_VERSION          = 4 // Androd等当前客户端使用版本
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
	reader  io.Reader
	writer  io.Writer
	AuthKey []byte
}

func NewMTProto(reader io.Reader, writer io.Writer) *MTProto {
	return &MTProto{
		reader: reader,
		writer: writer,
	}
}

func (s *MTProto) RemoteAddr() net.Addr {
	return s.reader.(net.Conn).RemoteAddr()
}

func (s *MTProto) LocalAddr() net.Addr {
	return s.reader.(net.Conn).LocalAddr()
}

func (s *MTProto) Read() (interface{}, error) {

	return nil, nil
}

func (s *MTProto) Write(data interface{}) error {

	return nil
}
