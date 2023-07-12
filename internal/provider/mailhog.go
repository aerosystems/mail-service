package provider

type Mailhog struct{}

func (m *Mailhog) SendEmail(to, subject, body string) error {
	// TODO: implement mailhog provider
	return nil
}
