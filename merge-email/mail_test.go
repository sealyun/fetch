package main

import (
	"fmt"
	"testing"
)

func TestSendTemplateMail(t *testing.T) {
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
		`, "cuisongliu")
	SendHtmlMail("user_key", body, "cuisongliu@qq.com", "Install kubernetes HA in one step!")
}
