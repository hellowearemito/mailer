package mailer_test

import (
	"testing"

	. "chatbot/rest-api/components/mailer"

	. "github.com/smartystreets/goconvey/convey"
)

//go:generate counterfeiter -o fakesender_test.go -pkg ${GOPACKAGE}_test github.com/nkovacs/gophermail.Sender

func TestMailer(t *testing.T) {
	Convey("Mailer service", t, func() {
		Convey("Should be able to create a mailer", func() {
			mailer := NewSMTPMailer("noreply@chatbot.example.org", ":1025", false)
			So(mailer, ShouldNotBeNil)
			/*
				mess, err := mailer.NewPlainTextMessage("noreply@chatbot.example.org", []string{"tester@chatbot.example.org"}, "Test message", "Test body")
				So(err, ShouldBeNil)
				So(mess.Send(), ShouldBeNil)
			*/
		})
		Convey("Should be able to create a mailer with auth", func() {
			mailer := NewSMTPMailerWithAuth("noreply@chatbot.example.org", ":1025", "", "chatbot@example.org", "password", "example.org", false)
			So(mailer, ShouldNotBeNil)
			/*
				mess, err := mailer.NewPlainTextMessage("noreply@chatbot.example.org", []string{"tester@chatbot.example.org"}, "Test message", "Test body")
				So(err, ShouldBeNil)
				So(mess.Send(), ShouldBeNil)
			*/
		})
		Convey("With fake sender", func() {
			sender := new(FakeSender)
			mailer := NewMailer("noreply@chatbot.example.org", sender)
			So(mailer, ShouldNotBeNil)
			Convey("should send plain text mail", func() {
				sender.SendMailReturns(nil)
				subject := "Test email plain text"
				body := "This is a test email."
				mess, err := mailer.NewPlainTextMessage([]string{"tester@chatbot.example.org"}, subject, body)
				So(err, ShouldBeNil)
				err = mess.Send()
				So(err, ShouldBeNil)
				sentMessage := sender.SendMailArgsForCall(0)
				So(sentMessage.From.Name, ShouldEqual, "")
				So(sentMessage.From.Address, ShouldEqual, "noreply@chatbot.example.org")
				So(sentMessage.To, ShouldHaveLength, 1)
				So(sentMessage.To[0].Name, ShouldEqual, "")
				So(sentMessage.To[0].Address, ShouldEqual, "tester@chatbot.example.org")
				So(sentMessage.Subject, ShouldEqual, subject)
				So(sentMessage.Body, ShouldEqual, body)
			})

			Convey("should send html mail", func() {
				sender.SendMailReturns(nil)
				subject := "Test email plain text"
				body := "This is a <em>test</em> email."
				mess, err := mailer.NewHTMLMessage([]string{"tester@chatbot.example.org"}, subject, body)
				So(err, ShouldBeNil)
				err = mess.Send()
				So(err, ShouldBeNil)
				sentMessage := sender.SendMailArgsForCall(0)
				So(sentMessage.From.Name, ShouldEqual, "")
				So(sentMessage.From.Address, ShouldEqual, "noreply@chatbot.example.org")
				So(sentMessage.To, ShouldHaveLength, 1)
				So(sentMessage.To[0].Name, ShouldEqual, "")
				So(sentMessage.To[0].Address, ShouldEqual, "tester@chatbot.example.org")
				So(sentMessage.Subject, ShouldEqual, subject)
				So(sentMessage.HTMLBody, ShouldEqual, body)
			})

			Convey("should support changing body and subject", func() {
				sender.SendMailReturns(nil)
				subject := "Test email plain text"
				body := "This is a test email."
				subject2 := "Test email plain text modified"
				body2 := "This is a test email modified."
				mess, err := mailer.NewPlainTextMessage([]string{"tester@chatbot.example.org"}, subject, body)
				So(err, ShouldBeNil)
				mess.SetSubject(subject2)
				mess.SetBody(body2)
				err = mess.Send()
				So(err, ShouldBeNil)
				sentMessage := sender.SendMailArgsForCall(0)
				So(sentMessage.From.Name, ShouldEqual, "")
				So(sentMessage.From.Address, ShouldEqual, "noreply@chatbot.example.org")
				So(sentMessage.To, ShouldHaveLength, 1)
				So(sentMessage.To[0].Name, ShouldEqual, "")
				So(sentMessage.To[0].Address, ShouldEqual, "tester@chatbot.example.org")
				So(sentMessage.Subject, ShouldEqual, subject2)
				So(sentMessage.Body, ShouldEqual, body2)
			})

			Convey("should support adding html body to plain text message", func() {
				sender.SendMailReturns(nil)
				subject := "Test email plain text"
				body := "This is a test email."
				body2 := "This is a <em>test</em> email."
				mess, err := mailer.NewPlainTextMessage([]string{"tester@chatbot.example.org"}, subject, body)
				So(err, ShouldBeNil)
				mess.SetHTMLBody(body2)
				err = mess.Send()
				So(err, ShouldBeNil)
				sentMessage := sender.SendMailArgsForCall(0)
				So(sentMessage.From.Name, ShouldEqual, "")
				So(sentMessage.From.Address, ShouldEqual, "noreply@chatbot.example.org")
				So(sentMessage.To, ShouldHaveLength, 1)
				So(sentMessage.To[0].Name, ShouldEqual, "")
				So(sentMessage.To[0].Address, ShouldEqual, "tester@chatbot.example.org")
				So(sentMessage.Subject, ShouldEqual, subject)
				So(sentMessage.Body, ShouldEqual, body)
				So(sentMessage.HTMLBody, ShouldEqual, body2)
			})

			Convey("should parse email addresses", func() {
				sender.SendMailReturns(nil)
				subject := "Test email plain text"
				body := "This is a test email."
				mess, err := mailer.NewPlainTextMessage([]string{"Test Testersson <tester@chatbot.example.org>"}, subject, body)
				So(err, ShouldBeNil)
				err = mess.SetFrom("Chatbot <noreply@chatbot.example.org>")
				So(err, ShouldBeNil)
				err = mess.Send()
				So(err, ShouldBeNil)
				sentMessage := sender.SendMailArgsForCall(0)
				So(sentMessage.From.Name, ShouldEqual, "Chatbot")
				So(sentMessage.From.Address, ShouldEqual, "noreply@chatbot.example.org")
				So(sentMessage.To, ShouldHaveLength, 1)
				So(sentMessage.To[0].Name, ShouldEqual, "Test Testersson")
				So(sentMessage.To[0].Address, ShouldEqual, "tester@chatbot.example.org")
				So(sentMessage.Subject, ShouldEqual, subject)
				So(sentMessage.Body, ShouldEqual, body)
			})

			Convey("should error on invalid sender address", func() {
				sender.SendMailReturns(nil)
				mess, err := mailer.NewPlainTextMessage([]string{"tester@chatbot.example.org"}, "Error", "error")
				So(err, ShouldBeNil)
				err = mess.SetFrom("noreply_chatbot.example.org")
				So(err, ShouldNotBeNil)
			})

			Convey("should error on invalid default sender address", func() {
				mailer := NewMailer("noreply_chatbot.example.org", sender)
				So(mailer, ShouldNotBeNil)
				sender.SendMailReturns(nil)
				_, err := mailer.NewPlainTextMessage([]string{"tester@chatbot.example.org"}, "Error", "error")
				So(err, ShouldNotBeNil)
			})

			Convey("should error on invalid recipient address", func() {
				sender.SendMailReturns(nil)
				_, err := mailer.NewPlainTextMessage([]string{"tester_chatbot.example.org"}, "Error", "error")
				So(err, ShouldNotBeNil)
			})

		})
	})
}
