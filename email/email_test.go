package email

import (
	"fmt"
	"testing"
)

func Test_SendEmail(t *testing.T) {
	var config = &MailConfig{}
	config.Username = "developer_mail@163.com"
	config.Host = "smtp.163.com"
	config.Password = "rkrntactzdinzcjk"
	config.Port = "25"
	config.Secure = false

	var e = NewHtmlMessage("title", "<a href='http://www.google.com'>Google</a>")
	e.From = "From<developer_mail@163.com>"
	e.To = []string{"917996695@qq.com"}

	fmt.Println(SendMail(config, e))
}