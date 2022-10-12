package serve

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"receive-files/global"
	"strconv"
	"time"
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
		log.Printf("[ERROR] (HandlerSingleFile) Get form file failed: %s", err)
		return
	}
	defer formFile.Close()
	// 创建保存文件
	destFile, err := os.Create("./upload/" + fn)
	if err != nil {
		log.Printf("[ERROR] (HandlerSingleFile) Create Save failed: %s", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("[ERROR] (HandlerSingleFile) Write file failed: %s", err)
		return
	}
	log.Printf("[INFO] Have Receive File: %v", fn) //输出上传的文件名
	str, _ := os.Getwd()
	dic := fmt.Sprintf("%v/upload/", str)
	if fn == "xcm.jpeg" {
		global.GLO_RECV_CHAN <- []string{dic + "jkm.jpeg", dic + "xcm.jpeg"}
	}

	w.Write([]byte("Successful...")) //这个写入到w的是输出到客户端的
}

//
// DeclarationService
//  @Description: 申报服务
//  @param files
//
func DeclarationService(files []string) {
	time.Sleep(3 * time.Second)
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Printf("[ERROR] 删除失败, %v", err)
			continue
		}
		log.Printf("[INFO] delete file %v", file)
	}
}

//
// CompressImageResource
//  @Description: 压缩JPEG文件
//  @param imagePath
//
func CompressImageResource(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Open File failed err: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Decode error: %v", err)
	}
	buf := bytes.Buffer{}
	quality, err := strconv.Atoi(global.GLO_CONFIG.Quality)
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Quality convert from config error: %v", err)
	}
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Encode error: %v", err)
	}
	//保存到新文件中
	newFile, err := os.Create(imagePath)
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Create Compress File failed err: %v", err)
	}
	defer newFile.Close()
	_, err = newFile.Write(buf.Bytes())
	if err != nil {
		log.Printf("[ERROR] (compressImageResource) Write Compress File failed err: %v", err)
	} else {
		log.Printf("[INFO] (compressImageResource) Compress %v success", imagePath)
	}
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
		log.Printf("接收文件: %v", files[i].Filename) //输出上传的文件名
	}
	w.Write([]byte("successfully")) //这个写入到w的是输出到客户端的
}
