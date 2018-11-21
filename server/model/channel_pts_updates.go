package model

type ChannelPtsUpdates struct {
	Model
	ChannelId  int32  `gorm:""`
	Pts        int32  `gorm:""`
	PtsCount   int32  `gorm:""`
	UpdateType int8   `gorm:""`
	UpdateData string `gorm:"type:longtext"`
	Date2      int32  `gorm:""`
}

type ChannelPtsUpdatesDao struct{}

var channelPtsUpdatesDao = &ChannelPtsUpdatesDao{}
