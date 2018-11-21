package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
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

func (s *MessageDao) GetMessagesByIDList(userid int64, idlist []int32) []*Message {
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
