package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func r1HelloHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"this r1 hello handler",
	})
}

func r1PostHandler(c *gin.Context){

	name:=c.DefaultPostForm("name","zaosi")

	c.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("%s",name),
	})
}

func LoadR1(e *gin.Engine){
	r1Group:=e.Group("/r1")
	r1Group.GET("/hello",r1HelloHandler)
	r1Group.POST("/post",r1PostHandler)
}

