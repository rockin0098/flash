package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

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

func (s *ModelManager) CheckPhoneExists(phone string) bool {

	db := datasource.DataSourceInstance().Master()

	user := &User{}
	err := db.Where("phone=?", phone).Find(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return false
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return false
	}

	return true
}
