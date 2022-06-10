package services

import (
	"errors"
	"github.com/google/uuid"
	"mini-douyin/repository"
)

var tokenRepo = repository.TokenRepo

func UpdateToken(userID uint) (string, error) {
	token := uuid.NewString()
	err := tokenRepo.SetToken(token, userID)
	if err != nil {
		logger.Println(err)
		return "", errors.New("更新token失败")
	}
	return token, nil
}

func GetUID(token string) uint {
	return tokenRepo.GetUID(token)
}
