package model

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/crypto"
	"github.com/rockin0098/meow/base/datasource"
	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
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
}

type UserDao struct{}

var userDao = &UserDao{}

func (s *UserDao) CheckPhoneExists(phone string) bool {

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

func (s *UserDao) GetUserByPhoneNumber(phone string) *User {

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

func (s *UserDao) GetUserByID(userid int64) *User {

	db := datasource.DataSourceInstance().Master()

	user := &User{}
	err := db.Where("id=?", userid).Find(user).Error
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

func (s *UserDao) CreateNewUser(user *User, authid int64) error {

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

func (s *UserDao) GetUsersByIDList(ids []int64) []*User {

	db := datasource.DataSourceInstance().Master()

	var res []*User
	err := db.Where("id in (?)", ids).Find(&res).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return res

}

func (s *UserDao) GetUsersBySelfAndIDList(selfid int64, ids []int64) []mtproto.TLObject {

	if len(ids) == 0 {
		var os []mtproto.TLObject
		return os
	}

	users := s.GetUsersByIDList(ids)
	var tlusers []mtproto.TLObject
	for _, u := range users {
		tlu := makeTLUserByUser(selfid, u)
		tlusers = append(tlusers, tlu)
	}

	return tlusers
}

func (s *UserDao) GetUserStatus(userID int64) mtproto.TLObject {
	now := time.Now().Unix()
	up := userPresenceDao.GetUserPresenceByID(userID)
	if up == nil {
		return mtproto.New_TL_userStatusEmpty()
	}

	if now <= up.LastSeenAt+5*60 {
		status := &mtproto.TL_userStatusOnline{
			M_expires: int32(up.LastSeenAt + 5*30),
		}
		return status
	} else {
		status := &mtproto.TL_userStatusOffline{
			M_was_online: int32(up.LastSeenAt),
		}
		return status
	}
}

func (s *UserDao) GetDefaultUserPhotoID(userID int64) int64 {

	user := s.GetUserByID(userID)
	if user == nil {
		return 0
	}

	if len(user.Photos) == 0 {
		return 0
	}

	// 暂时只存一个头像图片id
	defaultid, err := strconv.ParseInt(user.Photos, 10, 64)
	ASSERT(err == nil)

	return defaultid
}
