package mailer

import "stori/pkg/mailer/gmail"

type Mailer interface {
	Send(message string, to string) error
}

func NewMailer() Mailer {
	// Replace by your own email flavor
	return gmail.Gmailer{}
}
