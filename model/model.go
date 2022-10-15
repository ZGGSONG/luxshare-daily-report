package model

type LoginResp struct {
	IsSuccess bool        `json:"IsSuccess"`
	ErrMsg    interface{} `json:"ErrMsg"`
	Data      Data        `json:"Data"`
}

type Data struct {
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
