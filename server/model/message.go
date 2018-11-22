package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

const (
	MESSAGE_TYPE_UNKNOWN         = 0
	MESSAGE_TYPE_MESSAGE_EMPTY   = 1
	MESSAGE_TYPE_MESSAGE         = 2
	MESSAGE_TYPE_MESSAGE_SERVICE = 3
)
const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
	MESSAGE_BOX_TYPE_CHANNEL  = 2
)

const (
	PTS_UNKNOWN             = 0
	PTS_MESSAGE_OUTBOX      = 1
	PTS_MESSAGE_INBOX       = 2
	PTS_READ_HISTORY_OUTBOX = 3
	PTS_READ_HISTORY_INBOX  = 4
)

type Message struct {
	Model
	UserID           int64  `gorm:""`
	UserMessageBoxID int32  `gorm:""`
	DialogMessageID  int64  `gorm:""`
	SenderUserID     int32  `gorm:""`
	MessageBoxType   int8   `gorm:""`
	PeerType         int8   `gorm:""`
	PeerID           int32  `gorm:""`
	RandomID         int64  `gorm:""`
	MessageType      int8   `gorm:""`
	MessageData      string `gorm:"type:longtext"`
	Date2            int32  `gorm:""`
	Deleted          int8   `gorm:""`
}

type MessageDao struct{}

var messageDao = &MessageDao{}

func (s *MessageDao) GetMessagesByIDList(userid int32, idlist []int32) []*Message {
	db := datasource.DataSourceInstance().Master()

	var messages []*Message

	err := db.Where("user_id = ? and deleted = 0 and user_message_box_id in (?)", userid, idlist).
		Order("user_message_box_id desc").
		Find(&messages).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return messages
}

