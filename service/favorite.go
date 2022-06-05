package service

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