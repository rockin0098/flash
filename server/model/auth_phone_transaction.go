package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

type AuthPhoneTransaction struct {
	Model
	AuthKeyID        int64  `gorm:""`
	PhoneNumber      string `gorm:"size:64"`
	Code             string `gorm:"size:32"`
	CodeExpired      int32  `gorm:""`
	TransactionHash  string `gorm:"size:64"`
	SentCodeType     int8   `gorm:""`
	FlashCallPattern string `gorm:"size:64"`
	NextCodeType     int8   `gorm:""`
	State            int8   `gorm:""`
	ApiID            int32  `gorm:""`
	ApiHash          string `gorm:"size:64"`
	Attempts         int32  `gorm:""`
	CreatedTime      int64  `gorm:""`
	IsDeleted        int8   `gorm:""`
}

type AuthPhoneTransactionDao struct{}

var authPhoneTransactionDao = &AuthPhoneTransactionDao{}

func (s *AuthPhoneTransactionDao) GetAuthPhoneTransactionByHash(hash string) *AuthPhoneTransaction {
	db := datasource.DataSourceInstance().Master()

	res := &AuthPhoneTransaction{}
	err := db.Where("transaction_hash=?", hash).Find(res).Error
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
