package serve

import (
	"io"
	"log"
	"net/http"
	"os"
)

//
// HandlerRoot
//  @Description: 根GET访问处理
//  @param w
//  @param r
//
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Receive File Api

Just personal api for company daily reports...

https://github.com/zggsong`))
}

//
// HandlerSingleFile
//  @Description: 单文件上传处理
//  @param w
//  @param r
//
func HandlerSingleFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte(`Receive File Api

Please using POST method to upload files...

https://github.com/zggsong`))
		return
	}

	fn := r.FormValue("name")
	// 根据字段名获取表单文件
	formFile, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("[ERROR] (HandlerSingleFile) Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()
	// 创建保存文件
	destFile, err := os.Create("./upload/" + fn)
	if err != nil {
		log.Printf("[ERROR] (HandlerSingleFile) Create Save failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("[ERROR] (HandlerSingleFile) Write file failed: %s\n", err)
		return
	}
	log.Printf("[INFO] Have Receive File: %v\n", fn) //输出上传的文件名

	w.Write([]byte("Successful...")) //这个写入到w的是输出到客户端的
}

//
// HandlerMultiFiles
//  @Description: 多文件上传处理
//  @param w
//  @param r
//
func HandlerMultiFiles(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("接收文件: %v\n", files[i].Filename) //输出上传的文件名
	}
	w.Write([]byte("successfully")) //这个写入到w的是输出到客户端的
}
