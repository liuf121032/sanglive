package shop

import "github.com/gin-gonic/gin"

func LoadshopRouters(e *gin.Engine){
	shopGroup:=e.Group("/shop")
	{
		shopGroup.GET("/get",shopGetHandler)
		shopGroup.POST("/comment",shopCommentHandler)
	}
}