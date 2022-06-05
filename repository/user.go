package repository

import (
	"mini-douyin/models"
	"gorm.io/gorm"
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
	err := m.db.Model(&models.Followship{}).Where("ID = ?", id).Find(&followships).Error
	return followships, err
}

func (m *UserRepository) GetFollowees(id uint) ([]models.Followship, error) {
	var followships []models.Followship
	err := m.db.Model(&models.Followship{}).Where("FollowerID = ?", id).Find(&followships).Error
	return followships, err
}

func (m *UserRepository) CheckFollow(followerID uint, followeeID uint) (bool, error) {
	var followship models.Followship
	err := m.db.Where("FollowerID = ? AND UserID = ?", followerID, followeeID).First(&followship).Error
	if err != nil {
		return false, err
	}
	return true, nil
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
	err := m.db.Where("author_id = ?", id).Find(&videos).Error
	return videos, err
}

func (m *UserRepository) GetFavorites(id uint) ([]models.Video, error) {
	var videos []models.Video
	err := m.db.Model(&models.User{}).Where("id = ?", id).Association("Favorites").Find(&videos)
	return videos, err
}

func (m *UserRepository) AddFavorite(userID uint, videoID uint) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}	
	return m.db.Model(&models.User{}).Where("id = ?", userID).Association("Favorites").Append(&video)
}

func (m *UserRepository) DeleteFavorite(userID uint, videoID uint) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}	
	return m.db.Model(&models.User{}).Where("id = ?", userID).Association("Favorites").Delete(&video)
}

