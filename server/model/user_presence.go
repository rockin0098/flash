package model

type UserPresence struct {
	Model
	UserID            int64  `gorm:""`
	LastSeenAt        int64  `gorm:""`
	LastSeenAuthKeyId int64  `gorm:""`
	LastSeenIp        string `gorm:"size:64"`
	Version           int64  `gorm:""`
}
