package mailer

import (
	"crypto/tls"
	"net/smtp"

	"github.com/nkovacs/gophermail"
)

// Service is a mailing service.
type Service interface {
	NewMessage(to []string, subject string, htmlBody string, body string) (Message, error)
	NewHTMLMessage(to []string, subject string, htmlBody string) (Message, error)
	NewPlainTextMessage(to []string, subject string, body string) (Message, error)
}

type mailer struct {
	senderAddress string
	sender        gophermail.Sender
}

func (m *mailer) sendMessage(message *gophermail.Message) error {
	return m.sender.SendMail(message)
}

// NewMessage creates a a new message. The sender address will be set to the mailer's
// default sender address, but can be changed using SetFrom.
func (m *mailer) NewMessage(to []string, subject string, htmlBody string, body string) (Message, error) {
	message := &message{
		Message: gophermail.Message{
			Subject:  subject,
			HTMLBody: htmlBody,
			Body:     body,
		},
		mailer: m,
	}

	err := message.SetFrom(m.senderAddress)
	if err != nil {
		return nil, err
	}
	err = message.AddTo(to...)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (m *mailer) NewHTMLMessage(to []string, subject string, htmlBody string) (Message, error) {
	return m.NewMessage(to, subject, htmlBody, "") // TODO autogenerate plain text body
}

func (m *mailer) NewPlainTextMessage(to []string, subject string, body string) (Message, error) {
	return m.NewMessage(to, subject, "", body)
}

// NewMailer returns a mailer that uses the given sender.
func NewMailer(defaultSenderAddress string, sender gophermail.Sender) Service {
	return &mailer{defaultSenderAddress, sender}
}

// NewSMTPMailer creates an smtp mailer without authentication.
func NewSMTPMailer(defaultSenderAddress string, smtpAddress string, insecureSkipVerify bool) Service {
	sender := gophermail.NewSMTPSender(smtpAddress, nil, &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	})
	return &mailer{defaultSenderAddress, sender}
}

// NewSMTPMailerWithAuth creates an smtp mailer with authentication.
// see PlainAuth in net/smtp for details on authentication parameters.
func NewSMTPMailerWithAuth(defaultSenderAddress string, smtpAddress, identity, username, password, host string, insecureSkipVerify bool) Service {
	sender := gophermail.NewSMTPSender(smtpAddress, smtp.PlainAuth(identity, username, password, host), &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	})
	return &mailer{defaultSenderAddress, sender}
}
