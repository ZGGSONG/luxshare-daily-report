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
	filePath = filepath.Join(filePath, "/config/config.yml")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(filePath)
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("配置文件更新: %v", e.Name)
		global.GLO_CONFIG_CHAN <- getConfig()
	})
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("未找到配置，请先添加配置")
		os.Exit(1)
	}
	log.Infof("Loaded Config Success...")
	return getConfig()
}

func getConfig() model.Config {
	var config model.Config
	config.Quality = viper.GetString("images.quality")
	config.UserName = viper.GetString("user.username")
	config.PassWord = viper.GetString("user.password")
	config.MsgEnabled = viper.GetBool("message.enabled")
	config.MsgType = viper.GetString("message.type")
	config.BarkUrl = viper.GetString("message.bark.url")
	config.BarkKey = viper.GetString("message.bark.key")
	config.MailHost = viper.GetString("message.mail.host")
	config.MailProtocol = viper.GetString("message.mail.protocol")
	config.MailPort = viper.GetInt("message.mail.port")
	config.MailUser = viper.GetString("message.mail.username")
	config.MailPwd = viper.GetString("message.mail.password")
	config.MailFromName = viper.GetString("message.mail.from_name")
	config.MailTo = viper.GetStringSlice("message.mail.to")
	return config
}
