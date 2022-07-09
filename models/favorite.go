package models

type Favorite struct {
	VideoID int64 `gorm:"primaryKey"`
	UserID  int64 `gorm:"primaryKey"`
}
