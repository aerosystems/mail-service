package models

type Message struct {
	To          string
	ToName      string
	From        string
	FromName    string
	Cc          string
	Subject     string
	Body        string
	Attachments []string
	DataMap     map[string]any
}
