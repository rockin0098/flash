package model

type UserQtsUpdates struct {
	Model
	UserID     int64  `gorm:""`
	Qts        int32  `gorm:""`
	UpdateType int32  `gorm:""`
	UpdateData []byte `gorm:"type:longtext"`
	Date2      int32  `gorm:""`
}

type UserQtsUpdatesDao struct{}

var userQtsUpdatesDao = &UserQtsUpdatesDao{}
