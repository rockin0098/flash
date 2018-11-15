package model

type AuthSeqUpdates struct {
	Model
	AuthID     int64  `gorm:""`
	UserID     int64  `gorm:""`
	Seq        int32  `gorm:""`
	UpdateType int32  `gorm:""`
	UpdateData []byte `gorm:"type:longtext"`
	Date2      int32  `gorm:""`
}
