package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	PlayURL     string
	CoverURL    string
	Title       string
	UserID      uint
	FavoritedBy []User `gorm:"many2many:video_favorites;"`
	Comments    []Comment
}

func GetVideoFeed(latestTime int64, limit int) ([]Video, error) {
	var videos []Video
	err := db.Where("created_at < ?", latestTime).Order("created_at desc").Limit(limit).Find(&videos).Error
	return videos, err
}


func AddVideo(video *Video) error {
	return db.Create(video).Error
}

func (video *Video) GetFavoritesCount() int64 {
	return db.Model(video).Association("Favorites").Count()
}

func (video *Video) GetCommentsCount() int64 {
	return db.Model(video).Association("Comments").Count()
}