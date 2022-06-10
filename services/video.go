package services

import (
	"errors"
	"mini-douyin/config"
	"mini-douyin/models"
	"mini-douyin/repository"
	"time"
)

var videoRepo = repository.VideoRepo

func GetVideo(videoID uint) (*models.Video, error) {
	video, err := videoRepo.Get(videoID)
	if err != nil {
		return nil, errors.New("获取视频失败")
	}
	return video, nil
}

func GetVideoFeed(latestTime time.Time, limit int) ([]uint, error) {
	videos, err := videoRepo.GetFeed(latestTime, limit)
	if err != nil {
		return nil, errors.New("获取视频列表失败")
	}
	return videos, nil
}

func GetVideosByAuthor(authorID uint) ([]uint, error) {
	videos, err := videoRepo.GetVideosByAuthor(authorID)
	if err != nil {
		return nil, errors.New("获取视频列表失败")
	}
	return videos, nil
}

func PublishVideo(userID uint, title string, videoPath string, coverPath string) error {

	video := models.Video{
		UserID:   userID,
		Title:    title,
		PlayURL:  config.Config.StaticBaseURL + videoPath,
		CoverURL: config.Config.StaticBaseURL + coverPath,
	}

	if err := videoRepo.Create(&video); err != nil {
		return errors.New("发布视频失败")
	}
	return nil
}

func CountFavorited(videoID uint) int64 {
	count := videoRepo.CountFavorited(videoID)
	return count
}
