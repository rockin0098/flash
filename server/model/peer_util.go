package model

import (
	"fmt"

	"github.com/rockin0098/meow/proto/mtproto"
)

const (
	PEER_EMPTY   = 0
	PEER_SELF    = 1
	PEER_USER    = 2
	PEER_CHAT    = 3
	PEER_CHANNEL = 4
	PEER_USERS   = 5
	PEER_CHATS   = 6
	// PEER_ALL     = 7
	PEER_UNKNOWN = -1
)

type PeerUtil struct {
	selfId     int32
	PeerType   int32
	PeerId     int32
	AccessHash int64
}

func (p PeerUtil) String() (s string) {
	switch p.PeerType {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	//case PEER_ALL:
	//	return fmt.Sprintf("PEER_ALL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	}
	// return
}

func FromInputPeer(peer mtproto.TLObject) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.ClassID() {
	case mtproto.TL_CLASS_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.TL_CLASS_inputPeerSelf:
		p.PeerType = PEER_SELF
	case mtproto.TL_CLASS_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.(*mtproto.TL_inputPeerUser).Get_user_id()
		p.AccessHash = peer.(*mtproto.TL_inputPeerUser).Get_access_hash()
	case mtproto.TL_CLASS_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.(*mtproto.TL_inputPeerChat).Get_chat_id()
	case mtproto.TL_CLASS_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.(*mtproto.TL_inputPeerChannel).Get_channel_id()
		p.AccessHash = peer.(*mtproto.TL_inputPeerChannel).Get_access_hash()
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return
}

func FromInputPeer2(selfId int32, peer mtproto.TLObject) (p *PeerUtil) {
	p = &PeerUtil{
		selfId: selfId,
	}
	switch peer.ClassID() {
	case mtproto.TL_CLASS_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.TL_CLASS_inputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case mtproto.TL_CLASS_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.(*mtproto.TL_inputPeerUser).Get_user_id()
		p.AccessHash = peer.(*mtproto.TL_inputPeerUser).Get_access_hash()
	case mtproto.TL_CLASS_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.(*mtproto.TL_inputPeerChat).Get_chat_id()
	case mtproto.TL_CLASS_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.(*mtproto.TL_inputPeerChannel).Get_channel_id()
		p.AccessHash = peer.(*mtproto.TL_inputPeerChannel).Get_access_hash()
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return

}

func (p *PeerUtil) ToInputPeer() (peer mtproto.TLObject) {
	switch p.PeerType {
	case PEER_EMPTY:
		peer = &mtproto.TL_inputPeerEmpty{
			M_classID: mtproto.TL_CLASS_inputPeerEmpty,
		}
	case PEER_SELF:
		peer = &mtproto.TL_inputPeerSelf{
			M_classID: mtproto.TL_CLASS_inputPeerSelf,
		}
	case PEER_USER:
		peer = &mtproto.TL_inputPeerUser{
			M_classID:     mtproto.TL_CLASS_inputPeerUser,
			M_user_id:     p.PeerId,
			M_access_hash: p.AccessHash,
		}
	case PEER_CHAT:
		peer = &mtproto.TL_inputPeerChat{
			M_classID: mtproto.TL_CLASS_inputPeerChat,
			M_chat_id: p.PeerId,
		}
	case PEER_CHANNEL:
		peer = &mtproto.TL_inputPeerChannel{
			M_classID:     mtproto.TL_CLASS_inputPeerChannel,
			M_channel_id:  p.PeerId,
			M_access_hash: p.AccessHash,
		}
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", p))
	}
	return
}

func FromPeer(peer mtproto.TLObject) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.ClassID() {
	case mtproto.TL_CLASS_peerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.(*mtproto.TL).GetUserId()
	case mtproto.TLConstructor_CRC32_peerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetData2().GetChatId()
	case mtproto.TLConstructor_CRC32_peerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetData2().GetChannelId()
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer *mtproto.Peer) {
	switch p.PeerType {
	case PEER_SELF:
		if p.PeerId != 0 {
			peer = &mtproto.Peer{
				Constructor: mtproto.TLConstructor_CRC32_peerUser,
				Data2:       &mtproto.Peer_Data{UserId: p.PeerId},
			}
		} else if p.selfId != 0 {
			peer = &mtproto.Peer{
				Constructor: mtproto.TLConstructor_CRC32_peerUser,
				Data2:       &mtproto.Peer_Data{UserId: p.selfId},
			}
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", p))
		}
	case PEER_USER:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerUser,
			Data2:       &mtproto.Peer_Data{UserId: p.PeerId},
		}
	case PEER_CHAT:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChat,
			Data2:       &mtproto.Peer_Data{ChatId: p.PeerId},
		}
	case PEER_CHANNEL:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChannel,
			Data2:       &mtproto.Peer_Data{ChannelId: p.PeerId},
		}
	default:
		panic(fmt.Sprintf("ToPeer(%v) error!", p))
	}
	return
}

func FromInputNotifyPeer(peer *mtproto.InputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputNotifyPeer:
		p = FromInputPeer(peer.GetData2().GetPeer())
	case mtproto.TLConstructor_CRC32_inputNotifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TLConstructor_CRC32_inputNotifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TLConstructor_CRC32_inputNotifyAll:
	//	p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromInputNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer *mtproto.InputNotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyPeer,
			Data2: &mtproto.InputNotifyPeer_Data{
				Peer: p.ToInputPeer(),
			},
		}
	case PEER_USERS:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyUsers,
			Data2:       &mtproto.InputNotifyPeer_Data{},
		}
	case PEER_CHATS:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyChats,
			Data2:       &mtproto.InputNotifyPeer_Data{},
		}
	//case PEER_ALL:
	//	peer = &mtproto.InputNotifyPeer{
	//		Constructor: mtproto.TLConstructor_CRC32_inputNotifyAll,
	//		Data2:       &mtproto.InputNotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

func FromNotifyPeer(peer *mtproto.NotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_notifyPeer:
		p = FromPeer(peer.GetData2().GetPeer())
	case mtproto.TLConstructor_CRC32_notifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TLConstructor_CRC32_notifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TLConstructor_CRC32_notifyAll:
	//	p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToNotifyPeer() (peer *mtproto.NotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyPeer,
			Data2: &mtproto.NotifyPeer_Data{
				Peer: p.ToPeer(),
			},
		}
	case PEER_USERS:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyUsers,
			Data2:       &mtproto.NotifyPeer_Data{},
		}
	case PEER_CHATS:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyChats,
			Data2:       &mtproto.NotifyPeer_Data{},
		}
	//case PEER_ALL:
	//	peer = &mtproto.NotifyPeer{
	//		Constructor: mtproto.TLConstructor_CRC32_notifyAll,
	//		Data2:       &mtproto.NotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}

func ToPeerByTypeAndID(peerType int8, peerId int32) (peer *mtproto.Peer) {
	switch peerType {
	case PEER_USER:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerUser,
			Data2:       &mtproto.Peer_Data{UserId: peerId},
		}
	case PEER_CHAT:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChat,
			Data2:       &mtproto.Peer_Data{ChatId: peerId},
		}
	case PEER_CHANNEL:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChannel,
			Data2:       &mtproto.Peer_Data{ChannelId: peerId},
		}
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerId))
	}
	return
}
