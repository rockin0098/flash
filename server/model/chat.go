package model

import (
	"github.com/rockin0098/meow/base/datasource"
	"github.com/rockin0098/meow/proto/mtproto"
)

type Chat struct {
	Model
	CreatorUserID    int32  `gorm:""`
	AccessHash       int64  `gorm:""`
	RandomID         int64  `gorm:""`
	ParticipantCount int32  `gorm:""`
	Title            string `gorm:"size:256"`
	PhotoID          int64  `gorm:""`
	AdminsEnabled    int8   `gorm:""`
	Deactivated      int8   `gorm:""`
	Version          int32  `gorm:""`
	Date             int32  `gorm:""`
}

type ChatDao struct{}

var chatDao = &ChatDao{}

func (s *ChatDao) GetChatByID(id int64) *Chat {

	db := datasource.DataSourceInstance().Master()

	chat := &Chat{}

	err := db.Where("id=?").Find(chat).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return chat
}

func (s *ChatDao) GetChatListBySelfAndIDList(selfUserID int32, ids []int32) []mtproto.TLObject {

	if len(ids) == 0 {
		var os []mtproto.TLObject
		return os
	}

	var chats []mtproto.TLObject

	for _, id := range ids {
		c := s.GetChatByID(int64(id))
		if c == nil {
			emp := mtproto.New_TL_chatEmpty()
			chats = append(chats, emp)
		} else {
			chats = append(chats, Chat_to_TL_chat(c, selfUserID))
		}
	}

	return chats
}
