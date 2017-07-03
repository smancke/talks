# Workshop 6 - Web Development Teil 2

## HTTP2 mit Go
Seit Version 1.6 hat Go Unterstützung für HTTP2 über die bestehende HTTP API.
HTTP2 setzt TLS voraus.

```go
panic(http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil))
```

## Graceful Shutdown

```go
httpSrv := &http.Server{Addr: ":8080", Handler: handlerChain}

go func() {
	if err := httpSrv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("not accepting new connections")
		} else {
			fmt.Printf("error %v", err)
		   	os.Exit(1)
		}
	}
}()

stop := make(chan os.Signal)
signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

sig := <-stop
fmt.Printf("got %v, shutdown gracefully, now\n", sig)

ctx, ctxCancel := context.WithTimeout(context.Background(), gracePeriod)

httpSrv.Shutdown(ctx)
fmt.Println("down")
ctxCancel() // not needed, but good style
```
    
## Middleware

### Chaining von Handlern
Handler können einfach in einander geschachtelt werden.
Generische Handler fungieren dabei als sog. Middleware.

```go
loggingMiddleware(delegate http.Handler) http.Handler
```

Beispiel:
```go
	chain := LoggingMiddleware(
		AccessMiddleware(
			helloWorld
		)
	)
	panic(http.ListenAndServe(":8080", chain))
```


### Chaining mit `alice`
https://github.com/justinas/alice

```go
	chain := alice.New(
		LoggingMiddleware,
		AccessMiddleware,
	).Then(helloWorld)
```

### Handler aus dem Gorilla Toolkit
Handler sind generisch und können aus unterschielichen Frameworks kombiniert werden.

http://www.gorillatoolkit.org/pkg/handlers

## Context

### Go >= 1.7

Seit Version 1.7 ist es sehr einfach geworden requestspezifische Informationen zu kommunizieren.

#### Context
Mit 1.7 wurde der Typ Context in die Standardbibliothek aufgenommen und kann so von Structs und Methoden dort genutzt werden. 
```go 
type Context interface {
        Deadline() (deadline time.Time, ok bool)
        Done() <-chan struct{}
        Err() error
        Value(key interface{}) interface{}
}
```
Die Methode `Value(key interface{}) interface{}` kann genutzt werden um einen Wert aus einem Context zu holen.

Die Funktion `func WithValue(parent Context, key, val interface{}) Context` kann genutzt werden um einen Context mit einem neuen Wert anzureichern.

#### Request-Context

Auch der  `http.Request`-Typ hat 2 neue Methoden erhalten:
```go
func (r *Request) Context() context.Context
```
Diese Methode liefert den Context welcher zum Request gehört. 

```go
func (r *Request) WithContext(ctx context.Context) *Request
```
Diese Methode setzt einen neuen Context und liefert den Request mit neuem Context.

Hierdurch können sehr einfach Informationen im Requestscope durch Middlewares / Handlerchains hindurch transportiert und geteilt werden.

### Go <= 1.6

Request spezifische Daten lassen sich nur schwer an andere Handler mitgeben.

Lösung: Globale Map `*http.Request->Daten`.

```go
var context struct {
	data  map[*http.Request]interface{}
	mutex sync.Mutex
}

func GetUsername(r *http.Request) string {
    ..
}
func SetUsername(r *http.Request, username string) {
    ..
}
```

### gorilla/context
Fertige Lösung im Package `gorilla/context` des Gorilla Toolkit.

http://www.gorillatoolkit.org/pkg/context

## Web Frameworks & alternative Router

### `julienschmidt/httprouter`

Sehr schneller HTTP Router mit Parameter Parsing.
Basiert intern auf einem Suchbaum.

```go
import "github.com/julienschmidt/httprouter"


mux := httprouter.New()
mux.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", params.ByName("name"))
})
```

### Gorilla Toolkit
Das Gorilla Toolkit liefert eine Reihe von Packages, die unabhängig von einander verwendet werden können.

http://www.gorillatoolkit.org


- `gorilla/context` stores global request variables.
- `gorilla/mux` is a powerful URL router and dispatcher.
- `gorilla/reverse` produces reversible regular expressions for regexp-based muxes.
- `gorilla/rpc` implements RPC over HTTP with codec for JSON-RPC.
- `gorilla/schema` converts form values to a struct.
- `gorilla/securecookie` encodes and decodes authenticated and optionally encrypted cookie values.
- `gorilla/sessions` saves cookie and filesystem sessions and allows custom session backends.
- `gorilla/websocket` implements the WebSocket protocol defined in RFC 6455.


### Gocraft/web
Einfaches Web Framework, das Contexte weiter reicht und Objekt orientierte Handler ermöglicht.

```go
type Context struct {
	Username string
}


func (c *Context) Hello(w web.ResponseWriter, r *web.Request) {

	fmt.Fprintf(w, "Hello %v\n", c.Username)
}

mux := web.New(Context{}).
    Get("/", (*Context).Hello)
```
    
## Web Sockets
Gut Web Socket Implementierung in `gorilla/websocket`.
Die Go standard Implementierung kann keine Chunked Messages.

Web Sockets laufen durch ein *Upgrade* im normalen HttpHandling.

Server:
```go
import "github.com/gorilla/websocket"

var webSocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (factory *WSHandlerFactory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c, err := webSocketUpgrader.Upgrade(w, r, nil)
	_, bytes, err := c.ReadMessage()
    ..
	c.WriteMessage(websocket.BinaryMessage, bytes)
    ..
}
```

Client
```go
conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080", nil)
if err != nil {
	panic(err)
}
conn.WriteMessage(websocket.TextMessage, []byte(strings.Join(os.Args[1:], " ")))
_, body, err := conn.ReadMessage()
fmt.Println(string(body))
```

## Databases mit `jinzhu/gorm`
Gorm ist ein minimaler OR-Mapper.

Imports:
```go
import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)
```

Typdefinition:
```go
type User struct {
	ID        uint       `gorm:"primary_key"`
	UserName  string     `sql:"type:varchar(50);unique_index"json:"userName"`
	NickName  string     `sql:"type:varchar(50)"json:"nickName"`
	Link      string     `sql:"type:varchar(500)"json:"link"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
```

Öffnen einer SQLite Datenbank:
```go
gormdb, err := gorm.Open("sqlite3", filename)
if err == nil {
	if err := gormdb.DB().Ping(); err != nil {
		log.Println("error pinging database: %v", err)
	} else {
		log.Println("can ping database")
	}

	//gormdb.LogMode(true)
	gormdb.DB().SetMaxIdleConns(2)
	gormdb.DB().SetMaxOpenConns(5)
	gormdb.SingularTable(true)

	if err := gormdb.AutoMigrate(&User{}).Error; err != nil {
        log.Printf("error in schema migration: %v", err)
	    return err
    } else {
		log.Println("ensured db schema")
	}
} else {
	log.Println("error opening sqlite3 db %v: %v", filename, err)
    }
```

Queries:
```go
func CreateUser(user *User) error {
	return gormdb.Create(user).Error
}

func SaveUser(user *User) error {
	return gormdb.Save(user).Error
}

func UserByUsername(username string) (*User, error) {
	user := &User{}
	if err := gormdb.
		Where("user_name = ?", username).
		First(user).Error; err != nil {

		return user, err
	}
	return user, nil
}
```


## Übung: Testen eines Handlers
Schreibe einen Handler mit einem Framework Deiner Wahl und teste
diesen mit dem HttpRecorder.
