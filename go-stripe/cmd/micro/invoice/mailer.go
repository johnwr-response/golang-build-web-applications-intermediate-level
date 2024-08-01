package main

import (
	"bytes"
	"embed"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"time"
)

//go:embed email-templates
var emailTemplatesFS embed.FS

func (app *application) SendMail(from, to, subject, tmpl string, attachments []string, data interface{}) error {
	// build html content
	templateToRender := fmt.Sprintf("email-templates/%s.html.gohtml", tmpl)
	t, err := template.New("email-html").ParseFS(emailTemplatesFS, templateToRender)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", data); err != nil {
		app.errorLog.Println(err)
		return err
	}
	formattedMessage := tpl.String()

	// build plain text content
	templateToRender = fmt.Sprintf("email-templates/%s.plain.gohtml", tmpl)
	t, err = template.New("email-plain").ParseFS(emailTemplatesFS, templateToRender)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	if err := t.ExecuteTemplate(&tpl, "body", data); err != nil {
		app.errorLog.Println(err)
		return err
	}
	plainMessage := tpl.String()

	// build email message
	email := mail.NewMSG()
	email.SetFrom(from).
		AddTo(to).
		SetSubject(subject).
		SetBody(mail.TextHTML, formattedMessage).
		AddAlternative(mail.TextPlain, plainMessage)
	if len(attachments) > 0 {
		for _, attachment := range attachments {
			email.AddAttachment(attachment)
		}
	}

	app.infoLog.Println(formattedMessage, plainMessage)

	// send the mail
	server := mail.NewSMTPClient()
	server.Host = app.cfg.SmtpHost
	server.Port = app.cfg.SmtpPort
	server.Username = app.cfg.SmtpUsername
	server.Password = app.cfg.SmtpPassword
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}
	err = email.Send(smtpClient)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	app.infoLog.Printf("sent email %s from %s to %s\n", subject, from, to)

	return nil
}
