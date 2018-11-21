package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

type UserPresence struct {
	Model
	UserID            int64  `gorm:""`
	LastSeenAt        int64  `gorm:""`
	LastSeenAuthKeyId int64  `gorm:""`
	LastSeenIp        string `gorm:"size:64"`
	Version           int64  `gorm:""`
}

type UserPresenceDao struct{}

var userPresenceDao = &UserPresenceDao{}

func (s *UserPresenceDao) GetUserPresenceByID(userID int64) *UserPresence {
	db := datasource.DataSourceInstance().Master()

	up := &UserPresence{}
	err := db.Where("user_id=?", userID).Find(up).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	return up
}
