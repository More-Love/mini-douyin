package repository

import (
	"gorm.io/gorm"
	"mini-douyin/models"
)

type VideoRepository struct {
	db *gorm.DB
}


func (m *VideoRepository) Get(id uint) (*models.Video, error) {
	video := &models.Video{}
	err := m.db.First(video, id).Error
	return video, err
}

func (m *VideoRepository) GetFeed(latestTime int64, limit int) ([]models.Video, error) {
	var videos []models.Video
	err := m.db.Where("created_at <= ?", latestTime).Order("created_at desc").Limit(limit).Find(&videos).Error
	return videos, err
}


func (m *VideoRepository) Create(video *models.Video) error {
	return m.db.Create(video).Error
}

func (m *VideoRepository) GetComments(videoID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := m.db.Model(&models.Video{}).Where("id = ?", videoID).Association("Comments").Find(&comments)
	return comments, err
}

func (m *VideoRepository) CountComments(videoID uint) int64 {
	count := m.db.Model(&models.Video{}).Where("id = ?", videoID).Association("Comments").Count()
	return count
}

func (m *VideoRepository) AddComment(comment *models.Comment) error {
	return m.db.Model(&models.Video{}).Where("id = ?", comment.VideoID).Association("Comments").Append(comment)
}

func (m *VideoRepository) DeleteComment(commentID uint) error {
	return m.db.Delete(&models.Comment{}, commentID).Error
}

func (m *VideoRepository) CountFavorites(videoID uint) int64 {
	count := m.db.Model(&models.Video{}).Where("id = ?", videoID).Association("Favorites").Count()
	return count
}