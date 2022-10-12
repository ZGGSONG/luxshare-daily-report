package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	count = 0
	path  = "upload"
)

func main() {
	port := 7201
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handler2)

	os.Mkdir(path, 0755)
	// 监听目录下文件
	go listenFloder()

	fmt.Printf("开启服务 监听端口 %v...\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func listenFloder() {
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
			if count > 1 {
				go func() {
					fmt.Printf("%d\n", count)
				}()
			}
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
						count++
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

func handler2(w http.ResponseWriter, r *http.Request) {
	fn := r.FormValue("name")
	// 根据字段名获取表单文件
	formFile, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()
	// 创建保存文件
	destFile, err := os.Create("./" + path + "/" + fn)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}
	fmt.Printf("接收文件: %v\n", fn) //输出上传的文件名

	fmt.Fprintf(w, "上传成功!\n") //这个写入到w的是输出到客户端的
}

func handler(w http.ResponseWriter, r *http.Request) {
	//设置内存大小
	r.ParseMultipartForm(32 << 20)
	//获取上传的文件组
	files := r.MultipartForm.File["file"]
	length := len(files)
	for i := 0; i < length; i++ {
		//打开上传文件
		srcFile, err := files[i].Open()
		defer srcFile.Close()

		if err != nil {
			log.Fatal(err)
		}
		//创建上传目录
		os.Mkdir("./upload", os.ModePerm)
		//创建上传文件
		destFile, err := os.Create("./upload/" + files[i].Filename)

		defer destFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(destFile, srcFile)
		fmt.Printf("接收文件: %v\n", files[i].Filename) //输出上传的文件名
	}

	fmt.Fprintf(w, "Success!\n") //这个写入到w的是输出到客户端的
}
