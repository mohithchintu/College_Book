package mail

import (
	"bytes"
	"crypto/tls"
	"html/template"

	"github.com/mohithchintu/College_Book/backend/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	SenderMail string
	Username   string
	OTP        string
}

func MailService(data EmailData) error {
	d := gomail.NewDialer(config.AppConfig.Mail_Service, config.AppConfig.Mail_Port, config.AppConfig.Mail, config.AppConfig.Mail_Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", config.AppConfig.Mail)
	m.SetHeader("To", data.SenderMail)
	m.SetHeader("Subject", "Test Email")

	tmpl, err := template.ParseFiles("mail/templates/otp_template.html")
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		return err
	}

	m.SetBody("text/html", bodyBuffer.String())

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// func ExampleUsage() {
// 	data := EmailData{
// 		SenderMail: "onlyforgamesdc@gmail.com",
// 		Username:   "Chintu",
// 		OTP:        "456789",
// 	}
// 	if err := MailService(data); err != nil {
// 		log.Printf("Failed to send email: %v", err)
// 	}
// }
