package model

import (
	"github.com/nebulaim/telegramd/biz/base"
)

type MessageBox2 struct {
	OwnerId        int32
	MessageId      int32
	MessageBoxType int8
	MediaUnread    bool
	ReplyToMsgId   int32
	*MessageData
}

func makeDialogId(fromId, peerType, peerId int32) (did int64) {
	switch peerType {
	case base.PEER_SELF:
		did = int64(fromId)<<32 | int64(fromId)
	case base.PEER_USER:
		if fromId <= peerId {
			did = int64(fromId)<<32 | int64(peerId)
		} else {
			did = int64(peerId)<<32 | int64(fromId)
		}
	case base.PEER_CHAT, base.PEER_CHANNEL:
		did = int64(-peerId)
	}
	return
}
