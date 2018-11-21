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
		p.PeerId = peer.(*mtproto.TL_peerUser).Get_user_id()
	case mtproto.TL_CLASS_peerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.(*mtproto.TL_peerChat).Get_chat_id()
	case mtproto.TL_CLASS_peerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.(*mtproto.TL_peerChannel).Get_channel_id()
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer mtproto.TLObject) {
	switch p.PeerType {
	case PEER_SELF:
		if p.PeerId != 0 {
			peer = &mtproto.TL_peerUser{
				M_classID: mtproto.TL_CLASS_peerUser,
				M_user_id: p.PeerId,
			}
		} else if p.selfId != 0 {
			peer = &mtproto.TL_peerUser{
				M_classID: mtproto.TL_CLASS_peerUser,
				M_user_id: p.PeerId,
			}
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", p))
		}
	case PEER_USER:
		peer = &mtproto.TL_peerUser{
			M_classID: mtproto.TL_CLASS_peerUser,
			M_user_id: p.PeerId,
		}
	case PEER_CHAT:
		peer = &mtproto.TL_peerChat{
			M_classID: mtproto.TL_CLASS_peerChat,
			M_chat_id: p.PeerId,
		}
	case PEER_CHANNEL:
		peer = &mtproto.TL_peerChannel{
			M_classID:    mtproto.TL_CLASS_peerChannel,
			M_channel_id: p.PeerId,
		}
	default:
		panic(fmt.Sprintf("ToPeer(%v) error!", p))
	}
	return
}

func FromInputNotifyPeer(peer *mtproto.TL_inputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.ClassID() {
	case mtproto.TL_CLASS_inputNotifyPeer:
		p = FromInputPeer(peer.Get_peer())
	case mtproto.TL_CLASS_inputNotifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TL_CLASS_inputNotifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TL_CLASS_inputNotifyAll:
	//	p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromInputNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer mtproto.TLObject) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.TL_inputNotifyPeer{
			M_classID: mtproto.TL_CLASS_inputNotifyPeer,
			M_peer:    p.ToInputPeer(),
		}
	case PEER_USERS:
		peer = &mtproto.TL_inputNotifyUsers{
			M_classID: mtproto.TL_CLASS_inputNotifyUsers,
		}
	case PEER_CHATS:
		peer = &mtproto.TL_inputNotifyChats{
			M_classID: mtproto.TL_CLASS_inputNotifyChats,
		}
	//case PEER_ALL:
	//	peer = &mtproto.InputNotifyPeer{
	//		Constructor: mtproto.TL_CLASS_inputNotifyAll,
	//		Data2:       &mtproto.InputNotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

func FromNotifyPeer(peer mtproto.TLObject) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.ClassID() {
	case mtproto.TL_CLASS_notifyPeer:
		p = FromPeer(peer.(*mtproto.TL_notifyPeer).Get_peer())
	case mtproto.TL_CLASS_notifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TL_CLASS_notifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TL_CLASS_notifyAll:
	//	p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToNotifyPeer() (peer mtproto.TLObject) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.TL_notifyPeer{
			M_classID: mtproto.TL_CLASS_notifyPeer,
			M_peer:    p.ToPeer(),
		}
	case PEER_USERS:
		peer = &mtproto.TL_notifyUsers{
			M_classID: mtproto.TL_CLASS_notifyUsers,
		}
	case PEER_CHATS:
		peer = &mtproto.TL_notifyChats{
			M_classID: mtproto.TL_CLASS_notifyChats,
		}
	//case PEER_ALL:
	//	peer = &mtproto.NotifyPeer{
	//		Constructor: mtproto.TL_CLASS_notifyAll,
	//		Data2:       &mtproto.NotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}

func ToPeerByTypeAndID(peerType int8, peerID int32) (peer mtproto.TLObject) {
	switch peerType {
	case PEER_USER:
		peer = &mtproto.TL_peerUser{
			M_user_id: peerID,
		}
	case PEER_CHAT:
		peer = &mtproto.TL_peerChat{
			M_chat_id: peerID,
		}
	case PEER_CHANNEL:
		peer = &mtproto.TL_peerChannel{
			M_channel_id: peerID,
		}
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerID))
	}
	return
}
