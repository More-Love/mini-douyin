package models

type Favorite struct {
	VideoID uint `gorm:"primaryKey"`
	UserID  uint `gorm:"primaryKey"`
}
