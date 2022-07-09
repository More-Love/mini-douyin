package repository

import (
	"mini-douyin/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (m *UserRepository) GetID(name string) (int64, error) {
	var user models.User
	err := m.db.Select("ID").Where("user_name = ?", name).First(&user).Error
	return user.ID, err
}

func (m *UserRepository) Get(id int64) (*models.User, error) {
	user := &models.User{}
	err := m.db.First(user, id).Error
	return user, err
}

func (m *UserRepository) Create(name string, password [32]byte) (int64, error) {
	user := &models.User{
		UserName: name,
		Password: password[:],
	}
	err := m.db.Create(user).Error
	return user.ID, err
}

func (m *UserRepository) GetFollowers(id int64) ([]int64, error) {
	var followships []models.Followship
	err := m.db.Where("user_id = ?", id).Select("follower_id").Find(&followships).Error
	if err != nil {
		return nil, err
	}
	followerIDs := make([]int64, len(followships))
	for i, followship := range followships {
		followerIDs[i] = followship.FollowerID
	}
	return followerIDs, err
}

func (m *UserRepository) GetFollowees(id int64) ([]int64, error) {
	var followships []models.Followship
	err := m.db.Where("follower_id = ?", id).Select("user_id").Find(&followships).Error
	if err != nil {
		return nil, err
	}
	followeeIDs := make([]int64, len(followships))
	for i, followship := range followships {
		followeeIDs[i] = followship.UserID
	}
	return followeeIDs, err
}

func (m *UserRepository) CheckFollow(followerID int64, followeeID int64) (bool, error) {
	var cnt int64
	err := m.db.Model(&models.Followship{}).Where("follower_id = ? AND user_id = ?", followerID, followeeID).Count(&cnt).Error
	return cnt > 0, err
}

func (m *UserRepository) AddFollower(followerID int64, followeeID int64) error {
	return m.db.Create(&models.Followship{
		UserID:     followeeID,
		FollowerID: followerID,
	}).Error
}

func (m *UserRepository) DeleteFollower(followerID int64, followeeID int64) error {
	return m.db.Where("user_id = ? AND follower_id = ?", followeeID, followerID).Delete(&models.Followship{}).Error
}

func (m *UserRepository) CountFollowers(id int64) (int64, error) {
	var count int64
	err := m.db.Model(&models.Followship{}).Where("user_id = ?", id).Count(&count).Error
	return count, err
}

func (m *UserRepository) CountFollowees(id int64) (int64, error) {
	var count int64
	err := m.db.Model(&models.Followship{}).Where("follower_id = ?", id).Count(&count).Error
	return count, err
}

func (m *UserRepository) GetVideos(id int64) ([]models.Video, error) {
	var videos []models.Video
	err := m.db.Where("user_id = ?", id).Find(&videos).Error
	return videos, err
}

func (m *UserRepository) CheckFavorite(userID int64, videoID int64) (bool, error) {
	var cnt int64
	err := m.db.Model(&models.Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Count(&cnt).Error
	return cnt > 0, err
}

func (m *UserRepository) GetFavorites(id int64) ([]int64, error) {
	var videos []models.Video
	err := m.db.Model(idToUser(id)).Association("Favorites").Find(&videos)
	if err != nil {
		return nil, err
	}
	videoIDs := make([]int64, len(videos))
	for i, video := range videos {
		videoIDs[i] = video.ID
	}
	return videoIDs, nil
}

func (m *UserRepository) AddFavorite(userID int64, videoID int64) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}
	return m.db.Model(idToUser(userID)).Association("Favorites").Append(&video)
}

func (m *UserRepository) DeleteFavorite(userID int64, videoID int64) error {
	var video models.Video
	err := m.db.First(&video, videoID).Error
	if err != nil {
		return err
	}
	return m.db.Model(idToUser(userID)).Association("Favorites").Delete(&video)
}

func idToUser(id int64) *models.User {
	user := &models.User{}
	user.ID = id
	return user
}
