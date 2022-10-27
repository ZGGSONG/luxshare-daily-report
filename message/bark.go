package message

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"

	"luxshare-daily-report/global"
	"net/http"
)

type Bark struct {
	url string
	key string
}

var b = Bark{}

func initBark() {
	b.url = global.GLO_CONFIG.BarkUrl
	b.key = global.GLO_CONFIG.BarkKey
	Register("bark", b)
}

func (m Bark) Send(message Body) {
	log.Println("[bark] Sending by bark...")
	var reqBody = Request{
		DeviceKey: b.key,
		Title:     message.Title,
		Body:      message.Content,
		Icon:      "https://m.luxshare-ict.com/favicon.ico",
		//Url:       "https://github.com/zggsong",
	}
	req, _ := json.Marshal(reqBody)
	resp, err := http.Post(m.url, "application/json; charset=utf-8", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("[bark] http post failed: %v\n", err)
	}
	defer resp.Body.Close()
	log.Printf("[bark] Send successful")
}

type Request struct {
	Body      string `json:"body"`
	DeviceKey string `json:"device_key"`
	Title     string `json:"title"`
	Badge     int    `json:"badge"`
	Category  string `json:"category"`
	Sound     string `json:"sound"`
	Icon      string `json:"icon"`
	Group     string `json:"group"`
	Url       string `json:"url"`
}

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
