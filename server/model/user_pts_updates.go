package model

type UserPtsUpdates struct {
	Model
	UserId     int32  `gorm:""`
	Pts        int32  `gorm:""`
	PtsCount   int32  `gorm:""`
	UpdateType int8   `gorm:""`
	UpdateData string `gorm:"type:longtext"`
	Date2      int32  `gorm:""`
}
