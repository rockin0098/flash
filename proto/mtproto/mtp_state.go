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
}

func NewMTProtoState() *MTProtoState {
	return &MTProtoState{
		State: MTProtoStateInitial,
	}
}
