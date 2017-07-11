# Libraries und Packages

## awesome-go
https://github.com/avelino/awesome-go

## Package `time`

Type `Duration`: https://golang.org/pkg/time/#Duration
```
const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)
func ParseDuration(s string) (Duration, error)
```

Timer Funktionen über Channels und Closures:
```
    func After(d Duration) <-chan Time
    func Sleep(d Duration)
    func Tick(d Duration) <-chan Time

    type Timer
        func AfterFunc(d Duration, f func()) *Timer
        func NewTimer(d Duration) *Timer
        func (t *Timer) Reset(d Duration) bool
        func (t *Timer) Stop() bool
```

Formatieren und Parsen:
Angabe des Layouts über ein Referenzdatum: `Mon Jan 2 15:04:05 -0700 MST 2006`
```
const (
        ANSIC       = "Mon Jan _2 15:04:05 2006"
        UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
        RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
        RFC822      = "02 Jan 06 15:04 MST"
        RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
        RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
        RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
        RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
        RFC3339     = "2006-01-02T15:04:05Z07:00"
        RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
        Kitchen     = "3:04PM"
        // Handy time stamps.
        Stamp      = "Jan _2 15:04:05"
        StampMilli = "Jan _2 15:04:05.000"
        StampMicro = "Jan _2 15:04:05.000000"
        StampNano  = "Jan _2 15:04:05.000000000"
)
        
    func Parse(layout, value string) (Time, error)
    func (t Time) Format(layout string) string
```

## Argumente und Umgebungsvariablen

### Package `flag`

Komandozeilenparameter können über das `flag` Package eingelesen werden.

Beispiel einer Anwendungskonfiguration:
```

func DefaultConfig() *Config {
        return &Config{
                Host:         "localhost",
                Port:         8080,
        }
}                                                                                                                                                                          
                                                                                                                                                                           
const envPrefix = "MYAPP_"                                                                                                                                               

type Config struct {
        Host         string
        Port         int
}

// ConfigureFlagSet adds all flags to the supplied flag set
func (c *Config) ConfigureFlagSet(f *flag.FlagSet) {
        f.StringVar(&c.Host, "host", c.Host, "The host to listen on")
        f.IntVar(&c.Port, "port", c.Port, "The port to listen on")
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
```

### Anotations basierte Parameter

```
import (
	"github.com/alexflint/go-arg"
	"github.com/caarlos0/env"
}

type Args struct {
	Listen      string `arg:"-l,help: [Host:]Port the address to listen on (:8080)" env:"GUBLE_LISTEN"`
	LogInfo     bool   `arg:"--log-info,help: Log on INFO level (false)" env:"GUBLE_LOG_INFO"`
	LogDebug    bool   `arg:"--log-debug,help: Log on DEBUG level (false)" env:"GUBLE_LOG_DEBUG"`
	StoragePath string `arg:"--storage-path,help: The path for storing messages and key value data if 'file' is enabled (/var/lib/guble)" env:"GUBLE_STORAGE_PATH"`
	KVBackend   string `arg:"--kv-backend,help: The storage backend for the key value store to use: file|memory (file)" env:"GUBLE_KV_BACKEND"`
	MSBackend   string `arg:"--ms-backend,help: The message storage backend : file|memory (file)" env:"GUBLE_MS_BACKEND"`
	GcmEnable   bool   `arg:"--gcm-enable: Enable the Google Cloud Messaging Connector (false)" env:"GUBLE_GCM_ENABLE"`
	GcmApiKey   string `arg:"--gcm-api-key: The Google API Key for Google Cloud Messaging" env:"GUBLE_GCM_API_KEY"`
}

func main() {
    ..
    args := loadArgs()
    ..
}
    
func loadArgs() Args {
	args := Args{
		Listen:      ":8080",
		KVBackend:   "file",
		MSBackend:   "file",
		StoragePath: "/var/lib/guble",
	}

	env.Parse(&args)
	arg.MustParse(&args)
	return args
}
```

## Logging
### Package `log`
Einfaches, aber limitiertes Logging Framework in der Standard Library.

```
    func SetOutput(w io.Writer)
    func SetPrefix(prefix string)

    func Fatal(v ...interface{})
    func Fatalf(format string, v ...interface{})
    func Fatalln(v ...interface{})

    func Panic(v ...interface{})
    func Panicf(format string, v ...interface{})
    func Panicln(v ...interface{})

    func Print(v ...interface{})
    func Printf(format string, v ...interface{})
    func Println(v ...interface{})
```
    
### Sirupsen/logrus
https://github.com/Sirupsen/logrus

- Strukturiertes Logging
- Loglevel
- Context Logger
- Highlighting in der Konsole
- Json Logger
- Replacement für `log` Package
- Viele Log Backends

## Bolt
Bolt ist eine library für einen Key-Value Store: https://github.com/boltdb/bolt

- Buckets und Sub Buckets
- Transactions
- Prefix & Range Scans
