package repository

import (
	"github.com/go-redis/redis/v8"
	"context"
	"crypto/rand"
	"encoding/base64"
	"strconv"
	"time"
)


type TokenRepository struct {
	rdb *redis.Client
	ctx context.Context
}


func (m *TokenRepository) GenerateToken() string {
	  b := make([]byte, 4)
	  rand.Read(b)
	  return base64.StdEncoding.EncodeToString(b)
}

func (m *TokenRepository) SetToken(token string, userID uint) error {
	return m.rdb.Set(m.ctx, strconv.FormatUint(uint64(userID), 10), token, time.Hour*2).Err()
}

func (m *TokenRepository) GetToken(userID uint) (string, error) {
	return m.rdb.Get(m.ctx, strconv.FormatUint(uint64(userID), 10)).Result()
}
