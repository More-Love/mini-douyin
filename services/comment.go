package services

import (
	"errors"
	"mini-douyin/models"
)

func GetVideoComments(videoID int64) ([]models.Comment, error) {
	comments, err := videoRepo.GetComments(videoID)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取评论失败")
	}
	return comments, nil
}

func CountVideoComments(videoID int64) int64 {
	count := videoRepo.CountComments(videoID)
	return count
}

func GetComment(commentID int64) (*models.Comment, error) {
	comment, err := videoRepo.GetComment(commentID)
	if err != nil {
		logger.Println(err)
		return nil, errors.New("获取评论失败")
	}
	return comment, nil
}

func AddComment(userID int64, videoID int64, content string) (int64, error) {
	comment := models.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	if err := videoRepo.AddComment(&comment); err != nil {
		logger.Println(err)
		return 0, errors.New("添加评论失败")
	}
	return comment.ID, nil
}

func DeleteComment(commentID int64) error {
	if err := videoRepo.DeleteComment(commentID); err != nil {
		logger.Println(err)
		return errors.New("删除评论失败")
	}
	return nil
}
