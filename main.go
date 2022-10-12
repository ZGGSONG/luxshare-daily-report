package main

import (
	"log"
	"net/http"
	"os"
	"receive-files/config"
	"receive-files/global"
	"receive-files/serve"
)

func init() {
	global.GLO_CONFIG = config.InitialConfig()
	_, err := os.Stat("upload")
	if err != nil {
		err = os.Mkdir("upload", 0755)
		if err != nil {
			log.Fatal("[ERROR] Could not create upload directory...")
			return
		}
		log.Printf("[INFO] Created Upload Directory")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serve.HandlerRoot)
	mux.HandleFunc("/upload", serve.HandlerSingleFile)

	// 监听目录下文件
	//go util.ListeningDirectory("upload")
	global.GLO_RECV_CHAN = make(chan string)

	log.Printf("[INFO] Service Listener Port At %v...\n", global.GLO_CONFIG.Port)
	err := http.ListenAndServe(":"+global.GLO_CONFIG.Port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
