package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func shopGetHandler(c *gin.Context){

	cookie,err:=c.Cookie("key_cookie")

	if err!=nil{
		cookie = "NotSet"
		c.SetCookie("key_cookie","value_cookie111",60,"/","localhost",false,true)
	}

	fmt.Printf("cookie的值是：%s\n",cookie)

	c.JSON(http.StatusOK,gin.H{
		"msg":"this shop post handler",
	})
}

func shopCommentHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"this shop comment handler",
	})
}