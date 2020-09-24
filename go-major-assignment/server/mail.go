package main

import (
	"log"
	"net/smtp"
)

func send(from, to, subject, message string) {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:" + message + "\n\n" + message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print(msg)
}
