package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(32);not null;unique"`
	Password  string `gorm:"type:varchar(32)"`
	Followers []User `gorm:"many2many:followship;"`
	Videos    []Video
	Favorites []Video `gorm:"many2many:video_favorites;"`
	Comments  []Comment
}
