package service

import (
	"mini-douyin/models"
	"mini-douyin/repository"
	"errors"
)

var videoRepo = repository.VideoRepo


func GetVideoFeed(latestTime int64, limit int) ([]models.Video, error) {
	videos, err := videoRepo.GetFeed(latestTime, limit)
	if err != nil {
		return nil, errors.New("获取视频列表失败")
	}
	return videos, nil
}

func PublishVideo(userID uint, title string, playURL string) error {
	video := models.Video{
		AuthorID: userID,
		Title: title,
		PlayURL: playURL,
	}
	if err := videoRepo.Create(&video); err != nil {
		return errors.New("发布视频失败")
	}
	return nil
}