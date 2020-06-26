package utils

import (
	"net/smtp"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

// SendEmail Allow to send a mail through gmail's service with a gmail account
func SendEmail(to []string, cc []string, bcc []string, subject string, text string, html string, from string, attachmentURL string) error {
	env, _ := godotenv.Read()

	newMail := email.NewEmail()
	newMail.From = from
	newMail.To = to
	newMail.Bcc = bcc
	newMail.Cc = cc
	newMail.Subject = subject
	newMail.Text = []byte(text)
	newMail.HTML = []byte(html)

	newMail.AttachFile(attachmentURL)

	err := newMail.Send(env["SMTP_HOST"]+env["SMTP_PORT"], smtp.PlainAuth("", env["SMTP_USER"], env["SMTP_PASSWORD"], env["SMTP_HOST"]))

	return err
}
