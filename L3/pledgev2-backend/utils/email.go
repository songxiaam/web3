package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
	"pledgev2-backend/config"
)

func SendEmail(data []byte, dataType int) error {
	e := &email.Email{
		To:      config.Config.Email.To,
		Cc:      config.Config.Email.Cc,
		From:    config.Config.Email.From,
		Subject: config.Config.Email.Subject,
		Headers: textproto.MIMEHeader{},
	}
	if dataType == 1 {
		e.Text = data
	} else {
		e.HTML = data
	}
	return e.Send(config.Config.Email.Host+":"+config.Config.Email.Port, smtp.PlainAuth("", config.Config.Email.UserName, config.Config.Email.Password, config.Config.Email.Host))
}

// SendEmailWithAttach dataType 1 text, 2 html
func SendEmailWithAttach(data []byte, dataType int, filename string) error {
	e := &email.Email{
		To:      config.Config.Email.To,
		Cc:      config.Config.Email.Cc,
		From:    config.Config.Email.From,
		Subject: config.Config.Email.Subject,
		Headers: textproto.MIMEHeader{},
	}
	if dataType == 1 {
		e.Text = data
	} else {
		e.HTML = data
	}
	_, err := e.AttachFile(filename)
	if err != nil {
		return err
	}
	return e.Send(config.Config.Email.Host+config.Config.Email.Port, smtp.PlainAuth("", config.Config.Email.UserName, config.Config.Email.Password, config.Config.Email.Host))
}
