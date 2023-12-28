package util

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"videoweb/config"
)

func SendEmail(userEmail string, imgPath string) error {
	conf := config.Config.Email
	e := email.NewEmail()
	e.From = conf.Sender
	e.To = []string{userEmail}
	e.Subject = "验证码请扫描"

	_, err := e.AttachFile(imgPath)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", e.From, conf.Password, conf.Host) // 身份验证

	err = e.Send(conf.Addr, auth) // 发送
	if err != nil {
		return err
	}
	return err
}
