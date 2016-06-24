package mobileconfig

import (
	"io"
	"text/template"

	"github.com/twinj/uuid"
)

var t = template.Must(template.New("mobileconfig").Parse(mailTemplate))

type Config struct {
	EmailAddress string
	DisplayName string
	Identifier string
	Organization string
	AccountDescription string

	Imap *Imap
	Smtp *Smtp

	Description string
	ContentUuid string
	Uuid string
}

type Imap struct {
	Hostname string
	Port int
	Tls bool

	Username string
	Password string
}

type Smtp struct {
	Hostname string
	Port int
	Tls bool

	// Leave Username blank to do not use SMTP authentication.
	Username string
	// Leave Password blank to use IMAP credentials.
	Password string
}

func (c *Config) WriteTo(w io.Writer) error {
	if c.ContentUuid == "" {
		uuid := uuid.NewV4()
		c.ContentUuid = uuid.String()
	}
	if c.Uuid == "" {
		uuid := uuid.NewV4()
		c.Uuid = uuid.String()
	}

	return t.Execute(w, c)
}
