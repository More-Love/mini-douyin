package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Content string
}

func GetCommentsByVideoID(videoID uint) ([]Comment, error) {
	var comments []Comment
	err := db.Where("video_id = ?", videoID).Find(&comments).Error
	return comments, err
}

func AddComment(comment *Comment) error {
	return db.Create(comment).Error
}

func DeleteComment(comment *Comment) error {
	return db.Delete(comment).Error
}
