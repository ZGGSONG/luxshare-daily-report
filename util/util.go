package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"luxshare-daily-report/global"
	"luxshare-daily-report/model"
	"net/http"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

//
// ListeningDirectory
//  @Description: 监听目录
//  @param path 目录路径
//
func ListeningDirectory(path string) {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	err = watch.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("修改权限 : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()

	//循环
	select {}
}

//
// Send
//  @Description: bark通知
//  @param msg
//
func Send(msg string) {
	favicon := "https://m.luxshare-ict.com/favicon.ico"
	if global.GLO_CONFIG.BarkUrl == "" {
		log.Printf("[ERROR] bark url is empty!!!")
		return
	}
	getUrl := fmt.Sprintf("%s/KSAT MRSB Serve/%s?icon=%s", global.GLO_CONFIG.BarkUrl, msg, favicon)
	resp, err := http.Get(getUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("[ERROR] Send to bark err: %v", err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var barkResp model.BarkResp
	if err = json.Unmarshal(body, &barkResp); err != nil {
		log.Printf("[ERROR] bark unmarshal err: %v", err)
		return
	}
	log.Printf("[INFO] Send to bark success")

}
