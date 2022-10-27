package global

import (
	"luxshare-daily-report/model"
)

var (
	GLO_CONFIG      model.Config
	GLO_CONFIG_CHAN chan model.Config
	GLO_RECV_CHAN   chan map[string]string
)
