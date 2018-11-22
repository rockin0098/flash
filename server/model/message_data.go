package model

type MessageData struct {
	Model
	MessageDataID   int64  `gorm:""`
	DialogID        int64  `gorm:""`
	DialogMessageID int32  `gorm:""`
	SenderUserID    int32  `gorm:""`
	PeerType        int8   `gorm:""`
	PeerID          int32  `gorm:""`
	RandomID        int64  `gorm:""`
	MessageType     int8   `gorm:""`
	MessageData     string `gorm:"type:longtext"`
	HasMediaUnread  int8   `gorm:""`
	Date            int32  `gorm:""`
	EditMessage     string `gorm:"type:longtext"`
	EditDate        int32  `gorm:""`
	Views           int32  `gorm:""`
	Deleted         int8   `gorm:""`
}

type MessageDataDao struct{}

var messageDataDao = &MessageDataDao{}
