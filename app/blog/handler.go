package blog

import (
	"fmt"
	conf "gin/config"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"time"
)

var cfg = new(conf.AppConf)

func blogPostHandler(c *gin.Context){
	name:=c.PostForm("name")
	c.JSON(http.StatusOK,gin.H{
		"msg":fmt.Sprintf("%s",name),
	})
}



func blogCommentHandler(c *gin.Context){


	err:=ini.MapTo(cfg,"./config/config.ini")
	if err!=nil{
		fmt.Printf("load ini failed,err:%v\n",err)
		return
	}

	fmt.Println("config配置文件，Address：",cfg.EdctConf.Address)

	c.JSON(http.StatusOK,gin.H{
		"msg":"this blog comment handler",
		"Address":cfg.EdctConf.Address,
	})
}

type Login struct {
	User string `form:"user1" json:"user2" uri:"user3" xml:"user4"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}


func loginJson(c *gin.Context){
	var params Login

  //绑定json数据 ，获取post json的数据
	if err:=c.ShouldBindJSON(&params);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	fmt.Printf("params = %+v\n",params)
	fmt.Println(params.User)

	//return

	//if params.User !="root" || params.Password!="admin"{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"status":"304",
	//	})
	//	return
	//}
if params.User !="root" || params.Password!="admin"{
		c.JSON(http.StatusBadRequest,gin.H{
			"params":params,
			"status":"304",
		})
		return
	}


	var msg struct{
		Message string
	}

	msg.Message ="这只是一个测试"


	c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")

	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"data":params,
		"user":params.User,
		"pwd":params.Password,
		"msg":fmt.Sprintf("%s,%s",params.User,params.Password),
		"message":msg,
	})

}

func long_async(c *gin.Context){
	copyContext:=c.Copy()
	routers.MiddleWare()
	go func(){
		time.Sleep(3 * time.Second)
		log.Println("异步执行:"+ copyContext.Request.URL.Path)
	}()

	req,_:=c.Get("request")

	log.Println("request的内容是：",req)


}

func long_sync(c *gin.Context){
	time.Sleep(3*time.Second)
	log.Println("同步执行"+ c.Request.URL.Path)
}


func test(c *gin.Context){
	// a:=[3]int{1,2,3}
	// var b[20]int
	// b[19]=1
	// b[0]=101
	//
	//fmt.Println(a)
	//fmt.Println(b[0])
	//
	// c.JSON(http.StatusMovedPermanently,gin.H{
	// 	"data":b[0],
	// 	"code":http.StatusMovedPermanently,
	// 	"msg":"success",
	// })


}


func test03(n1 *int) {
	*n1 = *n1 + 10
	*n1++
	fmt.Println("n1=", *n1)
}



