package model

type AuthPhoneTransaction struct {
	Model
	AuthKeyID        int64  `gorm:""`
	PhoneNumber      string `gorm:"size:64"`
	Code             string `gorm:"size:32"`
	CodeExpired      int32  `gorm:""`
	TransactionHash  string `gorm:"size:64"`
	SentCodeType     int8   `gorm:""`
	FlashCallPattern string `gorm:"size:64"`
	NextCodeType     int8   `gorm:""`
	State            int8   `gorm:""`
	ApiID            int32  `gorm:""`
	ApiHash          string `gorm:"size:64"`
	Attempts         int32  `gorm:""`
	CreatedTime      int64  `gorm:""`
	IsDeleted        int8   `gorm:""`
}
