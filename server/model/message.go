package model

import "github.com/rockin0098/meow/base/datasource"

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

func (s *ModelManager) GetMessagesByIDList(userid int64, idlist []int32) []*Message {
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
