package message

import (
	"fmt"
	"github.com/jordan-wright/email"
	log "github.com/sirupsen/logrus"
	"luxshare-daily-report/global"
	"net/smtp"
)

type Mail struct {
	Host     string   `json:"host" yaml:"host"`
	Protocol string   `json:"protocol" yaml:"protocol"`
	Port     int      `json:"port" yaml:"port"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	FromName string   `json:"from_name" yaml:"from_name"`
	To       []string `json:"to" yaml:"to"`
}

var m Mail

func initMail() {
	m = Mail{
		Host:     global.GLO_CONFIG.MailHost,
		Port:     global.GLO_CONFIG.MailPort,
		Username: global.GLO_CONFIG.MailUser,
		Password: global.GLO_CONFIG.MailPwd,
		FromName: global.GLO_CONFIG.MailFromName,
		To:       global.GLO_CONFIG.MailTo,
	}
	Register("mail", m)
}

func (m Mail) Send(message Body) {
	log.Println("[mail] Sending by mail...")
	e := email.NewEmail()
	e.From = m.FromName
	e.To = m.To
	e.Subject = message.Title
	e.Text = []byte(message.Content)
	addr := fmt.Sprintf("%v:%v", m.Host, m.Port)
	err := e.Send(addr, smtp.PlainAuth("", m.Username, m.Password, m.Host))
	if err != nil {
		log.Fatalf("[mail] Send failed: %v\n", err)
	}
	log.Printf("[mail] Send successful")
}
