package main

import (
	"luxshare-daily-report/config"
	"luxshare-daily-report/global"
	"luxshare-daily-report/model"
	"luxshare-daily-report/serve"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// 监听目录下文件
	global.GLO_RECV_CHAN = make(chan map[string]string)
	global.GLO_CONFIG_CHAN = make(chan model.Config)

	// 初始化日志
	logger := &lumberjack.Logger{
		//Filename:   "./Log/Receive_File_Log" + time.Now().Format("20060102_150405") + ".txt",
		Filename:   "./log/Daily_Report_Log.txt",
		MaxSize:    10,   // 日志文件大小，单位是 MB
		MaxBackups: 3,    // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间，单位 天
		Compress:   true, // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
		LocalTime:  true, // 是否使用本地时间，默认是使用UTC时间
	}
	log.SetOutput(logger) // logrus 设置日志的输出方式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Infof("======")
	// 初始化配置文件
	global.GLO_CONFIG = config.InitialConfig()

	// 检查输出目录
	_, err := os.Stat("upload")
	if err != nil {
		err = os.Mkdir("upload", 0755)
		if err != nil {
			log.Fatal("Could not create upload directory...")
			return
		}
		log.Infof("Created Upload Directory")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serve.HandlerRoot)
	mux.HandleFunc("/upload", serve.HandlerSingleFile)

	// 监听每日申报
	go chanHandler()

	log.Infof("Service Listener Port At 7201...")
	err := http.ListenAndServe(":7201", mux)
	if err != nil {
		log.Fatalf("[FATAL] Start Server Error %s", err)
	}
}

// chanHandler
//
//	@Description: 信号处理
func chanHandler() {
	for {
		select {
		case newConf := <-global.GLO_CONFIG_CHAN:
			global.GLO_CONFIG = newConf
		case result := <-global.GLO_RECV_CHAN:
			for _, v := range result {
				serve.CompressImageResource(v)
			}
			serve.DeclarationService(result)
		}
	}
}
