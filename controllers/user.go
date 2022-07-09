package controllers

import (
	"mini-douyin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginOrRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginOrRegisterResponse struct {
	StatusCode    int    `json:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMessage string `json:"status_message"` // 返回状态描述
	UserID        int64  `json:"user_id"`        // 用户ID
	Token         string `json:"token"`          // 用户鉴权token
}

func Login(c *gin.Context) {
	var request LoginOrRegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, LoginOrRegisterResponse{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
	}

	userID, err := services.LoginUser(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusOK, LoginOrRegisterResponse{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	token, err := services.UpdateToken(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status_code":    0,
		"status_message": "登录成功",
		"user_id":        userID,
		"token":          token,
	})
}

func Register(c *gin.Context) {
	var request LoginOrRegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, LoginOrRegisterResponse{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
	}

	userID, err := services.RegisterUser(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusOK, LoginOrRegisterResponse{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	token, err := services.UpdateToken(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status_code":    0,
		"status_message": "注册成功",
		"user_id":        userID,
		"token":          token,
	})
}

type UserInfoRequest struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserInfoResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	User          *User  `json:"user"`
}

func UserInfo(c *gin.Context) {
	var request UserInfoRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, LoginOrRegisterResponse{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}
	user := getUserInfo(uid, request.UserID)

	if user == nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			StatusCode:    1,
			StatusMessage: "获取用户信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode:    0,
		StatusMessage: "获取用户信息成功",
		User:          user,
	})
}
