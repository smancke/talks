# Golang Schulung

## Grundlagen

### [01 Basics](01_basics.md)
- Hello World
- Datentypen & Variablen
- Strings
- Call by Value
- Text Ausgaben
- Funktionen
- Kontrollstrukturen
- Übung: Multiplikationstabelle

### [02 Dateien](02_dateien.md)
- Das Package `OS`
- Schreiben in eine Datei
- Exec Beispiel
- Übung: cat Befehl
- Übung: wc Befehl
- Übung: tac Befehl

### [03 Structs und Pointer](03_structs_und_pointer.md)
- Structs
- Pointer
- Pointer und Structs
- Pointer und Slices/Maps
- Übung: Key-Value Store

### [04 Projektstruktur](04_projektstruktur.md)
- Gopath
- Go Toolchain
- Packages

### [05 Testing](05_testing.md)
- Tests
- stretchr/testify
- Test mehrerer Packages
- Coverage anschauen
- Data Driven Tests
- Benchmarks
- Übung: Taschenrechner Programm
- Übung: Benchmarking des Key-Value Stores

### [06 Weitere Grundlagen](06_weitere_grundlagen.md)
- defer
- Panic, Recover
- Type switches

### [07 Objekte](07_objekte.md)
- Methoden
- Objekte
- Konstruktoren
- Embedding
- Überschreiben
- Interfaces
- Mocking
- Übung: Key-Value Objekt Orientiert
- Übung: Testen mit Mocks

## Web Development

### [08 Web Development Teil1](08_web_development_teil1.md)
- Http Basics
- Templating
- Nette Features im Http Package
- `http.Client`
- Package `http/httptest`

### [09 Web Development Teil2](09_web_development_teil2.md)
- HTTP2 mit Go
- Graceful Shutdown
- Middleware
- Context
- Web Frameworks & alternative Router
- Web Sockets
- Databases mit `jinzhu/gorm`

## Erweiterte Themen

### [10 Nebenlaufigkeit](10_nebenlaufigkeit.md)
- Goroutinen
- Channel
- Buffered Channel
- Das `sync` Package

### [11 Native Development](11_native_development.md)
- CGO
- Laden von Shared Libraries
- Cross Compiling

### [12 Golang im Docker Container](12_go_docker.md)

### [13 Libs](13_libs.md)
- awesome-go
- Package `time`
- Argumente und Umgebungsvariablen
- Logging
- Bolt

### [Debugging und Profiling](14_debugging_profiling.md)
- Race detection
- Debugging mit delve
- Profiling
- GODEBUG

### [Dependency Management](15_dependency_management.md)
- `vendor` Verzeichnis
- glide
- Alternativen

# Gute Videos
- [Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM), Rob Pike
- [Concurrency Is Not Parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg), Rob Pike
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs), Rob Pike
- [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw), Sameer Ajmani
- [Dont Just Check Errors Handle Them Gracefully](https://www.youtube.com/watch?v=lsBF58Q-DnY), Dave Cheney
