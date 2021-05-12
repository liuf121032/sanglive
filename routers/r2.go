package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func r2HelloHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"this r2 hello handler",
	})
}

func r2PostHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"this r2 post handler",
	})
}

func LoadR2(e *gin.Engine){
	r2Group:=e.Group("/r2")
	r2Group.GET("/hello",r2HelloHandler)
	r2Group.POST("/post",r2PostHandler)
}