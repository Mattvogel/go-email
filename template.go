package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Email struct {
	To               string
	From             string
	InputArguments   interface{}
	Template         *template.Template
	RenderedTemplate string
}

func (e *Email) Render() error {
	var rdr bytes.Buffer
	err := e.Template.Execute(&rdr, e.InputArguments)
	if err != nil {
		return err
	}
	e.RenderedTemplate = rdr.String()
	// do something with the template
	return nil
}

func (e *Email) Send(server string) {
	// send the email
	c, err := smtp.Dial(server)
	if err != nil {
		// handle error
	}
	c.Mail(e.From)
	c.Rcpt(e.To)
	wc, err := c.Data()
	if err != nil {
		// handle error
	}
	_, err = wc.Write([]byte(e.RenderedTemplate))
	if err != nil {
		// handle error
	}
	err = wc.Close()
	if err != nil {
		// handle error
	}
	c.Quit()
}
