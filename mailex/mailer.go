package mailex

import (
	"os"

	"github.com/chouandy/goex/awsex/service/cloudwatcheventsex"
	"github.com/chouandy/goex/osex"
)

// Mailer mailer instance
var Mailer *mailer

type mailer struct {
	SMTPSettings *SMTPSettings
	Options      *Options
}

// SMTPSettings smtp settings
type SMTPSettings struct {
	Address  string
	Port     int
	Username string
	Password string
}

// Options options struct
type Options struct {
	From string
}

// InitMailerEventMiddleware init mailer event middleware
func InitMailerEventMiddleware(ctx *cloudwatcheventsex.Context) error {
	if Mailer == nil {
		Mailer = &mailer{
			SMTPSettings: &SMTPSettings{
				Address:  os.Getenv("MAILER_SMTP_ADDRESS"),
				Port:     osex.GetenvParseInt("MAILER_SMTP_PORT"),
				Username: os.Getenv("MAILER_SMTP_USERNAME"),
				Password: os.Getenv("MAILER_SMTP_PASSWORD"),
			},
			Options: &Options{
				From: os.Getenv("MAILER_FROM"),
			},
		}
	}

	return nil
}
