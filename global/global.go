package global

import "luxshare-daily-report/config"

var (
	GLO_CONFIG    config.Config
	GLO_RECV_CHAN chan map[string]string
)
