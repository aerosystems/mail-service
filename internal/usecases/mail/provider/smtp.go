package provider

import (
	"bytes"
	"github.com/aerosystems/mail-service/internal/models"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"time"
)

type Smtp struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
}

func NewSmtp(domain, host string, port int, username, password, encryption string) *Smtp {
	return &Smtp{
		Domain:     domain,
		Host:       host,
		Port:       port,
		Username:   username,
		Password:   password,
		Encryption: encryption,
	}
}

func (s Smtp) SendEmail(msg models.Message) error {
	data := map[string]any{
		"message": msg.Body,
	}
	msg.DataMap = data

	formattedMessage, err := s.buildHtmlMessage(msg)
	if err != nil {
		return err
	}

	plainMessage, err := s.buildPlainTextMessage(msg)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = s.Host
	server.Port = s.Port
	server.Username = s.Username
	server.Password = s.Password
	server.Encryption = s.getEncryption(s.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		AddCc(msg.Cc).
		SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}

	if err := email.Send(smtpClient); err != nil {
		return err
	}
	return nil
}

func (s Smtp) buildHtmlMessage(msg models.Message) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = s.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (s Smtp) buildPlainTextMessage(msg models.Message) (string, error) {
	templateToRender := "./templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (s Smtp) inlineCSS(str string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(str, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (s Smtp) getEncryption(str string) mail.Encryption {
	switch str {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
