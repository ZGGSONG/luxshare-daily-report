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
	"strconv"
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
		return "", "", errors.New(fmt.Sprintf("(Login) Request Error: %v", err))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var loginModel model.LoginResp
	err = json.Unmarshal(body, &loginModel)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("(Login) Resp Json Unmarshal Error: %v", err))
	}
	//失败
	if !loginModel.IsSuccess {
		err = errors.New(util.Strval(loginModel.ErrMsg))
		return "", "", err
	}
	ticket := loginModel.Data.Ticket

	data, user := model.LoginRespData{}, model.UserInfo{}
	if loginModel.Data == data || loginModel.Data.UserInfo == user {
		err = errors.New(fmt.Sprintf("Login not return user information"))
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
		return nil, errors.New(fmt.Sprintf("(Upload2Azure) Error creating upload request: %v", err))
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("__user__", user)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("(Upload2Azure) Request Error: %v", err))
	}
	//resp, err := http.Post(uploadUrl, contentType, strings.NewReader(postData.Encode()))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var uploadModel model.Upload2AzureResp
	err = json.Unmarshal(body, &uploadModel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("(Upload2Azure) Resp Json Unmarshal Error: %v", err))
	}
	return uploadModel.Data.ImagePaths, nil
}

//
// GetLVIQuestInitModel
//  @Description: 获取以往个人信息
//  @param auth
//  @param user
//  @return model.LVIQuestInitModelData
//  @return error
//
func GetLVIQuestInitModel(auth, user string) (model.EpidemicQuestLVI, error) {
	var nilModel model.EpidemicQuestLVI
	var client = &http.Client{}
	url := "https://m.luxshare-ict.com/api/EpidemicSys/EpidemicRegistration/LVIQuestInit"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nilModel, errors.New(fmt.Sprintf("(GetLVIQuestInitModel) Error creating request: %v", err))
	}
	request.Header.Set("__user__", user)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)
	if err != nil {
		return nilModel, errors.New(fmt.Sprintf("(GetLVIQuestInitModel) Request Error: %v", err))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var livModel model.LVIQuestInitModel
	err = json.Unmarshal(body, &livModel)
	if err != nil {
		return nilModel, errors.New(fmt.Sprintf("(GetLVIQuestInitModel) Resp Json Unmarshal Error: %v", err))
	}
	return livModel.Data.EpidemicQuestLVI, nil
}

//
// EpidemicRegistration
//  @Description: 申报
//  @param auth
//  @param images
//  @return error
//
func EpidemicRegistration(auth, user string, images []string, data model.EpidemicQuestLVI) error {
	var client = &http.Client{}
	uploadUrl := "https://m.luxshare-ict.com/api/EpidemicSys/EpidemicRegistration/LVIQuestSave2"
	contentType := "application/x-www-form-urlencoded"

	postData := url.Values{}
	postData.Add("nowAddress", data.NowAddress)
	postData.Add("street", data.Street)
	postData.Add("isRisk", data.IsRisk)
	postData.Add("isWork", data.IsWork)
	postData.Add("isRoom", data.IsRoom)
	postData.Add("isRiskContact1", data.IsRiskContact1)
	postData.Add("isRiskContact2", data.IsRiskContact2)
	postData.Add("isDays14", data.IsDays14)
	postData.Add("isSymptom", data.IsSymptom)
	postData.Add("isVaccination", data.IsVaccination)
	postData.Add("vaccinationCount", strconv.Itoa(data.VaccinationCount))
	postData.Add("imagePaths[]", images[0])
	postData.Add("imagePaths[]", images[1])
	postData.Add("residCity", data.ResidCity)
	postData.Add("vaccineDt", data.VaccineDt[:10])
	postData.Add("vaccineDt2", data.VaccineDt2[:10])
	postData.Add("vaccineDt3", data.VaccineDt3[:10])
	postData.Add("noVaccineReason", data.NoVaccineReason)
	postData.Add("healthCodeColor", data.HealthCodeColor)
	postData.Add("vaccinType", data.VaccinType)
	postData.Add("temperature", data.Temperature)
	postData.Add("isGtPerson", data.IsGtPerson)
	postData.Add("gtAddress", data.GtAddress)
	postData.Add("gtShenzhenTime", data.GtShenzhenTime)
	postData.Add("gtDongguanTime", data.GtDongguanTime)
	postData.Add("dormArea", data.DormArea)
	postData.Add("dormBuilding", data.DormBuilding)
	postData.Add("dormRoom", data.DormRoom)
	postData.Add("byCity", data.ByCity)
	postData.Add("latitude", data.Latitude)
	postData.Add("longitude", data.Longitude)
	postData.Add("registLocal", data.RegistLocal)
	postData.Add("isXinGuanHistory", data.IsXinGuanHistory)
	postData.Add("historyDate", data.HistoryDate)
	postData.Add("isForeigner", data.IsForeigner)
	postData.Add("isEntry", data.IsEntry)
	postData.Add("entryDate", data.EntryDate)
	postData.Add("isTravelToShangHai", data.IsTravelToShangHai)

	request, err := http.NewRequest("POST", uploadUrl, strings.NewReader(postData.Encode()))
	if err != nil {
		return errors.New(fmt.Sprintf("(EpidemicRegistration) Error creating upload request: %v", err))
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("__user__", user)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)

	defer resp.Body.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("(EpidemicRegistration) Request Error: %v", err))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var epidemicRegistrationModel model.UniverseResp
	err = json.Unmarshal(body, &epidemicRegistrationModel)
	if err != nil {
		return errors.New(fmt.Sprintf("(EpidemicRegistration) Resp Json Unmarshal Error: %v", err))
	}

	if !epidemicRegistrationModel.IsSuccess {
		return errors.New(fmt.Sprintf("(EpidemicRegistration) Resp ErrMsg: %v", epidemicRegistrationModel.ErrMsg))
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
		return errors.New(fmt.Sprintf("(RefreshDoor) Error creating upload request: %v", err))
	}
	request.Header.Set("__user__", user)
	request.Header.Set("Authorization", fmt.Sprintf("BaseAuth %v", auth))

	resp, err := client.Do(request)

	defer resp.Body.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("(RefreshDoor) Request Error: %v", err))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var refreshDoorResp model.UniverseResp
	err = json.Unmarshal(body, &refreshDoorResp)
	if err != nil {
		return errors.New(fmt.Sprintf("(RefreshDoor) Resp Json Unmarshal Error: %v", err))
	}

	if !refreshDoorResp.IsSuccess {
		return errors.New(fmt.Sprintf("(RefreshDoor) Resp ErrMsg: %v", refreshDoorResp.ErrMsg))
	}
	return nil
}
