package services

import (
	"errors"
	"mini-douyin/models"
)

func GetVideoComments(videoID uint) ([]models.Comment, error) {
	comments, err := videoRepo.GetComments(videoID)
	if err != nil {
		return nil, errors.New("获取评论失败")
	}
	return comments, nil
}

func CountVideoComments(videoID uint) int64 {
	count := videoRepo.CountComments(videoID)
	return count
}

func GetComment(commentID uint) (*models.Comment, error) {
	comment, err := videoRepo.GetComment(commentID)
	if err != nil {
		return nil, errors.New("获取评论失败")
	}
	return comment, nil
}

func AddComment(userID uint, videoID uint, content string) (uint, error) {
	comment := models.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	if err := videoRepo.AddComment(&comment); err != nil {
		return 0, errors.New("添加评论失败")
	}
	return comment.ID, nil
}

func DeleteComment(commentID uint) error {
	if err := videoRepo.DeleteComment(commentID); err != nil {
		return errors.New("删除评论失败")
	}
	return nil
}
