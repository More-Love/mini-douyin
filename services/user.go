package services

import (
	"crypto/sha256"
	"errors"
	"mini-douyin/repository"
)

var userRepo = repository.UserRepo

func RegisterUser(userName string, password string) (int64, error) {

	// 先检查用户名是否已经存在
	if _, err := userRepo.GetID(userName); err == nil {
		logger.Println(err)
		return 0, errors.New("用户名已经存在")
	}

	passwdHash := sha256.Sum256([]byte(password))

	var id int64
	var err error
	// 创建用户
	if id, err = repository.UserRepo.Create(userName, passwdHash); err != nil {
		logger.Println(err)
		return 0, errors.New("创建用户失败")
	}

	return id, nil
}

func LoginUser(userName string, password string) (int64, error) {

	// 检查用户名是否存在
	id, err := userRepo.GetID(userName)
	if err != nil {
		logger.Println(err)
		return 0, errors.New("用户名不存在")
	}
	user, err := userRepo.Get(id)
	if err != nil {
		logger.Println(err)
		return 0, errors.New("获取用户信息失败")
	}

	passwdHash := sha256.Sum256([]byte(password))

	// 检查密码是否正确
	if user.Password != string(passwdHash[:]) {
		logger.Println(err)
		return 0, errors.New("密码错误")
	}

	return user.ID, nil
}

func GetUserName(id int64) (string, error) {
	user, err := userRepo.Get(id)
	if err != nil {
		logger.Println(err)
		return "", errors.New("获取用户信息失败")
	}
	return user.UserName, nil
}

func CountUserFollowers(id int64) (int64, error) {
	count, err := userRepo.CountFollowers(id)
	if err != nil {
		logger.Println(err)
		return 0, errors.New("获取粉丝数失败")
	}
	return count, nil
}
