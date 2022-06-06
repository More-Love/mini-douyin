package controllers

import (
	"mini-douyin/services"
)

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            uint   `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            uint   `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

func getUserInfo(sourceID uint, targetID uint) *User {
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

func getVideoInfo(sourceID uint, targetID uint) *Video {

	videoModel, err := services.GetVideo(targetID)
	if err != nil {
		return nil
	}

	commentCount := services.CountComments(targetID)
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
