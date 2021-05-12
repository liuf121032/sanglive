package routers

import (
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"fmt"
)

type Option func(engine *gin.Engine)

var options []Option

func MiddleWare() gin.HandlerFunc{
	return func(c * gin.Context){
		t :=time.Now()
		fmt.Println("中间件开始执行")
		c.Set("request","中间件")

		//执行函数
		c.Next()


		status:=c.Writer.Status()
		fmt.Println("中间件执行完毕",status)
		t2:=time.Since(t)
		fmt.Println("time:",t2)
	}
}


func Include(optObjs ...Option){
	options = append(options,optObjs...)
}

func InitRouters()*gin.Engine{

	gin.DisableConsoleColor() //写日志的方法

	f,_:=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r:=gin.Default()
	//r.Use(MiddleWare())
	for _,opt:=range options{
		opt(r)
	}
	return r
}



func helloHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"hello www.topger.com",
	})
}


func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/hello",helloHandler)
	return r
}