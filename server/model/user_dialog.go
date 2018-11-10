package model

import "github.com/rockin0098/meow/base/datasource"

type UserDialog struct {
	Model
	UserID              int64  `gorm:""`
	PeerType            int8   `gorm:""`
	PeerID              int32  `gorm:""`
	IsPinned            int8   `gorm:""`
	TopMessage          int32  `gorm:""`
	ReadInboxMaxID      int32  `gorm:""`
	ReadOutboxMaxID     int32  `gorm:""`
	UnreadCount         int32  `gorm:""`
	UnreadMentionsCount int32  `gorm:""`
	ShowPreviews        int8   `gorm:""`
	Silent              int8   `gorm:""`
	MuteUntil           int32  `gorm:""`
	Sound               string `gorm:"size:"256"`
	Pts                 int32  `gorm:""`
	DraftID             int32  `gorm:""`
	DraftType           int8   `gorm:""`
	DraftMessageData    string `gorm:"type:longtext"`
	Date2               int32  `gorm:""`
	Version             int64  `gorm:""`
}

func (s *ModelManager) GetDialogsByOffsetID(userid int64, isPinned bool, offsetid int32, limit int32) []*UserDialog {
	db := datasource.DataSourceInstance().Master()

	var dialogs []*UserDialog

	err := db.Where("user_id = ? and is_pinned = ? and top_message < ?").
		Order("top_message desc").
		Limit(limit).
		Find(&dialogs).Error
	if err != nil {
		Log.Error(err)
		return nil
	}

	return dialogs
}
