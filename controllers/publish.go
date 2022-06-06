package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"mini-douyin/services"
	"net/http"
)

type PublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data"`
	Token string                `form:"token"`
	Title string                `form:"title"`
}

type PublishActionResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func PublishAction(c *gin.Context) {
	var request PublishActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, PublishActionResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := services.GetUID(request.Token)
	fileName := uuid.New().String() + request.Data.Filename
	path := "videos/" + fileName
	c.SaveUploadedFile(request.Data, "./static/"+path)

	err := services.PublishVideo(uid, request.Title, path)

	if err != nil {
		c.JSON(http.StatusOK, PublishActionResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "成功发布视频",
	})
}

type PublishListRequest struct {
	Token  string `form:"token"`
	UserID uint   `form:"user_id"`
}

type PublishListResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

func PublishList(c *gin.Context) {
	var request PublishListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, PublishListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := services.GetUID(request.Token)

	videoIDs, err := services.GetVideosByAuthor(request.UserID)
	videos := make([]Video, len(videoIDs))
	for i, videoID := range videoIDs {
		videos[i] = *getVideoInfo(uid, videoID)
	}

	if err != nil {
		c.JSON(http.StatusOK, PublishListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取视频列表",
		VideoList:  videos,
	})
}
