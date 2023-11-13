package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
	"study_gva/plugin/email/global"

	"github.com/jordan-wright/email"
)

// Email测试方法
func EmailTest(subject string, body string) error {
	to := []string{global.GlobalConfig.To}
	return send(to, subject, body)
}

// Email发送方法
func send(to []string, subject, body string) error {
	from := global.GlobalConfig.From
	nickname := global.GlobalConfig.Nickname
	secret := global.GlobalConfig.Secret
	host := global.GlobalConfig.Host
	port := global.GlobalConfig.Port
	isSSL := global.GlobalConfig.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

// Email发送方法
func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}
