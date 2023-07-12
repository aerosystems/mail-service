package EmailService

type MailService struct {
	Provider EmailProvider
}

type EmailProvider interface {
	SendEmail(to, subject, body string) error
}

func (m *MailService) SetProvider(provider EmailProvider) {
	m.Provider = provider
}

func (m *MailService) SendEmail(to, subject, body string) error {
	return m.Provider.SendEmail(to, subject, body)
}
