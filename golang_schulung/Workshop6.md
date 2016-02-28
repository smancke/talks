# Workshop 6 - Web Development Teil 2

## HTTP2 mit Go1.6
Go 1.6 hat Unterstützung für HTTP2 über die bestehende HTTP API.
Um HTTP2 setzt TLS voraus.

```go
panic(http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil))
```

## Middleware

### Chaining von Handlern
Handler können einfach in einander geschachtelt werden.
Generische Handler fungieren dabei als sog. Middleware.

```go
loggingMiddleware(delegate http.Handler) http.Handler
```
### Handler aus dem Gorilla Toolkit
Handler sind generisch und können aus unterschielichen Frameworks kombiniert werden.

http://www.gorillatoolkit.org/pkg/handlers

## Context
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

## `julienschmidt/httprouter`

Sehr schneller HTTP Router mit Parameter Parsing.
Basiert intern auf einem Suchbaum.

```go
import "github.com/julienschmidt/httprouter"


mux := httprouter.New()
mux.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", params.ByName("name"))
})
```

## Web Frameworks
### Gorilla
### Gocraft/web
## Web Sockets
## JSON-RPC / Protobuf
