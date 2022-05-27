package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(32);not null;unique"`
	Password  string `gorm:"type:varchar(32)"`
	Followers []User `gorm:"many2many:user_followers;"`
	Videos    []Video
	Favorites []Video `gorm:"many2many:video_favorites;"`
	Comments  []Comment
}

func AddUser(user *User) error {
	return db.Create(user).Error
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	err := db.Where("user_name = ?", name).First(user).Error
	return user, err
}

func GetUserById(id uint) (*User, error) {
	user := &User{}
	err := db.Where("id = ?", id).First(user).Error
	return user, err
}

func (user *User) GetFollowers() ([]User, error) {
	var users []User
	err := db.Model(user).Association("Followers").Find(&users)
	return users, err
}


func (user *User) GetFollowersCount() int64 {
	return db.Model(user).Association("Followers").Count()
}

// to-do : GetFollowing, GetFollowingCount

func (user *User) Follow(target *User) error {
	return db.Model(user).Association("Followers").Append(target)
}

func (user *User) Unfollow(target *User) error {
	return db.Model(user).Association("Followers").Delete(target)
}

func (user *User) GetFavorites() ([]Video, error) {
	var videos []Video
	err := db.Model(user).Association("Favorites").Find(&videos)
	return videos, err
}

func (user *User) Favorite(video *Video) error {
	return db.Model(user).Association("Favorites").Append(video)
}

func (user *User) Unfavorite(video *Video) error {
	return db.Model(user).Association("Favorites").Delete(video)
}
