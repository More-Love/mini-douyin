package repository

import (
	"gorm.io/gorm"
	"mini-douyin/models"
)

type UserRepository struct {
	db *gorm.DB
}

func (m *UserRepository) GetID(name string) (uint, error) {
	var user models.User
	err := m.db.Select("ID").Where("user_name = ?", name).First(&user).Error
	return user.ID, err
}

func (m *UserRepository) Get(id uint) (*models.User, error) {
	user := &models.User{}
	err := m.db.First(user, id).Error
	return user, err
}

func (m *UserRepository) Create(name string, password string) (uint, error) {
	user := &models.User{
		UserName: name,
		Password: password,
	}
	err := m.db.Create(user).Error
	return user.ID, err
}

func (m *UserRepository) GetFollowers(id uint) ([]models.Followship, error) {
	var followships []models.Followship
	err := m.db.Model(&models.Followship{UserID: id}).Find(&followships).Error
	return followships, err
}

func (m *UserRepository) GetFollowees(id uint) ([]models.Followship, error) {
	var followships []models.Followship
	err := m.db.Model(&models.Followship{FollowerID: id}).Find(&followships).Error
	return followships, err
}

func (m *UserRepository) CheckFollow(followerID uint, followeeID uint) bool {
	var cnt int64
	err := m.db.Where("follower_id = ? AND user_id = ?", followerID, followeeID).First(&cnt).Error
	return err != nil
}

func (m *UserRepository) AddFollower(followerID uint, followeeID uint) error {
	return m.db.Create(&models.Followship{
		UserID:     followeeID,
		FollowerID: followerID,
	}).Error
}

func (m *UserRepository) DeleteFollower(followerID uint, followeeID uint) error {
	return m.db.Where("user_id = ? AND follower_id = ?", followeeID, followerID).Delete(&models.Followship{}).Error
}

func (m *UserRepository) CountFollowers(id uint) (int64, error) {
	var count int64
	err := m.db.Model(&models.Followship{UserID: id}).Count(&count).Error
	return count, err
}

func (m *UserRepository) CountFollowees(id uint) (int64, error) {
	var count int64
	err := m.db.Model(&models.Followship{FollowerID: id}).Count(&count).Error
	return count, err
}

func (m *UserRepository) GetVideos(id uint) ([]models.Video, error) {
	var videos []models.Video
	err := m.db.Where("user_id = ?", id).Find(&videos).Error
	return videos, err
}

func (m *UserRepository) CheckFavorite(userID uint, videoID uint) (bool, error) {
	var cnt int64
	err := m.db.Where("user_id = ? AND video_id = ?", userID, videoID).Count(&cnt).Error
	return cnt > 0, err
}

func (m *UserRepository) GetFavorites(id uint) ([]uint, error) {
	var videos []models.Video
	err := m.db.Model(idToUser(id)).Association("Favorites").Find(&videos)
	if err != nil {
		return nil, err
	}
	videoIDs := make([]uint, len(videos))
	for i, video := range videos {
		videoIDs[i] = video.ID
	}
	return videoIDs, nil
}

func (m *UserRepository) AddFavorite(userID uint, videoID uint) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}
	return m.db.Model(idToUser(userID)).Association("Favorites").Append(&video)
}

func (m *UserRepository) DeleteFavorite(userID uint, videoID uint) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}
	return m.db.Model(idToUser(userID)).Association("Favorites").Delete(&video)
}

func idToUser(id uint) *models.User {
	user := &models.User{}
	user.ID = id
	return user
}
