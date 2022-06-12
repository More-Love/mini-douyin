package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string  `gorm:"type:varchar(32);not null;unique"`
	Password  string  `gorm:"type:varchar(32);not null;unique"`
	Followers []*User `gorm:"many2many:followships;"`
	Videos    []*Video
	Favorites []*Video `gorm:"many2many:favorites;"`
	Comments  []*Comment
}
