package models

type Followship struct {
	FollowerID int64 `gorm:"primaryKey"`
	UserID     int64 `gorm:"primaryKey"`
}
