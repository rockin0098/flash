package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

type UserPtsUpdates struct {
	Model
	UserId     int32  `gorm:""`
	Pts        int32  `gorm:""`
	PtsCount   int32  `gorm:""`
	UpdateType int8   `gorm:""`
	UpdateData string `gorm:"type:longtext"`
	Date2      int32  `gorm:""`
}

type UserPtsUpdatesDao struct{}

var userPtsUpdatesDao = &UserPtsUpdatesDao{}

func (s *UserPtsUpdatesDao) GetUserPtsUpdatesByID(userID int64) *UserPtsUpdates {
	db := datasource.DataSourceInstance().Master()

	up := &UserPtsUpdates{}
	err := db.Where("user_id=?", userID).Find(up).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {

		upt := &UserPtsUpdates{
			Pts: 1,
		}

		err := db.Save(upt).Error
		if err != nil {
			Log.Error(err)
			return nil
		}

		return upt
	}

	return up
}
