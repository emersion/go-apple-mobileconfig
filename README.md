# go-apple-mobileconfig

[![GoDoc](https://godoc.org/github.com/ProtonMail/go-apple-mobileconfig?status.svg)](https://godoc.org/github.com/ProtonMail/go-apple-mobileconfig)

Create Apple mobileconfig configuration files.

Based on Node [`mobileconfig`](https://github.com/andris9/mobileconfig) module.

## Usage

```go
package main

import (
	"os"
	"github.com/ProtonMail/go-apple-mobileconfig"
)

func main() {
	c := &mobileconfig.Config{
		DisplayName: "Mail Account",
		EmailAddress: "root@nsa.gov",
		Identifier: "gov.nsa.mail",
		Organization: "NSA",
		Imap: &mobileconfig.Imap{
			Hostname: "mail.nsa.gov",
			Port: 993,
			Tls: true,
			Username: "root",
			Password: "snowden",
		},
		Smtp: &mobileconfig.Smtp{
			Hostname: "mail.nsa.gov",
			Port: 25,
			Tls: false,
			Username: "root",
		},
	}

	if err := c.WriteTo(os.Stdout); err != nil {
		panic(err)
	}
}
```

## License

MIT
