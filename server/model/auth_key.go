package model

import (
	"encoding/hex"

	"github.com/jinzhu/gorm"
	"github.com/rockin0098/meow/base/datasource"
)

type AuthKey struct {
	Model
	AuthID int64  `gorm:""`
	Body   string `gorm:"size:1024"`
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
