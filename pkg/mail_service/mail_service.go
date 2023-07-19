package EmailService

type MailService struct {
	Provider EmailProvider
}

type EmailProvider interface {
	SendEmail(msg Message) error
}

type Message struct {
	To          string
	ToName      string
	From        string
	FromName    string
	Subject     string
	Body        string
	Attachments []string
	DataMap     map[string]any
}

func (m *MailService) SetProvider(provider EmailProvider) {
	m.Provider = provider
}

func (m *MailService) SendEmail(msg Message) error {
	return m.Provider.SendEmail(msg)
}
