package service

import (
	"fmt"
)

func SendEmail(to string, code string) error {
	fmt.Println("Sending email to: ", to)
	fmt.Println("Code: ", code)
	return nil
}
