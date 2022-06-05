package service

import (
	"mini-douyin/repository"
)

var tokenRepo = repository.TokenRepo

func TestToken(userID  uint, token string) bool {
	expected, err := tokenRepo.GetToken(userID)
	if err != nil {
		return false
	}
	return token == expected
}


func ResetToken(userID uint) error {
	token := tokenRepo.GenerateToken()
	return tokenRepo.SetToken(token, userID)
}