package services

import (
	"errors"
)

func GetUserFavorites(userID int64) ([]int64, error) {
	favorites, err := userRepo.GetFavorites(userID)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取收藏失败")
	}
	return favorites, nil
}

func CheckFavorite(userID int64, videoID int64) bool {
	favorite, err := userRepo.CheckFavorite(userID, videoID)
	if err != nil {
		logger.Println(err)
		return false
	}
	return favorite
}

func Favorite(userID int64, videoID int64) error {
	if err := userRepo.AddFavorite(userID, videoID); err != nil {
		logger.Println(err)
		return errors.New("收藏失败")
	}
	return nil
}

func Unfavorite(userID int64, videoID int64) error {
	if err := userRepo.DeleteFavorite(userID, videoID); err != nil {
		logger.Println(err)
		return errors.New("取消收藏失败")
	}
	return nil
}
