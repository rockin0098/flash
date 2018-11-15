package model

type AuthUpdatesState struct {
	Model
	AuthKeyID int64  `gorm:""`
	UserIF    int64  `gorm:""`
	Pts       int32  `gorm:""`
	Pts2      int32  `gorm:""`
	Qts       int32  `gorm:""`
	Qts2      int32  `gorm:""`
	Seq       int32  `gorm:""`
	Seq2      int32  `gorm:""`
	Date      int32  `gorm:""`
	Date2     int32  `gorm:""`
	DeletedAt string `gorm:"size:64"`
}
