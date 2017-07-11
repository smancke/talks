package main

import (
	"flag"
	"github.com/smancke/mailigo/mail"
	"os"
	"strings"
	"time"
)

// DefaultConfig for mailigo
func DefaultConfig() *Config {
	return &Config{
		Host:         "localhost",
		Port:         8080,
		LogLevel:     "info",
		JwtSecret:    "secret",
		DBDriver:     "sqlite3",
		DBDataSource: "/var/lib/mailigo",
		GracePeriod:  5 * time.Second,
	}
}

const envPrefix = "MAILIGO_"

// Config for the loginsrv handler
type Config struct {
	Host         string
	Port         int
	LogLevel     string
	TextLogging  bool
	JwtSecret    string
	DBDriver     string
	DBDataSource string
	SMTPConfig   mail.SMTPConfig
	GracePeriod  time.Duration
}

// ConfigureFlagSet adds all flags to the supplied flag set
func (c *Config) ConfigureFlagSet(f *flag.FlagSet) {
	f.StringVar(&c.Host, "host", c.Host, "The host to listen on")
	f.IntVar(&c.Port, "port", c.Port, "The port to listen on")
	f.StringVar(&c.LogLevel, "log-level", c.LogLevel, "The log level")
	f.BoolVar(&c.TextLogging, "text-logging", c.TextLogging, "Log in text format instead of json")
	f.StringVar(&c.JwtSecret, "jwt-secret", c.JwtSecret, "The secret to sign the jwt token")
	f.StringVar(&c.DBDriver, "db-driver", c.DBDriver, "")
	f.StringVar(&c.DBDataSource, "db-datasource", c.DBDataSource, "")
	f.StringVar(&c.SMTPConfig.Host, "smtp-host", c.SMTPConfig.Host, "")
	f.IntVar(&c.SMTPConfig.Port, "smtp-port", c.SMTPConfig.Port, "")
	f.StringVar(&c.SMTPConfig.Username, "smtp-username", c.SMTPConfig.Username, "")
	f.StringVar(&c.SMTPConfig.Password, "smtp-password", c.SMTPConfig.Password, "")
	f.BoolVar(&c.SMTPConfig.SSL, "smtp-ssl", c.SMTPConfig.SSL, "")
	f.DurationVar(&c.GracePeriod, "grace-period", c.GracePeriod, "Graceful shutdown grace period")
}

// ReadConfig from the commandline args
func ReadConfig() *Config {
	c, err := readConfig(flag.NewFlagSet(os.Args[0], flag.ExitOnError), os.Args[1:])
	if err != nil {
		// should never happen, because of flag default policy ExitOnError
		panic(err)
	}
	return c
}

func readConfig(f *flag.FlagSet, args []string) (*Config, error) {
	config := DefaultConfig()
	config.ConfigureFlagSet(f)

	// prefer environment settings
	f.VisitAll(func(f *flag.Flag) {
		if val, isPresent := os.LookupEnv(envName(f.Name)); isPresent {
			f.Value.Set(val)
		}
	})

	err := f.Parse(args)
	if err != nil {
		return nil, err
	}

	return config, err
}

func envName(flagName string) string {
	return envPrefix + strings.Replace(strings.ToUpper(flagName), "-", "_", -1)
}