func (s *MessageDao) GetChannelMessage(channelid int32, id int32) *Message {
	db := datasource.DataSourceInstance().Master()

	channelMsgBox := &ChannelMessageBox{}
	err := db.Where("channel_id = ? and channel_message_box_id = ?", channelid, id).
		Find(channelMsgBox).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	message := &Message{}

	err = db.Where("message_id = ? and deleted = 0", channelMsgBox.MessageID).
		Limit(1).
		Find(message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	return message
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Loadhistory
func (s *MessageDao) SelectBackwardByOffsetLimit(user_id int32, dialog_id int64, user_message_box_id int32, limit int32) []MessageBox {
	var query = `
	select 
		user_id, 
		user_message_box_id, 
		dialog_id, 
		dialog_message_id, 
		message_data_id, 
		message_box_type, 
		reply_to_msg_id, 
		media_unread, 
		date2 
	from message_box 
	where user_id = ? and 
		dialog_id = ? and 
		user_message_box_id < ? and 
		deleted = 0 
		order by user_message_box_id desc limit ?`

	var values []MessageBox
	db := datasource.DataSourceInstance().Master()
	err := db.Raw(query, user_id, dialog_id, user_message_box_id, limit).Find(&values).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return values
}

func (s *MessageDao) SelectMessageList(dialog_id int64, idList []int32) []MessageData {
	var query = "select message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, has_media_unread, `date`, edit_message, edit_date, views from message_data where dialog_id = ? and dialog_message_id in (?)"

	var values []MessageData
	db := datasource.DataSourceInstance().Master()
	err := db.Raw(query, dialog_id, idList).Find(&values).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return values
}

// func (s *MessageDao) LoadBackwardHistoryMessages(userId int32, peerType, peerId int32, offset int32, limit int32) (messages []*mtproto.TL_message) {
// 	// TODO(@benqi): chat and channel

// 	did := makeDialogId(userId, peerType, peerId)

// 	switch peerType {
// 	case PEER_USER, PEER_CHAT:
// 		boxDOList := s.SelectBackwardByOffsetLimit(userId, did, offset, limit)
// 		if len(boxDOList) == 0 {
// 			messages = []*mtproto.TL_message{}
// 			return
// 		}

// 		dialogMessageIdList := make([]int32, 0, len(boxDOList))
// 		for i := 0; i < len(boxDOList); i++ {
// 			dialogMessageIdList = append(dialogMessageIdList, boxDOList[i].DialogMessageID)
// 		}
// 		mDataDOList := s.SelectMessageList(did, dialogMessageIdList)
// 		if len(mDataDOList) == 0 {
// 			messages = []*mtproto.TL_message{}

// 			// TODO(@benqi): logo
// 			return
// 		}

// 		for i := 0; i < len(boxDOList); i++ {
// 			for j := 0; j < len(mDataDOList); j++ {
// 				if boxDOList[i].DialogMessageID == mDataDOList[j].DialogMessageID {
// 					box := m.makeMessageBoxByDO(&boxDOList[i], &mDataDOList[j])
// 					messages = append(messages, box.ToMessage(userId))
// 					break
// 				}
// 			}
// 		}

// 	case base.PEER_CHANNEL:
// 		boxDOList := m.dao.ChannelMessagesDAO.SelectBackwardByOffsetLimit(peerId, offset, limit)
// 		for i := 0; i < len(boxDOList); i++ {
// 			box := m.makeChannelMessageBoxByDO(&boxDOList[i])
// 			messages = append(messages, box.ToMessage(userId))
// 		}
// 	default:
// 		// TODO(@benqi): log
// 	}
// 	return
// }

// func (m *MessageDao) LoadForwardHistoryMessages(userId int32, peerType, peerId int32, offset int32, limit int32) (messages []*mtproto.Message) {
// 	did := makeDialogId(userId, peerType, peerId)

// 	switch peerType {
// 	case base.PEER_USER, base.PEER_CHAT:
// 		boxDOList := m.dao.MessageBoxesDAO.SelectForwardByPeerOffsetLimit(userId, did, offset, limit)
// 		if len(boxDOList) == 0 {
// 			messages = []*mtproto.Message{}
// 			return
// 		}

// 		dialogMessageIdList := make([]int32, 0, len(boxDOList))
// 		for i := 0; i < len(boxDOList); i++ {
// 			dialogMessageIdList = append(dialogMessageIdList, boxDOList[i].DialogMessageId)
// 		}
// 		mDataDOList := m.dao.MessageDatasDAO.SelectMessageList(did, dialogMessageIdList)
// 		if len(mDataDOList) == 0 {
// 			messages = []*mtproto.Message{}

// 			// TODO(@benqi): log
// 			return
// 		}

// 		for i := 0; i < len(boxDOList); i++ {
// 			for j := 0; j < len(mDataDOList); j++ {
// 				if boxDOList[i].DialogMessageId == mDataDOList[j].DialogMessageId {
// 					box := m.makeMessageBoxByDO(&boxDOList[i], &mDataDOList[j])
// 					messages = append(messages, box.ToMessage(userId))
// 					break
// 				}
// 			}
// 		}

// 	case base.PEER_CHANNEL:
// 		boxDOList := m.dao.ChannelMessagesDAO.SelectForwardByOffsetLimit(peerId, offset, limit)
// 		for i := 0; i < len(boxDOList); i++ {
// 			box := m.makeChannelMessageBoxByDO(&boxDOList[i])
// 			messages = append(messages, box.ToMessage(userId))
// 		}
// 	default:
// 		// TODO(@benqi): log
// 	}
// 	return
// }

// func (m *MessageDao) GetUserMessagesByMessageIdList(userId int32, idList []int32) (messages []*mtproto.Message) {
// 	if len(idList) == 0 {
// 		messages = []*mtproto.Message{}
// 	} else {
// 		boxDOList := m.dao.MessageBoxesDAO.SelectByMessageIdList(userId, idList)
// 		glog.Info(boxDOList)
// 		if len(boxDOList) == 0 {
// 			messages = []*mtproto.Message{}
// 			return
// 		}

// 		messageDataIdList := make([]int64, 0, len(boxDOList))
// 		for i := 0; i < len(boxDOList); i++ {
// 			messageDataIdList = append(messageDataIdList, boxDOList[i].MessageDataId)
// 		}
// 		mDataDOList := m.dao.MessageDatasDAO.SelectMessageListByDataIdList(messageDataIdList)
// 		glog.Info(mDataDOList)
// 		if len(mDataDOList) == 0 {
// 			messages = []*mtproto.Message{}
// 			// TODO(@benqi): log
// 			return
// 		}

// 		for i := 0; i < len(boxDOList); i++ {
// 			for j := 0; j < len(mDataDOList); j++ {
// 				if boxDOList[i].DialogMessageId == mDataDOList[j].DialogMessageId {
// 					box := m.makeMessageBoxByDO(&boxDOList[i], &mDataDOList[j])
// 					messages = append(messages, box.ToMessage(userId))
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// func (m *MessageDao) GetPeerMessageListByMessageDataId(userId int32, messageDataId int64) (boxList []*MessageBox2) {
// 	doList := m.dao.MessageBoxesDAO.SelectPeerMessageList(userId, messageDataId)
// 	for _, do := range doList {
// 		// TODO(@benqi): check data
// 		box := &MessageBox2{
// 			OwnerId:        do.UserId,
// 			MessageId:      do.UserMessageBoxId,
// 			MessageBoxType: do.MessageBoxType,
// 			MediaUnread:    base2.Int8ToBool(do.MediaUnread),
// 		}
// 		boxList = append(boxList, box)
// 	}
// 	return
// }

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// func (m *MessageDao) GetPeerMessageId(userId, messageId, peerId int32) int32 {
// 	do := m.dao.MessageBoxesDAO.SelectPeerMessageId(peerId, userId, messageId)
// 	if do == nil {
// 		return 0
// 	} else {
// 		return do.UserMessageBoxId
// 	}
// }

// func (m *MessageDao) DeleteByMessageIdList(userId int32, idList []int32) {
// 	if len(idList) > 0 {
// 		m.dao.MessageBoxesDAO.DeleteMessagesByMessageIdList(userId, idList)
// 	}
// }

// func (m *MessageDao) GetPeerDialogMessageIdList(userId int32, idList []int32) map[int32][]int32 {
// 	doList := m.dao.MessageBoxesDAO.SelectPeerDialogMessageIdList(userId, idList)
// 	peerMessageIdListMap := make(map[int32][]int32)

// 	for _, do := range doList {
// 		if messageIdList, ok := peerMessageIdListMap[do.UserId]; !ok {
// 			peerMessageIdListMap[do.UserId] = []int32{do.UserMessageBoxId}
// 		} else {
// 			peerMessageIdListMap[do.UserId] = append(messageIdList, do.UserMessageBoxId)
// 		}
// 	}

// 	return peerMessageIdListMap
// }

// func (m *MessageDao) GetMessageIdListByDialog(userId int32, peer *PeerUtil) []int32 {
// 	did := makeDialogId(userId, peer.PeerType, peer.PeerId)
// 	doList := m.dao.MessageBoxesDAO.SelectDialogMessageIdList(userId, did)
// 	idList := make([]int32, 0, len(doList))
// 	for i := 0; i < len(doList); i++ {
// 		idList = append(idList, doList[i].UserMessageBoxId)
// 	}
// 	return idList
// }

// func (m *MessageDao) GetClearHistoryMessages(userId int32, peer *base.PeerUtil) (lastMessageBox *MessageBox2, idList []int32) {
// 	idList = []int32{}
// 	did := makeDialogId(userId, peer.PeerType, peer.PeerId)
// 	doList := m.dao.MessageBoxesDAO.SelectDialogMessageIdList(userId, did)
// 	for i := 0; i < len(doList); i++ {
// 		if i == 0 {
// 			var err error
// 			lastMessageBox, err  = m.GetMessageBox2(int32(base.PEER_USER), userId, doList[0].UserMessageBoxId)
// 			if err != nil {
// 				return
// 			}
// 		} else {
// 			idList = append(idList, doList[i].UserMessageBoxId)
// 		}
// 	}
// 	return
// }

// func (m *MessageDao) GetChannelMessagesViews(channelId int32, idList []int32, increment bool) ([]int32) {
// 	viewsDOList := m.dao.ChannelMessagesDAO.SelectMessagesViews(channelId, idList)
// 	viewsList := make([]int32, 0, len(idList))

// 	for _, id := range idList {
// 		views := int32(1)
// 		for i := 0; i < len(viewsDOList); i++ {
// 			if viewsDOList[i].ChannelMessageId == id {
// 				if increment {
// 					views = viewsDOList[i].Views + 1
// 				} else {
// 					views = viewsDOList[i].Views
// 				}
// 				break
// 			}
// 		}
// 		viewsList = append(viewsList, views)
// 	}

// 	return viewsList
// }

// func (m *MessageDao) IncrementChannelMessagesViews(channelId int32, idList []int32) {
// 	m.dao.ChannelMessagesDAO.UpdateMessagesViews(channelId, idList)
// }
