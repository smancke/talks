# Workshop 4 - Web Development

## Http Basics
### Http Handler 
Alles Http Handling läuft über das `Handler` interface

### Http Handler interface
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### http.ServeMux
`ServeMux` ist ein einfacher Multiplexer.
Er implementiert das `Handler` interface und erlaubt das
Registrieren von Handlern oder callbacks für bestimmte Pfade.
    
```go
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

### http.Server
`Server` ist der eigentliche Http Server, der gestartet wird.

```go
func (srv *Server) ListenAndServe() error
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
func (srv *Server) Serve(l net.Listener) error
```

### `DefaultServeMux`
Für einfache konfigurationen können default Funktionen verwendet werden.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
})
http.ListenAndServe(":8080", nil)
```

### `http.Request`
```go
type Request struct {
        Method string
        URL *url.URL
        Header Header

        Body io.ReadCloser
        ContentLength int64
        Host string
        Form url.Values
        PostForm url.Values
        MultipartForm *multipart.Form
        ...
}

func (r *Request) Cookies() []*Cookie
func (r *Request) FormValue(key string) string
func (r *Request) ParseForm() error
func (r *Request) Referer() string
func (r *Request) UserAgent() string
...
```

### `http.ResponseWriter`
```go
type ResponseWriter interface {
        Header() Header
        Write([]byte) (int, error)
        WriteHeader(int)
}
```

## Templating
Templating für text oder html in den Packages:
```go
import "text/template"
```

und
```go
import "html/template"
```

Gutes Intro in die Syntax:
https://gohugo.io/templates/go-templates/

### Ausführung
```go
t := template.New("template")
t.Funcs(functionMap)

template.Must(t.Parse(theTemplate))

err := t.Execute(os.Stdout, data)
```

## Nette Features im Http Package


### `http.FileServer`
Handler für Directory Listings.

```go
func FileServer(root FileSystem) Handler
```

Beispiel:
```go
http.Handle("/", http.FileServer(http.Dir(".")))
http.ListenAndServe(":8080", nil)
```

Das File System erfolgt über eine einfache Abstraktion,
so dass auch einfach ein Daten aus einem anderen Storage
zurück geliefert werden können. (Gutes Beispiel für Ducktyping)

### `http.Client`
Auch der Http-Client funktioniert denkbar einfach.

```go
type Client
       func (c *Client) Do(req *Request) (resp *Response, err error)
       func (c *Client) Get(url string) (resp *Response, err error)
       func (c *Client) Head(url string) (resp *Response, err error)
       func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
       func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```


```go
type Response struct {
        Status     string // e.g. "200 OK"
        StatusCode int    // e.g. 200
        Header Header
        Body io.ReadCloser
        ContentLength int64
        ...
}
```

### Package `http/httptest`

#### ResponseRecorder
Zum einfachen Testen von Handlern.

```go
handler := func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "something failed", http.StatusInternalServerError)
}

req, err := http.NewRequest("GET", "http://example.com/foo", nil)
if err != nil {
	log.Fatal(err)
}

w := httptest.NewRecorder()
handler(w, req)

fmt.Printf("%d - %s", w.Code, w.Body.String())
```

#### ResponseRecorder
Startet einen Server auf einem zufälligen freien Port.

```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, client")
}))
defer ts.Close()

res, err := http.Get(ts.URL)
```

## Übung: File Storage Server
Erstelle einen minimalen Server, der Dateien Speichern und Zurückliefern kann.

```shell
# Speichern einer Datei
# Antwort: Http 201
curl -X POST --data "some-text-data" 127.0.0.1:8080/path/to/file

# Abfragen einer Datei
# Antwort: Http 200 "some-text-data"
curl 127.0.0.1:8080/path/to/file

# Fehler bei nicht vorhandener Datei
# Antwort: Http 404
curl 127.0.0.1:8080/wrong/path
```

## Übung: REST Interface für den KV-Store
Versehe KV-Store aus den letzten Übungen mit einem REST interface.

```shell
# Speichern eines Wertes
# Antwort: Http 201
curl -X POST --data "some-text-data" 127.0.0.1:8080/key

# Abfragen eines Wertes
# Antwort: Http 200 "some-text-data"
curl 127.0.0.1:8080/key

# Fehler bei nicht vorhandenem KEy
# Antwort: Http 404
curl 127.0.0.1:8080/wrongKey
```

