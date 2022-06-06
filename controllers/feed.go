package controllers

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/services"
	"net/http"
	"time"
)

type FeedRequest struct {
	LatestTime int64  `form:"latest_time"`
	Token      string `form:"token"`
}

type FeedResponse struct {
	NextTime   int64   `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

func Feed(c *gin.Context) {
	var request FeedRequest
	var tm time.Time
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FeedResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if request.LatestTime == 0 {
		tm = time.Now()
	} else {
		tm = time.UnixMilli(request.LatestTime)
	}

	uid := services.GetUID(request.Token)
	videoIDs, err := services.GetVideoFeed(tm, 20)

	videos := make([]Video, len(videoIDs))
	for i, videoID := range videoIDs {
		videos[i] = *getVideoInfo(uid, videoID)
	}

	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取视频流",
		VideoList:  videos,
	})

}
