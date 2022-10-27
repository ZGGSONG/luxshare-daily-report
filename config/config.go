package config

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"luxshare-daily-report/global"
	"luxshare-daily-report/model"
	"os"
	"path"
	"path/filepath"
)

//
// InitialConfig
//  @Description: 初始化配置文件
//  @return Config
//
func InitialConfig() model.Config {
	workPath, _ := os.Executable()
	filePath := path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/.yml")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(filePath)
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("[INFO] 配置文件更新: %v", e.Name)
		global.GLO_CONFIG_CHAN <- getConfig()
	})
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[ERROR] 未找到配置，请先添加配置")
		os.Exit(1)
	}
	log.Printf("[INFO] Loaded Config Success...")
	return getConfig()
}

func getConfig() model.Config {
	var config model.Config
	config.Quality = viper.GetString("img.quality")
	config.UserName = viper.GetString("info.username")
	config.PassWord = viper.GetString("info.password")
	config.BarkUrl = viper.GetString("message.bark.url")
	return config
}
