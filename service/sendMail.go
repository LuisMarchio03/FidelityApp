package service

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, code string) error {
	// SMTP server configuration
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "your-email@gmail.com"
	senderPassword := "your-email-password"

	// Recipient email address
	recipientEmail := "recipient-email@example.com"

	// Message content
	subject := "Test email"
	body := "This is a test email sent from Golang."

	// Compose the email message
	message := "From: " + senderEmail + "\n" +
		"To: " + recipientEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Authentication
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Sending email
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}
	fmt.Println("Email sent successfully!")

	return nil
}
