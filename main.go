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

	feed := douyin.Group("/feed")
	{
		feed.GET("/", controllers.Feed)
	}

	publish := douyin.Group("/publish")
	{
		publish.POST("/action/", controllers.PublishAction)
		publish.GET("/list/", controllers.PublishList)
	}

	r.Static("/static", "./static")

	r.Run(":8080")
}
