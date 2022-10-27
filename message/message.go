package message

import "luxshare-daily-report/global"

var Messages = make(map[string]Message, 0)

type Message interface {
	Send(body Body)
}

type Body struct {
	Title   string
	Content string
}

// GetSupport 获取支持
func GetSupport() Message {
	initBark()
	initMail()
	//key := "bark"
	key := global.GLO_CONFIG.MsgType
	return Messages[key]
}

// Enabled 是否启用
func Enabled() bool {
	return global.GLO_CONFIG.MsgEnabled
}

// Register 注册
func Register(name string, message Message) {
	Messages[name] = message
}
