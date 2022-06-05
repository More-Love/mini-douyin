package service

import (
	"mini-douyin/repository"
	"errors"
)


var userRepo = repository.UserRepo

func RegisterUser(userName string, password string) (uint, error) {

	// 先检查用户名是否已经存在
	if _, err := userRepo.GetID(userName); err == nil {
		return 0, errors.New("用户名已经存在")
	}

	var id uint
	var err error
	// 创建用户
	if id, err = repository.UserRepo.Create(userName, password); err != nil {
		return 0, errors.New("创建用户失败")
	}

	return id, nil
}


func LoginUser(userName string, password string) (uint, error) {
	
	// 检查用户名是否存在
	id, err := userRepo.GetID(userName)
	if err != nil {
		return 0, errors.New("用户名不存在")
	}
	user, err := userRepo.Get(id)
	if err != nil {
		return 0, errors.New("获取用户信息失败")
	}

	// 检查密码是否正确
	if user.Password != password {
		return 0, errors.New("密码错误")
	}

	return user.ID, nil
}

func GetUserName(id uint) (string, error) {
	user, err := userRepo.Get(id)
	if err != nil {
		return "", errors.New("获取用户信息失败")
	}
	return user.UserName, nil
}

func CountUserFollowers(id uint) (int64, error) {
	count, err := userRepo.CountFollowers(id)
	if err != nil {
		return 0, errors.New("获取粉丝数失败")
	}
	return count, nil
}