package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 7201
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handler2)

	// 监听目录下文件
	go listenFloder("./upload")

	fmt.Printf("开启服务 监听端口 %v...\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func listenFloder(path string) {

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
	destFile, err := os.Create("./upload/" + fn)
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
