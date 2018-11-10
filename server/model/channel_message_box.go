package model

type ChannelMessageBox struct {
	Model
	SenderUserID        int64 `gorm:""`
	ChannelID           int32 `gorm:""`
	ChannelMessageBoxID int32 `gorm:""`
	MessageID           int64 `gorm:""`
	Date                int32 `gorm:"date"`
	Deleted             int8  `gorm:"deleted"`
}
