package main

import (
	"context"
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpHostname           string `envconfig:"HTTP_HOSTNAME" default:"0.0.0.0"`
	HttpPort               string `envconfig:"HTTP_PORT" default:"8080"`
	LogLevel               string `envconfig:"LOG_LEVEL" default:"info"`
	DigestInterval         string `envconfig:"DIGEST_INTERVAL" default:"weekly"`
	SmtpHostname           string `envconfig:"SMTP_HOSTNAME" default:"localhost"`
	SmtpPort               string `envconfig:"SMTP_PORT" default:"25"`
	SmtpUsername           string `envconfig:"SMTP_USERNAME" default:"mailweave"`
	SmtpPassword           string `envconfig:"SMTP_PASSWORD" default:"mailweave"`
	SmtpFrom               string `envconfig:"SMTP_FROM" default:"mailweave@localhost"`
	SmtpStartTLS           bool   `envconfig:"SMTP_STARTTLS" default:"true"`
	SmtpTLS                bool   `envconfig:"SMTP_TLS" default:"false"`
	SmtpInsecureSkipVerify bool   `envconfig:"SMTP_INSECURE_SKIP_VERIFY" default:"false"`
	DatabaseType           string `envconfig:"DATABASE_TYPE" default:"sqlite"`
	DatabasePath           string `envconfig:"DATABASE_PATH" default:"mailweave.db"`
	DatabaseHostname       string `envconfig:"DATABASE_HOSTNAME" default:"localhost"`
	DatabasePort           string `envconfig:"DATABASE_PORT" default:"5432"`
	DatabaseUsername       string `envconfig:"DATABASE_USERNAME" default:"mailweave"`
	DatabasePassword       string `envconfig:"DATABASE_PASSWORD" default:"mailweave"`
	DatabaseName           string `envconfig:"DATABASE_NAME" default:"mailweave"`
	MailboxType            string `envconfig:"MAILBOX_TYPE" default:"pop3"`
	POP3Hostname           string `envconfig:"POP3_HOSTNAME" default:"localhost"`
	POP3Port               string `envconfig:"POP3_PORT" default:"110"`
	POP3Username           string `envconfig:"POP3_USERNAME" default:"mailweave"`
	POP3Password           string `envconfig:"POP3_PASSWORD" default:"mailweave"`
	POP3SSL                bool   `envconfig:"POP3_SSL" default:"false"`
	POP3InsecureSkipVerify bool   `envconfig:"POP3_INSECURE_SKIP_VERIFY" default:"false"`
	IMAPHostname           string `envconfig:"IMAP_HOSTNAME" default:"localhost"`
	IMAPPort               string `envconfig:"IMAP_PORT" default:"143"`
	IMAPUsername           string `envconfig:"IMAP_USERNAME" default:"mailweave"`
	IMAPPassword           string `envconfig:"IMAP_PASSWORD" default:"mailweave"`
	IMAPSSL                bool   `envconfig:"IMAP_SSL" default:"false"`
	IMAPInsecureSkipVerify bool   `envconfig:"IMAP_INSECURE_SKIP_VERIFY" default:"false"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		slog.ErrorContext(context.Background())
	}
}
