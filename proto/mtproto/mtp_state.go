package mtproto

type EnumMTProtoState int

const (
	MTProtoStateInitial EnumMTProtoState = iota
)

// 保存握手信息
type MTProtoState struct {
	State       EnumMTProtoState
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	A           []byte
	P           []byte
	AuthKeyID   int64
	AuthKey     []byte
}

func NewMTProtoState() *MTProtoState {
	return &MTProtoState{
		State: MTProtoStateInitial,
	}
}
