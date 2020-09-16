package utils

import (
	"gopkg.in/gomail.v2"
)

var smtpCon *gomail.Dialer

// InitSMTPCon sets up global smtp con
func InitSMTPCon(stmpUser string, smtpPw string) {
	tempCon := gomail.NewDialer("smtp.gmail.com", 465, stmpUser, smtpPw)
	smtpCon = tempCon
}

// SendEmail Allow to send a mail through gmail's service with a gmail account
func SendEmail(to []string, cc []string, bcc []string, subject string, text string, html string, from string) error {
	mail := gomail.NewMessage()

	mail.SetHeader("From", from)
	mail.SetHeader("To", to...)
	mail.SetHeader("CC", cc...)
	mail.SetHeader("BCC", bcc...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", text)
	mail.SetBody("text/html", html)

	// Send the email to Bob, Cora and Dan.
	err := smtpCon.DialAndSend(mail)

	return err
}
