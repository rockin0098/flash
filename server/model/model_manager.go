package model

import (
	"encoding/hex"

	"github.com/jinzhu/gorm"

	"github.com/rockin0098/flash/base/datasource"
	. "github.com/rockin0098/flash/base/logger"
)

type ModelManager struct{}

var modelManager = &ModelManager{}

func GetModelManager() *ModelManager {
	return modelManager
}

func (s *ModelManager) ModelAdd(m interface{}) error {
	db := datasource.DataSourceInstance().Master()
	return db.Create(m).Error
}

func (s *ModelManager) GetAuthKeyByAuthID(authID int64) *AuthKey {
	db := datasource.DataSourceInstance().Master()

	auth := &AuthKey{}
	err := db.Where("auth_id=?", authID).Find(auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		Log.Warn(err)
		return nil
	}

	return auth
}

func (s *ModelManager) GetAuthKeyValueByAuthID(authID int64) []byte {
	auth := s.GetAuthKeyByAuthID(authID)
	if auth == nil {
		return nil
	}
	ak, err := hex.DecodeString(auth.Body)
	if err != nil {
		Log.Error(err)
		return nil
	}

	return ak
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
