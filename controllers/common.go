package controllers

import (
	"mini-douyin/services"

	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

func requireLogin(c *gin.Context, token string) int64 {
	uid := services.GetUID(token)
	if uid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code":    -1,
			"status_message": "未登录或登录已过期",
		})
		c.Abort()
	}
	return uid
}

func getUserInfo(sourceID int64, targetID int64) *User {
	name, err := services.GetUserName(targetID)
	if err != nil {
		return nil
	}
	followingCount, err := services.CountUserFollowing(targetID)
	if err != nil {
		return nil
	}
	followerCount, err := services.CountUserFollowers(targetID)
	if err != nil {
		return nil
	}
	isFollow := services.CheckFollow(sourceID, targetID)

	return &User{
		FollowCount:   followingCount,
		FollowerCount: followerCount,
		ID:            targetID,
		IsFollow:      isFollow,
		Name:          name,
	}
}

func getVideoInfo(sourceID int64, targetID int64) *Video {

	videoModel, err := services.GetVideo(targetID)
	if err != nil {
		return nil
	}

	commentCount := services.CountVideoComments(targetID)
	favoriteCount := services.CountFavorited(targetID)
	isFavorite := services.CheckFavorite(sourceID, targetID)
	author := getUserInfo(sourceID, videoModel.UserID)

	return &Video{
		Author:        *author,
		CommentCount:  commentCount,
		FavoriteCount: favoriteCount,
		ID:            targetID,
		IsFavorite:    isFavorite,
		PlayURL:       videoModel.PlayURL,
		CoverURL:      videoModel.CoverURL,
		Title:         videoModel.Title,
	}
}

func getCommentInfo(sourceID int64, targetID int64) *Comment {
	commentModel, err := services.GetComment(targetID)
	if err != nil {
		return nil
	}

	user := getUserInfo(sourceID, commentModel.UserID)

	return &Comment{
		Content:    commentModel.Content,
		CreateDate: commentModel.CreatedAt.Format("01-02"),
		ID:         targetID,
		User:       *user,
	}
}
