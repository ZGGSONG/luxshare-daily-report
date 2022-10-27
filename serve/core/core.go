package core

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"luxshare-daily-report/model"
	"luxshare-daily-report/util"
	"net/http"
	"net/url"
	"strings"
)

//
// Login
//  @Description: 登陆网站
//  @param userName
//  @param passwd
//  @return string
//
func Login(userName, passwd string) (string, string, error) {
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
		return "", "", errors.New(fmt.Sprintf("[ERROR] (Login) Request Error: %v", err))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var loginModel model.LoginResp
	err = json.Unmarshal(body, &loginModel)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("[ERROR] (Login) Resp Json Unmarshal Error: %v", err))
	}
	//失败
	if !loginModel.IsSuccess {
		err = errors.New(util.Strval(loginModel.ErrMsg))
		return "", "", err
	}
	ticket := loginModel.Data.Ticket

	data, user := model.LoginRespData{}, model.UserInfo{}
	if loginModel.Data == data || loginModel.Data.UserInfo == user {
		err = errors.New(fmt.Sprintf("[ERROR] Login not return user information"))
		return ticket, "", err
	}
	userStr := convert(loginModel.Data.UserInfo)

	return ticket, userStr, nil
}

func convert(info model.UserInfo) string {
	user := model.HeaderUser{
		CompanyOwner: info.CompanyOwner,
		CompanyCode:  info.CompanyCode,
		CompanyName:  info.CompanyName,
		BUCode:       info.BUCode,
		BUName:       info.BUName,
		DeptCode:     info.DeptCode,
		DeptName:     info.DeptName,
		Code:         info.Code,
		Name:         info.Name,
		IDCardNo:     info.IDCardNo,
		Gender:       info.Gender,
		Telephone:    info.Telephone,
		Email:        info.Email,
		Language:     "zh-cn",
		LoginType:    info.LoginType,
		DataSource:   "M",
	}
	bytes, _ := json.Marshal(user)
	//url编码
	escape := url.QueryEscape(string(bytes))
	return escape
}

//
// Upload2Azure
//  @Description: 上传图片到Azure服务器
//  @param auth
//  @param images
//  @return []string
//
func Upload2Azure(auth, user string, images map[string]string) ([]string, error) {
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
	request.Header.Set("__user__", user)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[ERROR] (Upload2Azure) Request Error: %v", err))
	}
	//resp, err := http.Post(uploadUrl, contentType, strings.NewReader(postData.Encode()))
	defer resp.Body.Close()

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
func EpidemicRegistration(auth, user string, images []string) error {
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
	request.Header.Set("__user__", user)
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
func RefreshDoor(auth, user string) error {
	var client = &http.Client{}
	refreshUrl := "https://m.luxshare-ict.com/api/EpidemicSys/EpidemicRegistration/RefreshDoor"

	request, err := http.NewRequest("POST", refreshUrl, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] (RefreshDoor) Error creating upload request: %v", err))
	}
	request.Header.Set("__user__", user)
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
