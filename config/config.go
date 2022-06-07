package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config struct {
	DatabaseDSN   string
	RedisAddr     string
	TokenDuration string
	StaticBaseURL string
}

func init() {
	path := os.Getenv("MINI_DOUYIN_CONFIG_PATH")
	if path == "" {
		path = "./config.json"
	}

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open config file failed, err: %v", err))
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(fmt.Errorf("decode config file failed, err: %v", err))
	}
}
