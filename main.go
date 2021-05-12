package main

import (
	"fmt"
	"gin/app/blog"
	"gin/app/shop"
	conf "gin/config"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"os"
	"reflect"
	"time"
	comment "gin/comment"
	"gorm.io/driver/mysql"

)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

var cfg = new(conf.AppConf)


type Result struct{
	ID int
	CALL_ID string
	PHONE string
	NAME string
	TYPE string
}

func main() {


	mysqlConfig := "ekt_app_user:P@ssw0rd@tcp(rm-2ze0xh993765uw4f3.mysql.rds.aliyuncs.com:3306)/lzrs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err1 := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix: "lz_",
		},
	})
	if err1!=nil{
		panic("failed to connect database")
	}

	var result Result
	db.Raw("SELECT id,call_id,phone,name,type from lz_work_order where id>=?",40).Scan(&result)


	fmt.Printf("resut 结构为 : %+T\n,数据为：%+v\n",result,result)

	fmt.Println(result)







	t :=comment.Comment{
		"get-test",
	}
	fmt.Println(t)
	//加载配置文件，配置选项在config.go中配置获取
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}

	fmt.Printf("测试读取配置文件：%v\n", cfg.EdctConf.Address)

	//第三种app中分路由
	routers.Include(shop.LoadshopRouters, blog.LoadBlogRouters)

	r := routers.InitRouters()

	//r.Use(routers.MiddleWare())

	if err := r.Run("127.0.0.1:9100"); err != nil {
		fmt.Println("err>>>>>", err.Error())
	}

	//第二简单的分组路由，加载不通的路由
	//r:=gin.Default()
	//
	//routers.LoadR1(r)
	//routers.LoadR2(r)
	//
	//if err:=r.Run("127.0.0.1:9100");err!=nil{
	//	fmt.Println("服务启动失败！err:%v\n",err)
	//}

	//最简答的分组路由
	//r:=routers2.SetupRouter()
	//if err:=r.Run("127.0.0.1:9100");err!=nil{
	//	fmt.Println("startup service failed,err:%v\n",err)
	//}

	//r := gin.Default()
	//api接口风格
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	action = strings.Trim(action, "/")
	//	c.String(http.StatusOK, name+" is "+action)
	//})

	//基本路由
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "hello,word!",
	//		"data":"list",
	//	})
	//	c.String(http.StatusOK,"")
	//})

	//URL风格
	//r.GET("/user", func(c *gin.Context) {
	//	name:=c.DefaultQuery("name","赵四")
	//	c.String(http.StatusOK,fmt.Sprintf("hello %s",name))
	//})

	//表单参数模式 post-json
	//r.POST("/user/add", func(c *gin.Context) {
	//	types := c.DefaultPostForm("types", "post") //这个只能接受x-www-form-urlencodeed提交的数据
	//	username := c.DefaultPostForm("username", "zaosi")
	//	password := c.DefaultPostForm("password", "liuneg")
	//
	//	body := make(map[string]interface{})
	//
	//	c.BindJSON(&body)
	//
	//	c.String(http.StatusOK, fmt.Sprintf("post-json:%s\n", body))
	//
	//	c.String(http.StatusOK, fmt.Sprintf("types:%s\n", types))
	//	c.String(http.StatusOK, fmt.Sprintf("username:%s\n", username))
	//	c.String(http.StatusOK, fmt.Sprintf("password:%s\n", password))
	//})

	//上传文件
	//r.MaxMultipartMemory = 8 << 20

	//v1:=r.Group("/v1")
	//{
	//	v1.GET("/login",login)
	//	v1.GET("/submit",submit)
	//}
	//
	//v2:=r.Group("/v2")
	//{
	//	v2.POST("/login",login)
	//	v2.POST("/submit",submit)
	//}

	//r.POST("/upload",upload)
	//r.Run("127.0.0.1:9100")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func upload(c *gin.Context) {
	file, err := c.FormFile("file") //获取文件
	if err != nil {
		c.String(http.StatusInternalServerError, "上传图片出错")
	}

	//获取文件后缀
	extstring := Ext(file.Filename)
	if extstring == "" {
		//err = errors.New("上传失败，文件类型不支持，只能上传 xlsx 格式的。")
		//return
		fmt.Printf("格式是：%v\n", extstring)
	}

	extStrSlice := []string{".xlsx", ".JPG"}
	if !ContainArray(extstring, extStrSlice) {
		fmt.Printf("上传失败，文件类型不支持，只能上传 xlsx 格式的。%v\n", extstring)
		//err = errors.New("上传失败，文件类型不支持，只能上传 xlsx 格式的。")
		//return
	}

	//filepath := 'resource/public/uploads/'
	filepath := "uploads/" //从配置文件里取
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}
	//上传到的路径
	//path := 'resource/public/uploads/20060102150405test.xlsx'
	file.Filename = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	path := filepath + file.Filename                                                        //路径+文件名上传

	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		fmt.Println(fmt.Sprintf("上传失败，%v", err))
		//err = errors.New(fmt.Sprintf("上传失败，%v", err))
		//return
	}
	c.String(http.StatusOK, "图片："+file.Filename+"上传成功")
}

//Contain 判断obj是否在target中，target支持的类型array,slice,map   false:不在 true:在
func ContainArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}
