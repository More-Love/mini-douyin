package controllers

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/services"
	"net/http"
)

type FollowActionRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int    `form:"action_type"`
}

type FollowActionResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func FollowAction(c *gin.Context) {
	var request FollowActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FollowActionResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := services.GetUID(request.Token)
	if request.ActionType == 1 {
		err := services.Follow(uid, request.ToUserId)
		if err != nil {
			c.JSON(http.StatusOK, FollowActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, FollowActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功关注",
		})
	} else if request.ActionType == 2 {
		err := services.Unfollow(uid, request.ToUserId)
		if err != nil {
			c.JSON(http.StatusOK, FollowActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, FollowActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功取消关注",
		})
	}
}

type FollowListRequest struct {
	Token  string `form:"token"`
	UserID int64  `form:"user_id"`
}

type FollowListResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

func FollowList(c *gin.Context) {
	var request FollowListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FollowListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}
	followList, err := services.GetUserFollowing(request.UserID)

	userList := make([]User, len(followList))
	for i, id := range followList {
		userList[i] = *getUserInfo(uid, id)
	}

	if err != nil {
		c.JSON(http.StatusOK, FollowListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取关注列表",
		UserList:   userList,
	})
}

type FollowerListRequest struct {
	Token  string `form:"token"`
	UserID int64  `form:"user_id"`
}

type FollowerListResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

func FollowerList(c *gin.Context) {
	var request FollowerListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, FollowerListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}
	userIDs, err := services.GetUserFollowers(request.UserID)
	if err != nil {
		c.JSON(http.StatusOK, FollowerListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	userList := make([]User, len(userIDs))
	for i, userID := range userIDs {
		userList[i] = *getUserInfo(uid, userID)
	}

	c.JSON(http.StatusOK, FollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取关注列表",
		UserList:   userList,
	})
}
