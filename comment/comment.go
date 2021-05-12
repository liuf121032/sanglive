package comment

import (
	conf "gin/config"
	"gopkg.in/ini.v1"
	"fmt"
)

type Comment struct {
	GetConfig string
}

var cfg = new(conf.AppConf)

func getConf()(string){
	//加载配置文件，配置选项在config.go中配置获取
	err:=ini.MapTo(cfg,"../config/config.ini")
	if err!=nil{
		fmt.Printf("load ini failed,err:%v\n",err)
		return "false"
	}
	return cfg.EdctConf.Address
	//fmt.Printf("测试读取配置文件：%v\n",cfg.EdctConf.Address)
}