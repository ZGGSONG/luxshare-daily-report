package core

import (
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"receive-files/model"
	"strings"
)

func Login(userName, passwd string) string {
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
	if err != nil {
		log.Printf("[LOGIN] Request Error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))
	var loginModel model.LoginResp
	err = json.Unmarshal(body, &loginModel)
	if err != nil {
		log.Printf("[LOGIN] Resp Json Unmarshal Error: %v", err)
	}
	ticket := loginModel.Data.Ticket
	return ticket
}
