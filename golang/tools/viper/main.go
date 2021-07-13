package viper

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

//配置文件名
var configFileName string

type configStruct struct {
}

var configDetail configStruct

//配置管理库viper
func loadConfigFromYaml() {
	vip := viper.New()
	vip.SetConfigFile(configFileName)
	//读取配置文件
	if err := vip.ReadInConfig(); err != nil {
		log.Println("读取配置失败")
	}
	if err := vip.Unmarshal(&configDetail); err != nil {
		log.Println("转义配置失败")
	}
	fmt.Println("配置信息:", configDetail)
}
