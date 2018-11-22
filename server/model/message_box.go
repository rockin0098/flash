package model

type MessageBox struct {
	Model
	UserID           int32 `gorm:""`
	UserMessageBoxID int32 `gorm:""`
	DialogID         int64 `gorm:""`
	DialogMessageID  int32 `gorm:""`
	MessageDataID    int64 `gorm:""`
	MessageBoxType   int8  `gorm:""`
	ReplyToMsgID     int32 `gorm:""`
	MediaUnread      int8  `gorm:""`
	Date2            int32 `gorm:""`
	Deleted          int8  `gorm:""`
}

type MessageBoxDao struct{}

var messageBoxDao = &MessageBoxDao{}
