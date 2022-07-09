package models

type User struct {
	ID        int64   `gorm:"primary_key"`
	UserName  string  `gorm:"type:varchar(32);not null;unique"`
	Password  []byte  `gorm:"type:bytea;not null;unique"`
	Followers []*User `gorm:"many2many:followships;"`
	Videos    []*Video
	Favorites []*Video `gorm:"many2many:favorites;"`
	Comments  []*Comment
}
