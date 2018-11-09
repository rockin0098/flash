package model

import (
	"encoding/hex"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/crypto"
	"github.com/rockin0098/meow/base/datasource"
)

type User struct {
	Model
	AccessHash     int64  `gorm:""`
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

func (s *ModelManager) GetUserByPhoneNumber(phone string) *User {

	db := datasource.DataSourceInstance().Master()

	user := &User{}
	err := db.Where("phone=?", phone).Find(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	return user
}

func (s *ModelManager) CreateNewUser(user *User, authid int64) error {

	db := datasource.DataSourceInstance().Master()

	tx := db.Begin()

	err := tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		Log.Error(err)
		return err
	}

	au := &AuthUser{
		AuthID: authid,
		UserID: user.ID,
	}

	err = tx.Create(au).Error
	if err != nil {
		tx.Rollback()
		Log.Error(err)
		return err
	}

	password := &UserPassword{
		UserID:     user.ID,
		ServerSalt: hex.EncodeToString(crypto.GenerateNonce(8)),
	}

	err = tx.Create(password).Error
	if err != nil {
		tx.Rollback()
		Log.Error(err)
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Log.Error(err)
		return err
	}

	return nil
}

func (s *ModelManager) GetUsersByIDList(ids []int64) []*User {

	db := datasource.DataSourceInstance().Master()

	var res []*User
	err := db.Where("id in (?)", ids).Find(res).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return res

}
