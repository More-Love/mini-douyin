package services

import (
	"errors"
	"mini-douyin/repository"
)

var tokenRepo = repository.TokenRepo

func UpdateToken(userID uint) (string, error) {
	token := tokenRepo.GenerateToken()
	err := tokenRepo.SetToken(token, userID)
	if err != nil {
		return "", errors.New("更新token失败")
	}
	return token, nil
}

func GetUID(token string) uint {
	return tokenRepo.GetUID(token)
}
