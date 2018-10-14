package service

import (
	"encoding/hex"

	"github.com/rockin0098/flash/base/datasource"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/server/model"
)

type ModelService struct{}

var modelService = &ModelService{}

func ModelServiceInstance() *ModelService {
	return modelService
}

func (s *ModelService) ModelAdd(m interface{}) error {
	db := datasource.DataSourceInstance().Master()
	return db.Create(m).Error
}

func (s *ModelService) GetAuthKeyByAuthID(authID int64) *model.AuthKey {
	db := datasource.DataSourceInstance().Master()

	auth := &model.AuthKey{}
	err := db.Where("auth_id=?", authID).Find(auth).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return auth
}

func (s *ModelService) GetAuthKeyValueByAuthID(authID int64) []byte {
	auth := s.GetAuthKeyByAuthID(authID)
	ak, err := hex.DecodeString(auth.Body)
	if err != nil {
		Log.Error(err)
		return nil
	}

	return ak
}
