package provider

type Postfix struct{}

func (m *Postfix) SendEmail(to, subject, body string) error {
	// TODO: implement postfix provider
	return nil
}
