package service

import (
	"errors"
	"mini-douyin/models"
)


func GetComments(videoID uint) ([]models.Comment, error) {
	comments, err := videoRepo.GetComments(videoID)
	if err != nil {
		return nil, errors.New("获取评论失败")
	}
	return comments, nil
}


func AddComment(userID uint, videoID uint, content string) error {
	comment := models.Comment{
		UserID: userID,
		VideoID: videoID,
		Content: content,
	}
	if err := videoRepo.AddComment(&comment); err != nil {
		return errors.New("添加评论失败")
	}
	return nil
}


func DeleteComment(commentID uint) error {
	if err := videoRepo.DeleteComment(commentID); err != nil {
		return errors.New("删除评论失败")
	}
	return nil
}