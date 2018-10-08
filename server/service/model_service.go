package service

import (
	"github.com/rockin0098/flash/base/datasource"
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

func (s *ModelService) AuthKeyAdd() {

}
