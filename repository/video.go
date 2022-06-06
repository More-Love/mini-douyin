package repository

import (
	"gorm.io/gorm"
	"mini-douyin/models"
	"time"
)

type VideoRepository struct {
	db *gorm.DB
}

func (m *VideoRepository) Get(id uint) (*models.Video, error) {
	video := &models.Video{}
	err := m.db.First(video, id).Error
	return video, err
}

func (m *VideoRepository) GetFeed(latestTime time.Time, limit int) ([]uint, error) {
	var videos []models.Video
	err := m.db.Select("id").Where("created_at <= ?", latestTime).Order("created_at desc").Limit(limit).Find(&videos).Error
	ret := make([]uint, len(videos))
	for i, v := range videos {
		ret[i] = v.ID
	}
	return ret, err
}

func (m *VideoRepository) GetVideosByAuthor(authorID uint) ([]uint, error) {
	var videos []models.Video
	author := &models.User{}
	author.ID = authorID
	err := m.db.Model(author).Association("Videos").Find(&videos)
	ret := make([]uint, len(videos))
	for i, v := range videos {
		ret[i] = v.ID
	}
	return ret, err
}

func (m *VideoRepository) Create(video *models.Video) error {
	return m.db.Create(video).Error
}

func (m *VideoRepository) GetComments(videoID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := m.db.Model(idToVideo(videoID)).Association("Comments").Find(&comments)
	return comments, err
}

func (m *VideoRepository) CountComments(videoID uint) int64 {
	count := m.db.Model(idToVideo(videoID)).Association("Comments").Count()
	return count
}

func (m *VideoRepository) GetComment(commentID uint) (*models.Comment, error) {
	comment := models.Comment{}
	err := m.db.First(&comment, commentID).Error
	return &comment, err
}

func (m *VideoRepository) AddComment(comment *models.Comment) error {
	return m.db.Model(idToVideo(comment.VideoID)).Association("Comments").Append(comment)
}

func (m *VideoRepository) DeleteComment(commentID uint) error {
	return m.db.Delete(&models.Comment{}, commentID).Error
}

func (m *VideoRepository) CountFavorited(videoID uint) int64 {
	count := m.db.Model(idToVideo(videoID)).Association("FavoritedBy").Count()
	return count
}

func idToVideo(id uint) *models.Video {
	video := &models.Video{}
	video.ID = id
	return video
}
