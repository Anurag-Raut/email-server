package tests

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func sendMail(to, subject, body string) error {
	from := "rautab0@gmail.com"
	smtpServer := "localhost:1025"
	reader := bytes.NewReader([]byte(body))
	err := smtp.SendMail(smtpServer, sasl.NewPlainClient("", "", ""), from, []string{to}, reader)
	if err != nil {
		return err
	}
	return nil
}

func TestSendMail(t *testing.T) {
	err := sendMail("anurag.raut.86@gmail.com", "Hello", "By byte")
	if err != nil {
		t.Fatalf("Failed to send email: %v", err)
	}
	fmt.Printf("Mail Sent successfully")
}
