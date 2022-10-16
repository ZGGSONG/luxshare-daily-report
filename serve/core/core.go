package core

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"receive-files/model"
	"strings"
)

//
// Login
//  @Description: 登陆网站
//  @param userName
//  @param passwd
//  @return string
//
func Login(userName, passwd string) (string, error) {
	loginUrl := "https://m.luxshare-ict.com/api/Account/Login"
	contentType := "application/x-www-form-urlencoded"

	userNameBase64 := base64.StdEncoding.EncodeToString([]byte(userName))
	passwdBase64 := base64.StdEncoding.EncodeToString([]byte(passwd))
	//fmt.Printf("username: %s, passwd: %v\n", userNameBase64, passwdBase64)

	postData := url.Values{}
	postData.Add("code", userNameBase64)
	postData.Add("password", passwdBase64)
	postData.Add("openid", "")
	postData.Add("dataSource", "M")
	resp, err := http.Post(loginUrl, contentType, strings.NewReader(postData.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return "", errors.New(fmt.Sprintf("[ERROR] (Login) Request Error: %v", err))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var loginModel model.LoginResp
	err = json.Unmarshal(body, &loginModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("[ERROR] (Login) Resp Json Unmarshal Error: %v", err))
	}
	ticket := loginModel.Data.Ticket
	return ticket, nil
}

//
// Upload2Azure
//  @Description: 上传图片到Azure服务器
//  @param auth
//  @param images
//  @return []string
//
func Upload2Azure(auth string, images map[string]string) ([]string, error) {
	var client = &http.Client{}
	uploadUrl := "https://p.luxshare-ict.com/api/Azure/TencentFileToAzure"
	contentType := "application/x-www-form-urlencoded"

	postData := url.Values{}
	postData.Add("dir", "~/Upload2/EpidemicSys/")
	postData.Add("mediaIds[]", fmt.Sprintf("data:image/jpg;base64,%v", images["xcm"]))
	postData.Add("mediaIds[]", fmt.Sprintf("data:image/jpg;base64,%v", images["jkm"]))
	postData.Add("source", "qyweixin")

	request, err := http.NewRequest("POST", uploadUrl, strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[ERROR] (Upload2Azure) Error creating upload request: %v", err))
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("__user__", `%7B%22CompanyOwner%22:888,%22CompanyCode%22:%22KSAT%22,%22CompanyName%22:%22%E6%B1%9F%E8%8B%8F%E6%9C%BA%E5%99%A8%E4%BA%BA%22,%22BUCode%22:%22U00001%22,%22BUName%22:%22%E6%99%BA%E8%83%BD%E5%88%B6%E9%80%A0%E5%BC%80%E5%8F%91%E4%B8%AD%E5%BF%83%22,%22DeptCode%22:%22U12544%22,%22DeptName%22:%22%E8%AE%BE%E5%A4%87%E4%BF%A1%E6%81%AF%E5%8C%96%E8%AF%BE%22,%22Code%22:%2213901424%22,%22Name%22:%22%E5%AE%8B%E5%A9%89%E5%86%9B%22,%22IDCardNo%22:%22340826199808161410%22,%22Gender%22:%22M%22,%22Telephone%22:%2217855513383%22,%22Email%22:%22Wanjun.Song@luxshare-ict.com%22,%22Language%22:%22zh-cn%22,%22LoginType%22:4,%22DataSource%22:%22M%22%7D`)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)
	//resp, err := http.Post(uploadUrl, contentType, strings.NewReader(postData.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[ERROR] (Upload2Azure) Request Error: %v", err))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var uploadModel model.Upload2AzureResp
	err = json.Unmarshal(body, &uploadModel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[ERROR] (Upload2Azure) Resp Json Unmarshal Error: %v", err))
	}
	return uploadModel.Data.ImagePaths, nil
}

//
// EpidemicRegistration
//  @Description: 申报
//  @param auth
//  @param images
//  @return error
//
func EpidemicRegistration(auth string, images []string) error {
	var client = &http.Client{}
	uploadUrl := "https://m.luxshare-ict.com/api/EpidemicSys/EpidemicRegistration/LVIQuestSave2"
	contentType := "application/x-www-form-urlencoded"

	postData := url.Values{}
	postData.Add("nowAddress", "鑫河湾")
	postData.Add("street", "锦溪镇")
	postData.Add("isRisk", "否")
	postData.Add("isWork", "是")
	postData.Add("isRoom", "否")
	postData.Add("isRiskContact1", "否")
	postData.Add("isRiskContact2", "否")
	postData.Add("isDays14", "否")
	postData.Add("isSymptom", "否")
	postData.Add("isVaccination", "是")
	postData.Add("vaccinationCount", "3")
	postData.Add("imagePaths[]", images[0])
	postData.Add("imagePaths[]", images[1])
	postData.Add("residCity", "江苏省+苏州市+昆山市")
	postData.Add("vaccineDt", "2021/05/13")
	postData.Add("vaccineDt2", "2021/06/11")
	postData.Add("vaccineDt3", "2021/12/18")
	postData.Add("noVaccineReason", "")
	postData.Add("healthCodeColor", "绿色")
	postData.Add("vaccinType", "三针型")
	postData.Add("temperature", "")
	postData.Add("isGtPerson", "")
	postData.Add("gtAddress", "")
	postData.Add("gtShenzhenTime", "")
	postData.Add("gtDongguanTime", "")
	postData.Add("dormArea", "")
	postData.Add("dormBuilding", "")
	postData.Add("dormRoom", "")
	postData.Add("byCity", "江苏省苏州市")
	postData.Add("latitude", "31.17725")
	postData.Add("longitude", "120.9165")
	postData.Add("registLocal", "江苏省苏州市昆山市学苑路")
	postData.Add("isXinGuanHistory", "")
	postData.Add("historyDate", "")
	postData.Add("isForeigner", "")
	postData.Add("isEntry", "")
	postData.Add("entryDate", "")
	postData.Add("isTravelToShangHai", "")

	request, err := http.NewRequest("POST", uploadUrl, strings.NewReader(postData.Encode()))
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (EpidemicRegistration) Error creating upload request: %v", err))
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("__user__", `%7B%22CompanyOwner%22:888,%22CompanyCode%22:%22KSAT%22,%22CompanyName%22:%22%E6%B1%9F%E8%8B%8F%E6%9C%BA%E5%99%A8%E4%BA%BA%22,%22BUCode%22:%22U00001%22,%22BUName%22:%22%E6%99%BA%E8%83%BD%E5%88%B6%E9%80%A0%E5%BC%80%E5%8F%91%E4%B8%AD%E5%BF%83%22,%22DeptCode%22:%22U12544%22,%22DeptName%22:%22%E8%AE%BE%E5%A4%87%E4%BF%A1%E6%81%AF%E5%8C%96%E8%AF%BE%22,%22Code%22:%2213901424%22,%22Name%22:%22%E5%AE%8B%E5%A9%89%E5%86%9B%22,%22IDCardNo%22:%22340826199808161410%22,%22Gender%22:%22M%22,%22Telephone%22:%2217855513383%22,%22Email%22:%22Wanjun.Song@luxshare-ict.com%22,%22Language%22:%22zh-cn%22,%22LoginType%22:4,%22DataSource%22:%22M%22%7D`)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)

	defer resp.Body.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (EpidemicRegistration) Request Error: %v", err))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var epidemicRegistrationModel model.UniverseResp
	err = json.Unmarshal(body, &epidemicRegistrationModel)
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (EpidemicRegistration) Resp Json Unmarshal Error: %v", err))
	}

	if !epidemicRegistrationModel.IsSuccess {
		return errors.New(fmt.Sprintf("[ERROR] (EpidemicRegistration) Resp ErrMsg: %v", epidemicRegistrationModel.ErrMsg))
	}
	return nil
}

