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

type ModelDao struct {
	AuthKeyDao              *AuthKeyDao
	AuthPhoneTransactionDao *AuthPhoneTransactionDao
	AuthSeqUpdatesDao       *AuthSeqUpdatesDao
	AuthUpdatesStateDao     *AuthUpdatesStateDao
	AuthUserDao             *AuthUserDao
	ChannelMessageBoxDao    *ChannelMessageBoxDao
	ChannelPtsUpdatesDao    *ChannelPtsUpdatesDao
	ChatDao                 *ChatDao
	MessageDao              *MessageDao
	PhotoDataDao            *PhotoDataDao
	UserContactDao          *UserContactDao
	UserDialogDao           *UserDialogDao
	UserPasswordDao         *UserPasswordDao
	UserPresenceDao         *UserPresenceDao
	UserPtsUpdatesDao       *UserPtsUpdatesDao
	UserQtsUpdatesDao       *UserQtsUpdatesDao
	UserDao                 *UserDao
	MessageDataDao          *MessageDataDao
}

var modelDao = &ModelDao{
	AuthKeyDao:              authKeyDao,
	AuthPhoneTransactionDao: authPhoneTransactionDao,
	AuthSeqUpdatesDao:       authSeqUpdatesDao,
	AuthUpdatesStateDao:     authUpdatesStateDao,
	AuthUserDao:             authUserDao,
	ChannelMessageBoxDao:    channelMessageBoxDao,
	ChannelPtsUpdatesDao:    channelPtsUpdatesDao,
	ChatDao:                 chatDao,
	MessageDao:              messageDao,
	PhotoDataDao:            photoDataDao,
	UserContactDao:          userContactDao,
	UserDialogDao:           userDialogDao,
	UserPasswordDao:         userPasswordDao,
	UserPresenceDao:         userPresenceDao,
	UserPtsUpdatesDao:       userPtsUpdatesDao,
	UserQtsUpdatesDao:       userQtsUpdatesDao,
	UserDao:                 userDao,
	MessageDataDao:          messageDataDao,
}

func GetModelDao() *ModelDao {
	return modelDao
}

func (s *ModelDao) ModelAdd(m interface{}) error {
	db := datasource.DataSourceInstance().Master()
	return db.Create(m).Error
}
