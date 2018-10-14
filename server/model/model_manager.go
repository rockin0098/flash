package model

import (
	"encoding/hex"

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
	if err != nil {
		Log.Error(err)
		return nil
	}

	return auth
}

func (s *ModelManager) GetAuthKeyValueByAuthID(authID int64) []byte {
	auth := s.GetAuthKeyByAuthID(authID)
	ak, err := hex.DecodeString(auth.Body)
	if err != nil {
		Log.Error(err)
		return nil
	}

	return ak
}
