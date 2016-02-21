# Golang Schulung

## Angeda

### Workshop 1
- Hello World
- Projektstruktur
- Sprachgrundlagen
  - Packages
  - Datentypen & Variablen
    - Build-in Datentypen
    - Array & Slices
    - Maps
    - Structs
    - Pointer
    - Pointer und Structs
    - Pointer und Slices/Maps
    - Funktionen
  - Kontrollstrukturen
    - Kontrollstrukturen - if {}
    - Kontrollstrukturen - for {}
    - Kontrollstrukturen - switch {}
- Das Package `os`
  - Schreiben in eine Datei
- Übungen
  - Übung 1: IDE
  - Übung 2: Key-Value Store

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
- Alternative Router
- Web Frameworks

### Workshop 7
* Größere Projekte organisieren
* Interssante Packages und Libraries
- Rundgang durch die Standard Library

### Workshop 8
* TBD

