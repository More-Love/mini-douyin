package models

import "time"

type Comment struct {
	ID        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UserID    int64
	VideoID   int64
	Content   string
}
