package model

type UserDialog struct {
	Model
	UserID              int32  `gorm:""`
	PeerType            int8   `gorm:""`
	PeerID              int32  `gorm:""`
	IsPinned            int8   `gorm:""`
	TopMessage          int32  `gorm:""`
	ReadInboxMaxID      int32  `gorm:""`
	ReadOutboxMaxID     int32  `gorm:""`
	UnreadCount         int32  `gorm:""`
	UnreadMentionsCount int32  `gorm:""`
	ShowPreviews        int8   `gorm:""`
	Silent              int8   `gorm:""`
	MuteUntil           int32  `gorm:""`
	Sound               string `gorm:"size:"256"`
	Pts                 int32  `gorm:""`
	DraftID             int32  `gorm:""`
	DraftType           int8   `gorm:""`
	DraftMessageData    string `gorm:"type:longtext"`
	Date2               int32  `gorm:""`
	Version             int64  `gorm:""`
}
