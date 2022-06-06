package main

import (
	"mini-douyin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	{
		user.POST("/login/", controllers.Login)
		user.POST("/register/", controllers.Register)
		user.GET("/", controllers.UserInfo)
	}

	douyin.GET("/feed", controllers.Feed)

	publish := douyin.Group("/publish")
	{
		publish.POST("/action/", controllers.PublishAction)
		publish.GET("/list/", controllers.PublishList)
	}

	favorite := douyin.Group("/favorite")
	{
		favorite.POST("/action/", controllers.FavoriteAction)
		favorite.GET("/list/", controllers.FavoriteList)
	}

	comment := douyin.Group("/comment")
	{
		comment.POST("/action/", controllers.CommentAction)
		comment.GET("/list/", controllers.CommentList)
	}

	relation := douyin.Group("/relation")
	{
		relation.POST("/action/", controllers.FollowAction)
		follow := relation.Group("/follow")
		{
			follow.GET("/list/", controllers.FollowList)
		}

		follower := relation.Group("/follower")
		{
			follower.GET("/list/", controllers.FollowerList)
		}
	}

	r.Static("/static", "./static")

	r.Run(":8080")
}
