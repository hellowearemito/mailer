package mailer

import "github.com/nkovacs/gophermail"

// Message is an e-mail message.
type Message interface {
	AddBcc(addresses ...string) error
	AddCc(addresses ...string) error
	AddTo(addresses ...string) error
	Bytes() ([]byte, error)
	SetFrom(address string) error
	SetReplyTo(address string) error
	SetBody(body string)
	SetHTMLBody(body string)
	SetSubject(subject string)
	// TODO attachment
	Send() error
}

type message struct {
	gophermail.Message
	mailer *mailer
}

// SetBody
func (m *message) SetBody(body string) {
	m.Body = body
}

//
func (m *message) SetHTMLBody(body string) {
	m.HTMLBody = body
}

func (m *message) SetSubject(subject string) {
	m.Subject = subject
}

func (m *message) Send() error {
	return m.mailer.sendMessage(&m.Message)
}
