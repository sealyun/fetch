package main

import (
	"bufio"
	"fmt"
	"io"
	"net/smtp"
	"os"
	"strings"
)

//Email  is
var Email map[string][]string

var sended map[string]int

func loadSended() {
	fi, err := os.Open("sended.dump")
	if err != nil {
		fmt.Println("open file error ", err)
	}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		sended[string(a)] = 1
	}
}

func alreadySend(email string) bool {
	_, ok := sended[email]
	return ok
}

func main() {
	Email = make(map[string][]string)
	sended = make(map[string]int)
	loadSended()

	if os.Getenv("API_USER") != "" {
		API_USER = os.Getenv("API_USER")
	}
	if os.Getenv("FROM") != "" {
		FROM = os.Getenv("FROM")
	}
	if os.Getenv("FROM_USER") != "" {
		FROM_USER = os.Getenv("FROM_USER")
	}
	if os.Getenv("KEY") != "" {
		KEY = os.Getenv("KEY")
	}

	fmt.Printf("api user: %s from: %s from user: %s key: %s\n", API_USER, FROM, FROM_USER, KEY)

	for _, file := range os.Args[1:] {
		fi, err := os.Open(file)
		if err != nil {
			fmt.Println("open file error ", err)
		}

		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			if contains(string(a)) {
				continue
			}

			sp := strings.Split(string(a), "|")
			email := sp[4]
			if email != "null" {
				Email[email] = sp
				fmt.Printf("dump merge : %s\n", string(a))
			}
		}
	}

	for k, v := range Email {
		if alreadySend(k) {
			fmt.Println("dump already send: ", k)
			continue
		}
		//send(v[0], k)
		sendCloud(v[0], k)
	}
}

func contains(a string) bool {
	b := strings.ToLower(a)
	return strings.Contains(b, "iflytek")  || strings.Contains(b, "spam")|| strings.Contains(b, "hefei") || strings.Contains(b, "anhui")
}

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

//using sendCloud to send email
func sendCloud(name, email string) {
/*
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
		`, name)
*/
	body := fmt.Sprintf(`
    		<html>
		<body>
		<h3>
		Dear %s,
		</h3>
		<p>
&nbsp;&nbsp;&nbsp;&nbsp;      I am steven, a coder from China, I am honoured to introduce a kubernetes HA installer for you. 
https://github.com/fanux/sealos , 
 <br /> &nbsp;&nbsp;&nbsp;&nbsp; you can build a high avalible kubernetes cluster in few minute using sealos. 
We using ipvs local proxy for masters LB, it is very simple and robust. 
<br />
     &nbsp;&nbsp;&nbsp;&nbsp;   Hope you like it, sorry about to disturb you :)

<br />
Regards,
<br />
Steven.

        <br />
		</p>
		</body>
		</html>
	`,name)

	SendHtmlMail(KEY, email, "A kubernetes HA installer", body)
}

func send(name, email string) {
	//user := "474785153@qq.com"
	//password := "opprovzrnltjbjid" // 需要在邮箱设置里配置
	user := "sealyun@cuisongliu.com"
	password := "wajwdkxxnqrfbjfh"

	host := "smtp.qq.com:25"
	to := email

	subject := ""

	githubUser := name

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
		fmt.Printf("Failed|%s|%s\n", name, email)
		fmt.Println(err)
	} else {
		fmt.Printf("Success|%s|%s\n", name, email)
	}
}
