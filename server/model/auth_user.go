package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

type AuthUser struct {
	Model
	AuthID        int64  `gorm:""`
	UserID        int64  `gorm:""`
	Hash          int64  `gorm:""`
	DeviceModel   string `gorm:"size:64"`
	Platform      string `gorm:"size:64"`
	SystemVersion string `gorm:"size:64"`
	ApiID         int32  `gorm:""`
	AppName       string `gorm:"size:64"`
	AppVersion    string `gorm:"size:64"`
	DateCreated   int32  `gorm:""`
	DateActive    int32  `gorm:""`
	Ip            string `gorm:"size:64"`
	Country       string `gorm:"size:64"`
	Region        string `gorm:"size:64"`
	DeletedAt     int64  `gorm:"deleted_at"`
}

func (s *ModelManager) GetAuthUserByAuthID(authID int64) *AuthUser {
	db := datasource.DataSourceInstance().Master()

	res := &AuthUser{}
	err := db.Where("auth_id=?", authID).Find(res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	return res
}
