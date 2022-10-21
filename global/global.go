package global

import "luxshare-mrsb/config"

var (
	GLO_CONFIG    config.Config
	GLO_RECV_CHAN chan map[string]string
)
