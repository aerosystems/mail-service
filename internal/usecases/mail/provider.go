package mail

import "errors"

type Provider struct {
	slug string
}

func (p Provider) String() string {
	return p.slug
}

var (
	UnknownProvider = Provider{"unknown"}
	Mailhog         = Provider{"mailhog"}
	Brevo           = Provider{"brevo"}
)

func FromString(provider string) (Provider, error) {
	switch provider {
	case "mailhog":
		return Mailhog, nil
	case "brevo":
		return Brevo, nil
	}
	return UnknownProvider, errors.New("unknown email provider")
}
