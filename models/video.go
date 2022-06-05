package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	PlayURL     string
	CoverURL    string
	Title       string
	AuthorID      uint
	FavoritedBy []User `gorm:"many2many:video_favorites;"`
	Comments    []Comment
}