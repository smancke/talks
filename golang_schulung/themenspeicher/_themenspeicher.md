
Wie gehts dann weiter:
- OO-Light Funktionen auf Typen
- Einfaches testen
- Was ist HTTP/curl
- Http Client verwenden
- Einfacher HTTP-Server
- Einfacher Rest-Server
- Beispiel Key-Value Store
- Einfaches Templating

Interessante Konzepte
- Regex + String-Verarbeitung
- UTF8
- JSON
- REST
- Performance Testing und Komplexitäten einschätzen
- Testing
- Pair Programming
- git
- Dokumentieren

Beispiele:
- Unix Tools, cat
- Lesen schreiben in Datei
- Komandozeilen Parameter
- Environment Variablen


-----------------------------------


### Workshop 2
- Weitere Grundlagen
  - defer
  - Panic, Recover
- Type switches
- Testen von Go code
  - Tests
  - stretchr/testify
  - Test mehrerer Packages
  - Coverage anschauen
- Objekte
  - Methoden
  - Objekte
  - Konstruktoren
  - Embedding
  - Überschreiben
- Interfaces
- Übungen
  - Übung 3: Key-Value Objekt Orientiert
  - Übung 4: Testen

### Workshop 3
- Workshop 3
- Goroutinen
  - Goroutinen Syntax
  - Maximale Threads begrenzen
  - Sheduling
- Channel
  - Basics
  - Unbuffered Channel
  - Buffered Channel
  - Schließen eines Channel
  - Mit range über Channels iterieren
  - Select
  - Channel Tricks: Exit on Signal
  - Channel Tricks: Close als Wait-Broadcast
  - Channel Tricks: Channel mit Callback
- Das `sync` Package
  - Mutexes
  - Wait Group
- Benchmarks
- Übungen
  - Übung 5a: Concurrent Key-Value Store Access
  - Übung 5b: Concurrency Test
  - Übung 6: Benchmarking des Key-Value Stores

### Workshop 4 - Tooling
- Dependency Management
  - GO15VENDOREXPERIMENT
  - glide
  - Alternativen
- Debugging mit delve
  - Installation unter Linux
  - Installation OSX, Windows
- Profiling
- GODEBUG
- runtime/trace
- Race detection

### Workshop 5 - Web Development Teil 1
- Http Basics
  - Http Handler 
  - Http Handler interface
  - http.ServeMux
  - http.Server
  - `DefaultServeMux`
  - `http.Request`
  - `http.ResponseWriter`
- Templating
  - Ausführung
- Nette Features im Http Package
  - `http.FileServer`
  - `http.Client`
  - Package `http/httptest`
- Übung: File Storage Server
- Übung: REST Interface für den KV-Store

### Workshop 6 - Web Development Teil 2
- HTTP2 mit Go1.6
- Middleware
  - Chaining von Handlern
  - Handler aus dem Gorilla Toolkit
- Context
- gorilla/context
- Web Frameworks & alternative Router
  - `julienschmidt/httprouter`
  - Gorilla Toolkit
  - Gocraft/web
- Web Sockets
- Databases mit `jinzhu/gorm`
- Übung: Testen eines Handlers

### Workshop 7


# Workshop 7 - Verschiedenes
- Golang im Docker Container
- Libraries und Packages
  - Package `time`
  - Argumente und Umgebungsvariablen
  - Logging
  - Bolt
  - UI Libraries
- Coole Anwendungen, in Golang
- Rundgang durch Guble
- Gute Videos

