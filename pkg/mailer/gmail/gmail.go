package gmail

import (
	"net/smtp"
	"os"
)

type Gmailer struct {
}

func (mailer Gmailer) Send(message string, to string) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	host := "smtp.gmail.com"
	port := "587"
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
