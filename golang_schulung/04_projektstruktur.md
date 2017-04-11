## 04 Projektstruktur
### Anlegen eines Workspace
```shell
mkdir hello_world
cd hello_world
export GOPATH=`pwd`
mkdir -p src/hello
```

### Übersetzen und Ausführen
```shell
go install hello
bin/hello
```

## Weitere Befehle
```shell
go build <package>   # Baut nur, ohne zu installieren
go test <packages>   # Fürt Tests in den *_test.go aus
go get <packages>    # Holt und installiert alle Abhängigkeiten
go get -t <package>  # Holt und installiert alle Abhängigkeiten inkl. der für die Tests
```

## Sprachgrundlagen
### Packages
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

