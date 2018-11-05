package model

import "time"

type Model struct {
	ID        int64 `gorm:"AUTO_INCREMENT;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthKey struct {
	Model
	AuthID int64  `gorm:""`
	Body   string `gorm:"size:1024"`
}

type User struct {
	Model
	AccessHash     string `gorm:"size:256"`
	FirstName      string `gorm:"size:32"`
	LastName       string `gorm:"size:32"`
	UserName       string `gorm:"size:32"`
	Phone          string `gorm:"size:64;unique"`
	CountryCode    string `gorm:"size:16"`
	Bio            string `gorm:"size:256"`
	About          string `gorm:"size:512"`
	State          int    `gorm:""`
	IsBot          int    `gorm:""`
	Banned         int64  `gorm:""`
	BannedReason   string `gorm:"size:256"`
	AccountDaysTtl int    `gorm:""`
	Photos         string `gorm:"size:1024"`
	Deleted        int64  `gorm:""`
	DeletedReason  string `gorm:"size:256"`
	BannedAt       time.Time
	DeletedAt      time.Time
}
