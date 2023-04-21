package service

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to string, code string) error {
	from := "john.doe@example.com"

	user := "9c1d45eaf7af5b"
	password := "ad62926fa75d0f"

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	msg := []byte("Code: " + code + " to verify your email address")

	auth := smtp.PlainAuth("", user, password, host)

	arrTo := []string{to}
	err := smtp.SendMail(addr, auth, from, arrTo, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
	fmt.Println("Code: " + code)

	return nil
}
