package controllers

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/services"
	"net/http"
)

type CommentActionRequest struct {
	Token       string `form:"token"`
	VideoID     int64  `form:"video_id"`
	ActionType  int    `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentID   int64  `form:"comment_id"`
}

type CommentActionResponse struct {
	StatusCode int      `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	Comment    *Comment `json:"comment"`
}

func CommentAction(c *gin.Context) {
	var request CommentActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
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
		id, err := services.AddComment(uid, request.VideoID, request.CommentText)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		comment := getCommentInfo(uid, id)
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功评论",
			Comment:    comment,
		})
	} else if request.ActionType == 2 {
		comment, err := services.GetComment(request.CommentID)
		if err != nil || comment.UserID != uid {
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "评论不存在",
			})
			return
		}

		err = services.DeleteComment(request.CommentID)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "成功删除评论",
		})
	}
}

type CommentListRequest struct {
	Token   string `form:"token"`
	VideoID int64  `form:"video_id"`
}

type CommentListResponse struct {
	StatusCode int        `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	Comments   []*Comment `json:"comment_list"`
}

func CommentList(c *gin.Context) {
	var request CommentListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	uid := requireLogin(c, request.Token)
	if uid == 0 {
		return
	}

	comments, err := services.GetVideoComments(request.VideoID)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	commentList := make([]*Comment, len(comments))
	for i, comment := range comments {
		commentList[i] = getCommentInfo(uid, comment.ID)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取评论列表",
		Comments:   commentList,
	})
}
