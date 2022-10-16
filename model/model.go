package model

// LoginResp
//  @Description: 登陆返回Model
//
type LoginResp struct {
	IsSuccess bool          `json:"IsSuccess"`
	ErrMsg    interface{}   `json:"ErrMsg"`
	Data      LoginRespData `json:"Data"`
}

type LoginRespData struct {
	IsLeaveLogin bool     `json:"IsLeaveLogin"`
	IsOnJob      bool     `json:"IsOnJob"`
	UserInfo     UserInfo `json:"UserInfo"`
	LoginTime    string   `json:"LoginTime"`
	Ticket       string   `json:"Ticket"`
}

type UserInfo struct {
	CompanyOwner int    `json:"CompanyOwner"`
	CompanyCode  string `json:"CompanyCode"`
	CompanyName  string `json:"CompanyName"`
	BUCode       string `json:"BUCode"`
	BUName       string `json:"BUName"`
	DeptCode     string `json:"DeptCode"`
	DeptName     string `json:"DeptName"`
	EmpCode      string `json:"EmpCode"`
	EmpName      string `json:"EmpName"`
	Gender       string `json:"Gender"`
	IDCardNo     string `json:"IDCardNo"`
	Telephone    string `json:"Telephone"`
	Email        string `json:"Email"`
	LoginType    int    `json:"LoginType"`
	DataSource   int    `json:"DataSource"`
	Code         string `json:"Code"`
	Name         string `json:"Name"`
	IdCardNo     string `json:"IdCardNo"`
	IsOnJob      bool   `json:"IsOnJob"`
}

//
// Upload2AzureResp
//  @Description: 上传图片返回Model
//
type Upload2AzureResp struct {
	IsSuccess bool                 `json:"IsSuccess"`
	ErrMsg    interface{}          `json:"ErrMsg"`
	Data      Upload2AzureRespData `json:"Data"`
}
type Upload2AzureRespData struct {
	ImagePaths []string `json:"ImagePaths"`
	ByCity     string   `json:"ByCity"`
	IsOk       bool     `json:"IsOk"`
}

//
// UniverseResp
//  @Description: 申报、刷新门禁返回Model
//
type UniverseResp struct {
	IsSuccess bool        `json:"IsSuccess"`
	ErrMsg    string      `json:"ErrMsg"`
	Data      interface{} `json:"Data"`
}

type BarkResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
