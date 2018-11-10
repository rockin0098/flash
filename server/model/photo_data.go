package model

type PhotoData struct {
	Model
	PhotoID    int64  `gorm:""`
	PhotoType  int8   `gorm:""`
	DcID       int32  `gorm:""`
	VolumeID   int64  `gorm:""`
	LocalID    int32  `gorm:""`
	AccessHash int64  `gorm:""`
	Width      int32  `gorm:""`
	Height     int32  `gorm:""`
	FileSize   int32  `gorm:""`
	FilePath   string `gorm:"size:512"`
	Ext        string `gorm:"size:512"`
	FileID     int64  `gorm:"file_id"`
	State      int8   `gorm:"state"`
}
