package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

//SendToMail is
func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: Install kubernetes HA in one step!" + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func main() {
	user := "474785153@qq.com"
	password := "opprovzrnltjbjid" // 需要在邮箱设置里配置
	host := "smtp.qq.com:25"
	to := "474785153@qq.com"

	subject := "使用Golang发送邮件"

	githubUser := "fanux"

	body := fmt.Sprintf(`
		<html>
		<body>
		<h3>
		Dear %s:
		</h3>
		<p>
		    Hi, It's my pleasure to introduce you a kubernetes HA install tool <a href="https://github.com/fanux/sealos"> sealos </a>, 

        <br />
        <br />
		Quick Start:
        <br />
        <br />
		<code>
		sealos init \<br />
			--master 192.168.0.2 \<br />
			--master 192.168.0.3 \<br />
			--master 192.168.0.4 \          # master addresses list <br />
			--node 192.168.0.5 \            # nodes list <br />
			--user root \                   # host username <br />
			--passwd your-server-password \ # host password <br />
			--pkg kube1.14.1.tar.gz  \      # offline package name, if you star sealos on github, you can download it free in http://store.lameleg.com <br />
			--version v1.14.1               # kubernetes version <br />
			</code>
        <br />
		That all!
        <br />

			Best wishes!
		</p>
		</body>
		</html>
		`, githubUser)
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}
