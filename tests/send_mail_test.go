package tests

import (
	//"bytes"
	"fmt"
	"io"
	"testing"

	//"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func sendMail(to, subject, body string) error {

	from := "rautab0@gmail.com"
	smtpServer := "localhost:1025"
	c, err := smtp.Dial(smtpServer)
	if err != nil {
		return err
	}
	if err := c.Mail(from, nil); err != nil {
		return err
	}
	if err := c.Rcpt(to, nil); err != nil {
		return err
	}
	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = io.WriteString(wc, body)
	if err != nil {
		return nil
	}

	err = wc.Close()
	if err != nil {
		return nil
	}
	err = c.Quit()
	if err != nil {
		return nil
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
