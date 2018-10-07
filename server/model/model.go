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
	Body   string `gorm:"size:512"`
}
