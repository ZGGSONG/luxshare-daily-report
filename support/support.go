package support

import (
	"fmt"
	"luxshare-daily-report/message"
)

// SendMessageError 发送错误信息
func SendMessageError(err error) {
	var body = message.Body{}
	body.Title = "LuxShareReportServer"
	body.Content = fmt.Sprintf("Error occurred: %s", err.Error())
	send(body)
}

// SendSuccess 发送成功信息
func SendSuccess(content string) {
	body := message.Body{
		Title:   "LuxShareReportServer",
		Content: content,
	}
	send(body)
}

// send 发送
func send(body message.Body) {
	m := message.GetSupport()
	if m == nil {
		return
	}
	enabled := message.Enabled()
	if enabled {
		m.Send(body)
	}
}
