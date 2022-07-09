package services

import (
	"errors"
)

func GetUserFollowers(id int64) ([]int64, error) {
	followerIDs, err := userRepo.GetFollowers(id)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取粉丝失败")
	}
	return followerIDs, nil
}

func CountFollowers(id int64) (int64, error) {
	count, err := userRepo.CountFollowers(id)
	if err != nil {
		logger.Println(err)
		return 0, errors.New("获取粉丝数量失败")
	}
	return count, nil
}

func GetUserFollowing(id int64) ([]int64, error) {
	followeeIDs, err := userRepo.GetFollowees(id)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取关注失败")
	}
	return followeeIDs, nil
}

func CountUserFollowing(id int64) (int64, error) {
	count, err := userRepo.CountFollowees(id)
	if err != nil {
		logger.Println(err)
		return 0, errors.New("获取关注数量失败")
	}
	return count, nil
}

func CheckFollow(followerID int64, followeeID int64) bool {
	hasFollowship, err := userRepo.CheckFollow(followerID, followeeID)
	if err != nil {
		logger.Println(err)
		return false
	}
	return hasFollowship
}

func Follow(followerID int64, followeeID int64) error {
	if err := userRepo.AddFollower(followerID, followeeID); err != nil {
		logger.Println(err)
		return errors.New("关注失败")
	}
	return nil
}

func Unfollow(followerID int64, followeeID int64) error {
	if err := userRepo.DeleteFollower(followerID, followeeID); err != nil {
		logger.Println(err)
		return errors.New("取消关注失败")
	}
	return nil
}
