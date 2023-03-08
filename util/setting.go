package util

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfigS struct {
	AppMode  string
	HttpPort string
}

type DBConfigS struct {
	Driver   string
	Port     string
	UserName string
	Password string
	Host     string
	Charset  string
	DBName   string
}

var (
	Appconfig *AppConfigS
	DBConfig  *DBConfigS
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path + "/config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取错误，请检查文件文件:", err)
	}

	// 开启实时监控

	if err = viper.UnmarshalKey("server", &Appconfig); err != nil {
		fmt.Println("server配置解析错误:", err)
	}
	if err = viper.UnmarshalKey("database", &DBConfig); err != nil {
		fmt.Println("数据库配置解析错误:", err)
	}
}
