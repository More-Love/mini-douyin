package models

import "time"

type Video struct {
	ID          int64 `gorm:"primary_key"`
	CreatedAt   time.Time `gorm:"not null"`
	PlayURL     string
	CoverURL    string
	Title       string
	UserID      int64
	FavoritedBy []*User `gorm:"many2many:favorites;"`
	Comments    []*Comment
}
