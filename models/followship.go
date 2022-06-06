package models

type Followship struct {
	FollowerID uint `gorm:"primaryKey"`
	UserID     uint `gorm:"primaryKey"`
}
