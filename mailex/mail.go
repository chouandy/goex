package mailex

import (
	"strings"

	"gopkg.in/gomail.v2"
)

// Mail message struct
type Mail struct {
	To          string
	Subject     string
	ContentType string
	Body        string
	Attachments []string
}

// Send send mail
func (m *Mail) Send() error {
	// New email message
	msg := gomail.NewMessage()
	// Set from
	msg.SetHeader("From", Mailer.Options.From)
	// Set email to
	msg.SetHeader("To", strings.Split(m.To, ",")...)
	// Set email subject
	msg.SetHeader("Subject", m.Subject)
	// Set email body
	msg.SetBody(m.ContentType, m.Body)
	// Set attachment
	for _, attachment := range m.Attachments {
		msg.Attach(attachment)
	}

	// New dialer
	d := gomail.NewDialer(
		Mailer.SMTPSettings.Address,
		Mailer.SMTPSettings.Port,
		Mailer.SMTPSettings.Username,
		Mailer.SMTPSettings.Password,
	)

	return d.DialAndSend(msg)
}
