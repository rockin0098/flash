package model

import (
	"time"

	"github.com/rockin0098/meow/base/datasource"
)

type Model struct {
	ID        int64 `gorm:"AUTO_INCREMENT;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ModelManager struct{}

var modelManager = &ModelManager{}

func GetModelManager() *ModelManager {
	return modelManager
}

func (s *ModelManager) ModelAdd(m interface{}) error {
	db := datasource.DataSourceInstance().Master()
	return db.Create(m).Error
}
