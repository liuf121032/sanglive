package blog

import (
	"gin/routers"
	"github.com/gin-gonic/gin"
)

func LoadBlogRouters(e *gin.Engine){
	blogGroup:=e.Group("/blog")

		blogGroup.POST("/post",blogPostHandler)
		blogGroup.GET("/comment",blogCommentHandler)
		blogGroup.POST("/login",loginJson)
		blogGroup.GET("/long_async",routers.MiddleWare(),long_async)  //局部加中间件过程
		blogGroup.GET("/long_sync",long_sync)
		blogGroup.GET("/test",test)

	}