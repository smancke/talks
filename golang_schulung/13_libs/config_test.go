package main

import (
	"flag"
	"github.com/smancke/mailigo/mail"
	. "github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestConfig_ReadConfigDefaults(t *testing.T) {
	originalArgs := os.Args
	os.Args = []string{"mailigo"}
	defer func() { os.Args = originalArgs }()

	defaultConfig := DefaultConfig()
	gotConfig := ReadConfig()
	Equal(t, defaultConfig, gotConfig)
}

func TestConfig_ReadConfig(t *testing.T) {
	input := []string{
		"--host=host",
		"--port=42",
		"--log-level=loglevel",
		"--text-logging=true",
		"--jwt-secret=jwtsecret",
		"--db-driver=dbdriver",
		"--db-datasource=dbdatasource",
		"--smtp-host=smtphost",
		"--smtp-port=43",
		"--smtp-username=smtpusername",
		"--smtp-password=smtppassword",
		"--smtp-ssl=true",
		"--grace-period=4s",
	}

	expected := &Config{
		Host:         "host",
		Port:         42,
		LogLevel:     "loglevel",
		TextLogging:  true,
		JwtSecret:    "jwtsecret",
		DBDriver:     "dbdriver",
		DBDataSource: "dbdatasource",
		SMTPConfig: mail.SMTPConfig{
			Host:     "smtphost",
			Port:     43,
			Username: "smtpusername",
			Password: "smtppassword",
			SSL:      true,
		},
		GracePeriod: 4 * time.Second,
	}

	cfg, err := readConfig(flag.NewFlagSet("", flag.ContinueOnError), input)
	NoError(t, err)
	Equal(t, expected, cfg)
}

func TestConfig_ReadConfigFromEnv(t *testing.T) {
	NoError(t, os.Setenv("MAILIGO_HOST", "host"))
	NoError(t, os.Setenv("MAILIGO_PORT", "42"))
	NoError(t, os.Setenv("MAILIGO_LOG_LEVEL", "loglevel"))
	NoError(t, os.Setenv("MAILIGO_TEXT_LOGGING", "true"))
	NoError(t, os.Setenv("MAILIGO_JWT_SECRET", "jwtsecret"))
	NoError(t, os.Setenv("MAILIGO_DB_DRIVER", "dbdriver"))
	NoError(t, os.Setenv("MAILIGO_DB_DATASOURCE", "dbdatasource"))
	NoError(t, os.Setenv("MAILIGO_SMTP_HOST", "smtphost"))
	NoError(t, os.Setenv("MAILIGO_SMTP_PORT", "43"))
	NoError(t, os.Setenv("MAILIGO_SMTP_USERNAME", "smtpusername"))
	NoError(t, os.Setenv("MAILIGO_SMTP_PASSWORD", "smtppassword"))
	NoError(t, os.Setenv("MAILIGO_SMTP_SSL", "true"))
	NoError(t, os.Setenv("MAILIGO_GRACE_PERIOD", "4s"))

	defer func() {
		os.Unsetenv("MAILIGO_HOST")
		os.Unsetenv("MAILIGO_PORT")
		os.Unsetenv("MAILIGO_LOG_LEVEL")
		os.Unsetenv("MAILIGO_TEXT_LOGGING")
		os.Unsetenv("MAILIGO_JWT_SECRET")
		os.Unsetenv("MAILIGO_DB_DRIVER")
		os.Unsetenv("MAILIGO_DB_DATASOURCE")
		os.Unsetenv("MAILIGO_SMTP_HOST")
		os.Unsetenv("MAILIGO_SMTP_PORT")
		os.Unsetenv("MAILIGO_SMTP_USERNAME")
		os.Unsetenv("MAILIGO_SMTP_PASSWORD")
		os.Unsetenv("MAILIGO_SMTP_SSL")
		os.Unsetenv("MAILIGO_GRACE_PERIOD")
	}()

	expected := &Config{
		Host:         "host",
		Port:         42,
		LogLevel:     "loglevel",
		TextLogging:  true,
		JwtSecret:    "jwtsecret",
		DBDriver:     "dbdriver",
		DBDataSource: "dbdatasource",
		SMTPConfig: mail.SMTPConfig{
			Host:     "smtphost",
			Port:     43,
			Username: "smtpusername",
			Password: "smtppassword",
			SSL:      true,
		},
		GracePeriod: 4 * time.Second,
	}

	cfg, err := readConfig(flag.NewFlagSet("", flag.ContinueOnError), []string{})
	NoError(t, err)
	Equal(t, expected, cfg)
}
