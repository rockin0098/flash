package model

type AuthUser struct {
	Model
	AuthID        int64  `gorm:""`
	UserID        int64  `gorm:""`
	Hash          int64  `gorm:""`
	DeviceModel   string `gorm:"size:64"`
	Platform      string `gorm:"size:64"`
	SystemVersion string `gorm:"size:64"`
	ApiID         int32  `gorm:""`
	AppName       string `gorm:"size:64"`
	AppVersion    string `gorm:"size:64"`
	DateCreated   int32  `gorm:""`
	DateActive    int32  `gorm:""`
	Ip            string `gorm:"size:64"`
	Country       string `gorm:"size:64"`
	Region        string `gorm:"size:64"`
	DeletedAt     int64  `gorm:"deleted_at"`
}
