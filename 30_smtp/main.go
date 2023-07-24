package main

import "net/smtp"

func main() {
	auth := smtp.PlainAuth("", "9ilbert.app@gmail.com", "password", "smtp.gmail.com")

	from := "sender@gmail.com"
	to := []string{"receiver@gmail.com"}

	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다.\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
