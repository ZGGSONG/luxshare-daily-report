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

//
// Config
//  @Description: config model
//
type Config struct {
	Quality      string
	UserName     string
	PassWord     string
	MsgEnabled   bool
	MsgType      string
	BarkUrl      string
	BarkKey      string
	MailHost     string
	MailProtocol string
	MailPort     int
	MailUser     string
	MailPwd      string
	MailFromName string
	MailTo       []string
}

//
// HeaderUser
//  @Description: Header User Model
//
type HeaderUser struct {
	CompanyOwner int    `json:"CompanyOwner"`
	CompanyCode  string `json:"CompanyCode"`
	CompanyName  string `json:"CompanyName"`
	BUCode       string `json:"BUCode"`
	BUName       string `json:"BUName"`
	DeptCode     string `json:"DeptCode"`
	DeptName     string `json:"DeptName"`
	Code         string `json:"Code"`
	Name         string `json:"Name"`
	IDCardNo     string `json:"IDCardNo"`
	Gender       string `json:"Gender"`
	Telephone    string `json:"Telephone"`
	Email        string `json:"Email"`
	Language     string `json:"Language"`
	LoginType    int    `json:"LoginType"`
	DataSource   string `json:"DataSource"`
}

//
// LVIQuestInitModel
//  @Description: 获取以前填写过的信息
//
type LVIQuestInitModel struct {
	IsSuccess bool                  `json:"IsSuccess"`
	ErrMsg    string                `json:"ErrMsg"`
	Data      LVIQuestInitModelData `json:"Data"`
}

type LVIQuestInitModelData struct {
	LviImages         []string         `json:"LviImages"`
	EpidemicQuestLVI  EpidemicQuestLVI `json:"EpidemicQuestLVI"`
	NoVaccineReasons  []string         `json:"NoVaccineReasons"`
	HealthCodeColors  []string         `json:"HealthCodeColors"`
	Contacts          []interface{}    `json:"Contacts"`
	VaccinTypeList    []string         `json:"VaccinTypeList"`
	VaccinType        string           `json:"VaccinType"`
	IsLuxSan          bool             `json:"IsLuxSan"`
	IsDongGuan        bool             `json:"IsDongGuan"`
	IsTodaySubmit     bool             `json:"IsTodaySubmit"`
	IsKeyDorm         bool             `json:"IsKeyDorm"`
	DormAreas         []interface{}    `json:"DormAreas"`
	Towns             []string         `json:"Towns"`
	CanRefresh        bool             `json:"CanRefresh"`
	ProvinceList      []string         `json:"ProvinceList"`
	QuList            []QuList         `json:"QuList"`
	IsLz              bool             `json:"IsLz"`
	EpidemicDayConfig interface{}      `json:"EpidemicDayConfig"`
}
type QuList struct {
	HighAreaProvince string `json:"HighAreaProvince"`
	HighAreaQu       string `json:"HighAreaQu"`
}

type EpidemicQuestLVI struct {
	Id                   string `json:"Id"`
	RegistrationTime     string `json:"RegistrationTime"`
	CompanyCode          string `json:"CompanyCode"`
	CompanyName          string `json:"CompanyName"`
	DeptCode             string `json:"DeptCode"`
	DeptName             string `json:"DeptName"`
	Line                 string `json:"Line"`
	Telephone            string `json:"Telephone"`
	EmpCode              string `json:"EmpCode"`
	EmpName              string `json:"EmpName"`
	IdentityCard         string `json:"IdentityCard"`
	NowAddress           string `json:"NowAddress"`
	Street               string `json:"Street"`
	IsRisk               string `json:"IsRisk"`
	IsWork               string `json:"IsWork"`
	IsRoom               string `json:"IsRoom"`
	IsRiskContact1       string `json:"IsRiskContact1"`
	IsRiskContact2       string `json:"IsRiskContact2"`
	IsDays14             string `json:"IsDays14"`
	IsSymptom            string `json:"IsSymptom"`
	IsVaccination        string `json:"IsVaccination"`
	VaccinationCount     int    `json:"VaccinationCount"`
	Contacts             string `json:"Contacts"`
	ResidCity            string `json:"ResidCity"`
	VaccineDt            string `json:"VaccineDt"`
	VaccineDt2           string `json:"VaccineDt2"`
	VaccineDt3           string `json:"VaccineDt3"`
	NoVaccineReason      string `json:"NoVaccineReason"`
	HealthCodeColor      string `json:"HealthCodeColor"`
	IsDisabled           bool   `json:"IsDisabled"`
	Image1               string `json:"Image1"`
	Image2               string `json:"Image2"`
	Image3               string `json:"Image3"`
	IsOutSidePerson      bool   `json:"IsOutSidePerson"`
	IsNewPerson          bool   `json:"IsNewPerson"`
	VaccinType           string `json:"VaccinType"`
	Temperature          string `json:"Temperature"`
	IsGtPerson           string `json:"IsGtPerson"`
	GtAddress            string `json:"GtAddress"`
	GtShenzhenTime       string `json:"GtShenzhenTime"`
	GtDongguanTime       string `json:"GtDongguanTime"`
	StartCitys           string `json:"StartCitys"`
	IsStart              string `json:"IsStart"`
	JianKangColor        string `json:"JianKangColor"`
	IsPhoneSame          string `json:"IsPhoneSame"`
	UpPhone              string `json:"UpPhone"`
	IsOCR                bool   `json:"IsOCR"`
	OCRTime              string `json:"OCRTime"`
	ErrMsg               string `json:"ErrMsg"`
	DormArea             string `json:"DormArea"`
	DormBuilding         string `json:"DormBuilding"`
	DormRoom             string `json:"DormRoom"`
	IsDoor               bool   `json:"IsDoor"`
	IsDoorRecover        bool   `json:"IsDoorRecover"`
	DoorRecoverTime      string `json:"DoorRecoverTime"`
	DoorRecoverMsg       string `json:"DoorRecoverMsg"`
	IsSanDoor            bool   `json:"IsSanDoor"`
	SanLZCode            string `json:"SanLZCode"`
	IsSanDoorRecover     bool   `json:"IsSanDoorRecover"`
	SanDoorRecoverTime   string `json:"SanDoorRecoverTime"`
	ByCity               string `json:"ByCity"`
	Latitude             string `json:"Latitude"`
	Longitude            string `json:"Longitude"`
	RegistLocal          string `json:"RegistLocal"`
	EmpDormRoom          string `json:"EmpDormRoom"`
	IsXinGuanHistory     string `json:"IsXinGuanHistory"`
	HistoryDate          string `json:"HistoryDate"`
	IsForeigner          string `json:"IsForeigner"`
	IsEntry              string `json:"IsEntry"`
	EntryDate            string `json:"EntryDate"`
	IsTravelToShangHai   string `json:"IsTravelToShangHai"`
	IsEntryBefor         bool   `json:"IsEntryBefor"`
	EnteryBeforeResultNG bool   `json:"EnteryBeforeResultNG"`
	EntryBeforeNGRemark  string `json:"EntryBeforeNGRemark"`
	EpiFlag              bool   `json:"EpiFlag"`
	EpiDay               string `json:"EpiDay"`
	EpiCheck             string `json:"EpiCheck"`
	IsOnJob              bool   `json:"IsOnJob"`
	OnJobDate            string `json:"OnJobDate"`
	IsCheckHSOk          bool   `json:"IsCheckHSOk"`
	FengXianDiId         string `json:"FengXianDiId"`
}
