package model

import (
	"github.com/rockin0098/meow/base/datasource"
)

type UserContact struct {
	Model
	OwnerUserID      int64  `gorm:""`
	ContactUserID    int64  `gorm:""`
	ContactPhone     string `gorm:"size:64"`
	ContactFirstName string `gorm:"size:64"`
	ContactLastName  string `gorm:"size:64"`
	Mutual           int8   `gorm:""`
	IsBlocked        int8   `gorm:""`
	IsDeleted        int8   `gorm:""`
	Date2            int32  `gorm:""`
}

func (s *ModelManager) GetContactsByUserID(ownerUserID int64) []*UserContact {
	db := datasource.DataSourceInstance().Master()

	var res []*UserContact
	err := db.Where("owner_user_id=? and is_deleted = 0", ownerUserID).
		Order("contact_user_id asc").
		Find(res).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return res
}

func (s *ModelManager) IsMyContact(myid int64, contactid int64) bool {
	db := datasource.DataSourceInstance().Master()

	var res []*UserContact
	err := db.Where("owner_user_id=? and contact_user_id=? and is_deleted = 0", myid, contactid).
		Find(res).Error
	if err != nil {
		Log.Error(err)
		return false
	}

	if len(res) == 0 {
		return false
	}

	return true
}
