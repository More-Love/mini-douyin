package controllers

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/services"
	"net/http"
)

type FavoriteActionRequest struct {
	Token      string `form:"token"`
	VideoID    uint   `form:"video_id"`
	ActionType int    `form:"action_type"`
}

type FavoriteActionResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func FavoriteAction(c *gin.Context) {
	var request FavoriteActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}
	if request.ActionType == 1 {
		err := services.Favorite(uid, request.VideoID)
		if err != nil {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, FavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功收藏",
		})
	} else if request.ActionType == 2 {
		err := services.Unfavorite(uid, request.VideoID)
		if err != nil {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, FavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功取消收藏",
		})
	}
}

type FavoriteListRequest struct {
	Token  string `form:"token"`
	UserID uint   `form:"user_id"`
}

type FavoriteListResponse struct {
	StatusCode int     `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

func FavoriteList(c *gin.Context) {
	var request FavoriteListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}
	videoIDs, err := services.GetUserFavorites(request.UserID)
	videos := make([]Video, len(videoIDs))
	for i, videoID := range videoIDs {
		videos[i] = *getVideoInfo(uid, videoID)
	}

	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取收藏列表",
		VideoList:  videos,
	})
}
