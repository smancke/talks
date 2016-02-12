# Workshop 4

# Dependency Management

## GO15VENDOREXPERIMENT
In go 1.5:
```shell
export GO15VENDOREXPERIMENT=1
```
Ab go 1.6 der default!

*Effekt*:
- Packages in einem Unterverzeichnis `/vendor` werden bei go builds priorisert angezogen.
- Files in `/vendor` können mit ins git repository eingecheckt werden, müssen aber nicht.
- Die Abhängigkeiten können mit `godep` oder `glide` gemanaged werden.

*Vorsicht:*
Die selbe Library in unterschieldichen vendor-Verzeichnissen macht Probleme.
Siehe auch: https://github.com/mattfarina/golang-broken-vendor

## glide
Glide ist ein package manager mit semantic versioning.

```shell
go get github.com/masterminds/glide
```

*Benutzung:*
https://github.com/Masterminds/glide

```shell
# Erstellen einer glide.yaml Konfigurationsdatei
glide init

# Update aller Dependencies ensprechend der glide.yaml
# Schreibt die konkreten Versionen in die glide.lock
glide update

# Installieren der Versionen aus glide.lock in /vendor
glide install
```

## Alternativen
Übersicht von Package Management Tools:
https://github.com/golang/go/wiki/PackageManagementTools

### godep
Etabliertes tool, dass jetzt auch `GO15VENDOREXPERIMENT` unterstützt.

### vendetta
Organisiert das vendor Verzeichnis über git submodules


# Debugging mit delve

Delve ist ein debugger extra für go.

## Installation unter Linux

Entweder:
```shell
git clone git@github.com:derekparker/delve.git
GO15VENDOREXPERIMENT=1 make install
```

oder:
```shell
go get -v -u github.com/derekparker/delve/cmd/dlv
```

## Installation OSX, Windows

https://github.com/derekparker/delve/wiki/Building

https://github.com/derekparker/delve/wiki/Tips-&-Troubleshooting


# Profiling
Go hat einen einfachen Profiler eingebaut.

```go
import _ "net/http/pprof"
..
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

```shell
# heap usage snapshot
go tool pprof http://localhost:6060/debug/pprof/heap

# 10sec cpu profiling snapshot
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

# show top cpu consumer ordered by cumulative consumption
(pprof) top -cum

# create an svg graph
(pprof) web

# create an svg graph for the calltree of doSomething()
(pprof) web doSomething

# Display callers and callees of doSomething()
(pprof) peek doSomething
```

# Race detection
Golang hat ein tool zur Suche nach race conditions direkt eingebaut.
```shell
go test -race mypkg    // test the package
go run -race mysrc.go  // compile and run the program
go build -race mycmd   // build the command
go install -race mypkg // install the package
```

# Code Generieren go:generate

# Rundgang durch die Standard Library
    
