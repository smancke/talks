# 04 Projektstruktur

## Gopath
Der GOPATH ist der Suchpfad für Go packages.

Seit go 1.8 ist der default bei nicht gesetztem GOPATH: `~/go`

### Anlegen eines Workspace
```shell
mkdir hello_world
cd hello_world
export GOPATH=`pwd`
mkdir -p src/hello
```

### Praktischer Alias
``
alias gopath="export GOPATH=\`pwd\`; export PATH=\`pwd\`/bin:$PATH"
``

### Übersetzen und Ausführen
```shell
go install hello
bin/hello
```

## Go Toolchain
```shell
go build <package>    # Baut nur, ohne zu installieren
go test <packages>    # Führt Tests in den *_test.go aus
go get <packages>     # Holt und installiert alle Abhängigkeiten
go get -t <package>   # Holt und installiert alle Abhängigkeiten inkl. der für die Tests
go vet <package>      # Statische Fehleranalyse
go fmt <package>      # Einheitliche Code Formatierung
go generate <package> # Source Code Generierung

```

### go get
Go's dependency Konzept

* Source Code Dependencies
* Kein library-Konzept (ab 1.8 aber .so aber möglich)
* Automatischer Download und Build über `go get package/name`

Beispiel:

    export GOPATH=`pwd`
    go get github.com/smancke/servelocal

Imports sind Referenzen auf Code Repositories:

    import "github.com/gorilla/handlers"

## Packages
* Eine Datei besitzt eine Package Deklaration
* Pro Verzeichnis gibt es nur ein package
* Das Package `main` wird als Entrypoint verwendet
* Import von Paketen
```go
import(
    "fmt"
    "math"
)
```
* Packages können voll qualifiziert importiert werden, z.B.  `import github.com/smancke/guble`
* Groß geschriebene Bezeichner in einem Package werden exportiert.

### init()

Die `init()`-Funktion eines Packages wird direkt nach dem Import des Packages (vor `main.main()`) aufgerufen,
auch wenn gar kein Code aus dem Package explizit verwendet wird.
```
package foo

func init() {
    // do some initial stuff here
}
```

```
package main

import (
    _ foo
)
```
