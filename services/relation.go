package services

import (
	"errors"
)

func GetUserFollowers(id uint) ([]uint, error) {
	followerIDs, err := userRepo.GetFollowers(id)
	if err != nil {
		return nil, errors.New("获取粉丝失败")
	}
	return followerIDs, nil
}

func CountFollowers(id uint) (int64, error) {
	count, err := userRepo.CountFollowers(id)
	if err != nil {
		return 0, errors.New("获取粉丝数量失败")
	}
	return count, nil
}

func GetUserFollowing(id uint) ([]uint, error) {
	followeeIDs, err := userRepo.GetFollowees(id)
	if err != nil {
		return nil, errors.New("获取关注失败")
	}
	return followeeIDs, nil
}

func CountUserFollowing(id uint) (int64, error) {
	count, err := userRepo.CountFollowees(id)
	if err != nil {
		return 0, errors.New("获取关注数量失败")
	}
	return count, nil
}

func CheckFollow(followerID uint, followeeID uint) bool {
	hasFollowship := userRepo.CheckFollow(followerID, followeeID)
	return hasFollowship
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
