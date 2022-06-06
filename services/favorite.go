package services

import (
	"errors"
	"mini-douyin/models"
)

func GetUserFavorites(userID uint) ([]models.Video, error) {
	favorites, err := userRepo.GetFavorites(userID)
	if err != nil {
		return nil, errors.New("获取收藏失败")
	}
	return favorites, nil
}

func CheckFavorite(userID uint, videoID uint) bool {
	favorite, err := userRepo.CheckFavorite(userID, videoID)
	if err != nil {
		return false
	}
	return favorite
}

func Favorite(userID uint, videoID uint) error {
	if err := userRepo.AddFavorite(userID, videoID); err != nil {
		return errors.New("收藏失败")
	}
	return nil
}

func Unfavorite(userID uint, videoID uint) error {
	if err := userRepo.DeleteFavorite(userID, videoID); err != nil {
		return errors.New("取消收藏失败")
	}
	return nil
}
