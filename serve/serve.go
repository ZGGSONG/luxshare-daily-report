package serve

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"luxshare-daily-report/global"
	"luxshare-daily-report/serve/core"
	"luxshare-daily-report/util"
	"net/http"
	"os"
	"strconv"
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
		log.Errorf("(HandlerSingleFile) Get form file failed: %s", err)
		return
	}
	defer formFile.Close()
	// 创建保存文件
	destFile, err := os.Create("./upload/" + fn)
	if err != nil {
		log.Errorf("(HandlerSingleFile) Create Save failed: %s", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Errorf("(HandlerSingleFile) Write file failed: %s", err)
		return
	}
	log.Infof("Have Receive File: %v", fn) //输出上传的文件名
	str, _ := os.Getwd()
	dic := fmt.Sprintf("%v/upload/", str)
	if fn == "xcm.jpeg" {
		//路径字典
		m := make(map[string]string, 2)
		m["xcm"] = dic + "xcm.jpeg"
		m["jkm"] = dic + "jkm.jpeg"
		global.GLO_RECV_CHAN <- m
	}

	w.Write([]byte("服务端接收成功...")) //这个写入到w的是输出到客户端的
}

//
// DeclarationService
//  @Description: 每日申报服务
//  @param files
//
func DeclarationService(files map[string]string) {
	//登陆获取auth
	auth, __user__, err := core.Login(global.GLO_CONFIG.UserName, global.GLO_CONFIG.PassWord)
	if auth == "" || err != nil {
		log.Errorf("Ticket: %v, err:%v", auth, err.Error())
		util.SendMessageError(err)
		return
	}

	//上传图片
	var m = make(map[string]string, 2)
	srcXcm, _ := ioutil.ReadFile(files["xcm"])
	srcJkm, _ := ioutil.ReadFile(files["jkm"])
	resXcm := base64.StdEncoding.EncodeToString(srcXcm)
	resJkm := base64.StdEncoding.EncodeToString(srcJkm)
	m["xcm"] = resXcm
	m["jkm"] = resJkm
	imagesLinks, err := core.Upload2Azure(auth, __user__, m)
	//log.Printf("[DEBUG] get images links: %s", imagesLinks)
	if err != nil {
		log.Errorf(err.Error())
		util.SendMessageError(err)
		return
	}
	if imagesLinks == nil {
		err = errors.New("get no images links")
		log.Errorf(err.Error())
		util.SendMessageError(err)
		return
	}

	//TODO: if model is nil
	epidemicQuestLVIData, err := core.GetLVIQuestInitModel(auth, __user__)
	if err != nil {
		log.Errorf("获取个人初始化信息失败: %v", err.Error())
	}

	//申报
	for i := 0; i < 3; i++ {
		err = core.EpidemicRegistration(auth, __user__, imagesLinks, epidemicQuestLVIData)
		if err != nil && i > 1 {
			log.Errorf("重试3次失败，%v", err.Error())
			util.SendMessageError(err)
			return
		} else {
			break
		}

	}
	log.Infof("每日申报成功")

	//刷新门禁
	for i := 0; i < 3; i++ {
		err = core.RefreshDoor(auth, __user__)
		if err != nil && i > 1 {
			log.Errorf("重试3次失败，%v", err.Error())
			util.SendMessageError(err)
			return
		} else {
			break
		}
	}
	log.Infof("刷新门禁成功")

	util.SendSuccess("【成功】每日申报、刷新门禁")
}

//
// CompressImageResource
//  @Description: 压缩JPEG文件
//  @param imagePath
//
func CompressImageResource(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Errorf("(compressImageResource) Open File failed err: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Errorf("(compressImageResource) Decode error: %v", err)
	}
	buf := bytes.Buffer{}
	quality, err := strconv.Atoi(global.GLO_CONFIG.Quality)
	if err != nil {
		log.Errorf("(compressImageResource) Quality convert from config error: %v", err)
	}
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Errorf("(compressImageResource) Encode error: %v", err)
	}
	//保存到新文件中
	newFile, err := os.Create(imagePath)
	if err != nil {
		log.Errorf("(compressImageResource) Create Compress File failed err: %v", err)
	}
	defer newFile.Close()
	_, err = newFile.Write(buf.Bytes())
	if err != nil {
		log.Errorf("(compressImageResource) Write Compress File failed err: %v", err)
	} else {
		log.Infof("Compress %v success", imagePath)
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
		log.Infof("接收文件: %v", files[i].Filename) //输出上传的文件名
	}
	w.Write([]byte("successfully")) //这个写入到w的是输出到客户端的
}
