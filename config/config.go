package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

//
// InitialConfig
//  @Description: 初始化配置文件
//  @return Config
//
func InitialConfig() Config {
	workPath, _ := os.Executable()
	filePath := path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/.yml")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("未找到配置，请先添加配置！！！")
		os.Exit(1)
	}
	log.Printf("[INFO] Loaded Config Success...")
	return getConfig()
}

type Config struct {
	Port     string
	Quality  string
	UserName string
	PassWord string
	UserInfo string
}

func getConfig() Config {
	var config Config
	config.Port = viper.GetString("server.port")
	config.Quality = viper.GetString("img.quality")
	config.UserName = viper.GetString("info.username")
	config.PassWord = viper.GetString("info.password")
	config.UserInfo = viper.GetString("info.userinfo")
	return config
}
