package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "yangchen<yangchenyc@mail.ustc.edu.cn>"
	e.To = []string{
		"2215092273@qq.com",
	}

	e.Subject = "验证码发送测试"

	e.HTML = []byte("您的验证码:<b>123456</b>")
	err := e.SendWithTLS("mail.ustc.edu.cn:465", smtp.PlainAuth("", "yangchenyc@mail.ustc.edu.cn", "qFTKVz3AUGKk5Zb6", "mail.ustc.edu.cn"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "mail.ustc.edu.cn"})
	if err != nil {
		t.Error(err)
	}
}
