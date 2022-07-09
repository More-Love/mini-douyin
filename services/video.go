package services

import (
	"errors"
	"mini-douyin/config"
	"mini-douyin/models"
	"mini-douyin/repository"
	"time"
)

var videoRepo = repository.VideoRepo

func GetVideo(videoID int64) (*models.Video, error) {
	video, err := videoRepo.Get(videoID)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取视频失败")
	}
	return video, nil
}

func GetVideoFeed(latestTime time.Time, limit int) ([]int64, error) {
	videos, err := videoRepo.GetFeed(latestTime, limit)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取视频列表失败")
	}
	return videos, nil
}

func GetVideosByAuthor(authorID int64) ([]int64, error) {
	videos, err := videoRepo.GetVideosByAuthor(authorID)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取视频列表失败")
	}
	return videos, nil
}

func PublishVideo(userID int64, title string, videoPath string, coverPath string) error {

	video := models.Video{
		UserID:   userID,
		Title:    title,
		PlayURL:  config.Config.StaticBaseURL + videoPath,
		CoverURL: config.Config.StaticBaseURL + coverPath,
	}

	if err := videoRepo.Create(&video); err != nil {
		logger.Println(err)
		return errors.New("发布视频失败")
	}
	return nil
}

func CountFavorited(videoID int64) int64 {
	count := videoRepo.CountFavorited(videoID)
	return count
}
