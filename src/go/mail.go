package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

var user = ""
var password = ""
var host = "smtp.163.com:25"
var to = ""
var subject = "test"


func SendToMail(body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}

	msg := []byte(
		"To: " + to + "\r\n" +
			"From: " + user + "\r\n" +
			"Subject: " + subject + "\r\n" +
			contentType + "\r\n" +
			"\r\n" +
			"" + body)

	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func main() {

	// 这里用的是反单引号，而非单引号；
	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")


	err := SendToMail(body, "html")

	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}
