package model

type UserPassword struct {
	Model
	UserID      int64  `gorm:""`
	ServerSalt  string `gorm:"size:64"`
	Hash        string `gorm:"size:64"`
	Salt        string `gorm:"size:64"`
	Hint        string `gorm:"size:64"`
	Email       string `gorm:"size:64"`
	HasRecovery int8   `gorm:""`
	Code        string `gorm:"size:64"`
	CodeExpired int32  `gorm:""`
	Attempts    int32  `gorm:""`
	State       int8   `gorm:""`
}
