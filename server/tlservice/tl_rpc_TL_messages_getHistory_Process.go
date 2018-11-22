package tlservice

import (
	"math"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
)

const (
	kLoadTypeBackward           = 0
	kLoadTypeForward            = 1
	kLoadTypeFirstUnread        = 2
	kLoadTypeFirstAroundMessage = 3
	kLoadTypeFirstAroundDate    = 4

	kLoadTypeLimit1 = 16
)

// TODO(@benqi): only android client
// limit = count
// offset_id = max_id
func calcLoadHistoryType(isChannel bool, offsetId, offsetDate, addOffset, limit, maxId, minId int32) int {
	if limit == 1 {
		return kLoadTypeLimit1
	}

	// check isChannel??
	if isChannel && addOffset == -1 && maxId != 0 {
		return kLoadTypeBackward
	}

	if addOffset == 0 {
		return kLoadTypeBackward
	} else if addOffset == -limit+5 {
		return kLoadTypeFirstAroundDate
	} else if addOffset == -limit/2 {
		return kLoadTypeFirstAroundMessage
	} else if addOffset == -limit-1 {
		return kLoadTypeForward
	} else if addOffset == -limit+6 {
		if maxId != 0 {
			return kLoadTypeFirstUnread
		}
	}
	return kLoadTypeBackward
}

func (s *TLService) loadHistoryMessage(loadType int, selfUserId int32, peer *model.PeerUtil, offsetId, offsetDate, addOffset, limit, maxId, minId int32) []*mtproto.TL_message {
	messages := []*mtproto.TL_message{}

	switch loadType {
	case kLoadTypeLimit1:
		// 1. Load dialog last messag
		offsetId = math.MaxInt32
		messages = s.Dao.MessageDao.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, limit)
	case kLoadTypeBackward:
		messages = s.Dao.MessageDao.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, addOffset+limit)
	case kLoadTypeFirstAroundDate:
	case kLoadTypeFirstAroundMessage:
		// LOAD_HISTORY_TYPE_FORWARD and LOAD_HISTORY_TYPE_BACKWARD
		// 按升序排
		messages1 := s.Dao.MessageDao.LoadForwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, -addOffset)
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages1[i], messages1[j] = messages1[j], messages1[i]
		}
		messages = append(messages, messages1...)
		// 降序
		messages2 := s.Dao.MessageDao.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, limit+addOffset)
		messages = append(messages, messages2...)
	case kLoadTypeForward:
		messages = s.Dao.MessageDao.LoadForwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, -addOffset)
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}
	case kLoadTypeFirstUnread:
		messages = s.Dao.MessageDao.LoadBackwardHistoryMessages(selfUserId, peer.PeerType, peer.PeerId, offsetId, addOffset+limit)
	}

	return messages
}

// // TL_messages_getPinnedDialogs
// func (s *TLService) TL_messages_getHistory_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
// 	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

// 	tlobj := object
// 	tl := tlobj.(*mtproto.TL_messages_getHistory)

// 	peer := model.FromInputPeer(tl.Get_peer())
// 	if peer.PeerType == model.PEER_SELF {
// 		peer.PeerType = model.PEER_USER
// 		peer.PeerId = csess.GetUserID()
// 	}

// 	offsetid := tl.Get_offset_id()
// 	addoffset := tl.Get_add_offset()
// 	limit := tl.Get_limit()

// 	var (
// 		isChannel = peer.PeerType == base.PEER_CHANNEL
// 		users     []*mtproto.User
// 		chats     []*mtproto.Chat
// 	)

// 	loadType := calcLoadHistoryType(isChannel, offsetId, request.GetOffsetDate(), addOffset, limit, request.GetMaxId(), request.GetMinId())
// 	messages := s.loadHistoryMessage(loadType, md.UserId, peer, offsetId, request.GetOffsetDate(), addOffset, limit, request.GetMaxId(), request.GetMinId())

// 	// messagesMessages.SetMessages(messages)
// 	userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
// 	if len(userIdList) > 0 {
// 		users = s.UserModel.GetUsersBySelfAndIDList(md.UserId, userIdList)
// 		// messagesMessages.Data2.Users = users
// 	} else {
// 		users = []*mtproto.User{}
// 	}

// 	if len(chatIdList) > 0 {
// 		chats = s.ChatModel.GetChatListBySelfAndIDList(md.UserId, chatIdList)
// 	} else {
// 		chats = []*mtproto.Chat{}
// 	}

// 	// TODO(@benqi): Add channel's pts
// 	messagesSlice := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
// 		Messages: messages,
// 		Chats:    chats,
// 		Users:    users,
// 	}}
// 	messagesMessages = messagesSlice.To_Messages_Messages()

// 	return messagesMessages, nil
// }