//
// RefreshDoor
//  @Description: 刷新门禁
//  @param auth
//  @return error
//
func RefreshDoor(auth string) error {
	var client = &http.Client{}
	refreshUrl := "https://m.luxshare-ict.com/api/EpidemicSys/EpidemicRegistration/RefreshDoor"

	request, err := http.NewRequest("POST", refreshUrl, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (RefreshDoor) Error creating upload request: %v", err))
	}
	request.Header.Set("__user__", `%7B%22CompanyOwner%22:888,%22CompanyCode%22:%22KSAT%22,%22CompanyName%22:%22%E6%B1%9F%E8%8B%8F%E6%9C%BA%E5%99%A8%E4%BA%BA%22,%22BUCode%22:%22U00001%22,%22BUName%22:%22%E6%99%BA%E8%83%BD%E5%88%B6%E9%80%A0%E5%BC%80%E5%8F%91%E4%B8%AD%E5%BF%83%22,%22DeptCode%22:%22U12544%22,%22DeptName%22:%22%E8%AE%BE%E5%A4%87%E4%BF%A1%E6%81%AF%E5%8C%96%E8%AF%BE%22,%22Code%22:%2213901424%22,%22Name%22:%22%E5%AE%8B%E5%A9%89%E5%86%9B%22,%22IDCardNo%22:%22340826199808161410%22,%22Gender%22:%22M%22,%22Telephone%22:%2217855513383%22,%22Email%22:%22Wanjun.Song@luxshare-ict.com%22,%22Language%22:%22zh-cn%22,%22LoginType%22:4,%22DataSource%22:%22M%22%7D`)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)

	defer resp.Body.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (RefreshDoor) Request Error: %v", err))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var refreshDoorResp model.UniverseResp
	err = json.Unmarshal(body, &refreshDoorResp)
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (RefreshDoor) Resp Json Unmarshal Error: %v", err))
	}

	if !refreshDoorResp.IsSuccess {
		return errors.New(fmt.Sprintf("[ERROR] (RefreshDoor) Resp ErrMsg: %v", refreshDoorResp.ErrMsg))
	}
	return nil
}
