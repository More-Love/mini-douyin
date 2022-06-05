package service

import (
	"errors"
	"mini-douyin/models"
)


func GetUserFollowers(id uint) ([]models.Followship, error) {
	followers, err := userRepo.GetFollowers(id)
	if err != nil {
		return nil, errors.New("获取粉丝失败")
	}
	return followers, nil
}

func GetUserFollowing(id uint) ([]models.Followship, error) {
	followees, err := userRepo.GetFollowees(id)
	if err != nil {
		return nil, errors.New("获取关注失败")
	}
	return followees, nil
}

func CheckFollow(followerID uint, followeeID uint) (bool, error) {
	hasFollowship, err := userRepo.CheckFollow(followerID, followeeID)
	if err != nil {
		return false, errors.New("检查关注失败")
	}
	return hasFollowship, nil
}

func Follow(followerID uint, followeeID uint) error {
	if err := userRepo.AddFollower(followerID, followeeID); err != nil {
		return errors.New("关注失败")
	}
	return nil
}

func Unfollow(followerID uint, followeeID uint) error {
	if err := userRepo.DeleteFollower(followerID, followeeID); err != nil {
		return errors.New("取消关注失败")
	}
	return nil
}