
Überblick
=================

* Relativ neue Programmiersprache
 * go 1.0 in 2012
 * aktuell: go 1.7
* BSD License
* Erfunden und maintained von google
* Statisches Typsystem
* Garbage Collector
* Statisch gelinkte binaries

Haupt Einsatzzwecke
===================

* Server side prgramming
* Datenberarbeitung
* HTTP Server
* Microservices

Projekte in Go
--------------
* Docker
* Drone
* etcd
* Kubernetes
* consul
* Prometheus
* InfluxDB

Warum ich go verwende? 
======================
Go hat mir den Spaß am Programmieren zurück gebracht!

* Go versucht nicht schön zu sein, sondern: einfach und gut.
* Go geht aus dem Weg und unterstützt einen sachen einfach umzusetzen.

Charakter
=================
TODO: Mindmap 

Syntax
=================
* Angelehnt an gewohnte C-Syntax, aber mit Vereinfachungen
* Sparsam: Keine Klammern, kein `;`
* Public Bezeichner werden groß geschrieben
* Typ Inferenz `:=` initialisiert und deklariert eine Variable
* Mächtiger `for` loop für alle Schleifentypen
* `if` mit Initialisierung
* Funktionen mit mehreren Rückgabewerten

Syntax Beispiel
=================

    package main

    import (
	    "regexp"
    )

    func main() {
	    terms := []string{"java", "is", "fun"}
	    for _, value := range terms {
	        rx := regexp.MustCompile("java")
	        value = rx.ReplaceAllString(value, "golang")
	        println(value)
	    }
    }

->
    
    $ go run examples/syntax_example.go 
    golang
    is
    fun


Tooling
=========
* Build + Test tools included
* Super schneller Compiler

Source Code Dependencies
=========================
* Quellcode Dependencies

Packages
=================

Einfaches gutes Package Konzept

Deklaration über:
```go
package mypackage
```

* Groß geschriebene Bezeichner werden exportiert, alle anderen sind nur im package sichtbar
* Die Aufteilung von Code auf Dateien in einem Package kann beliebig erfolgen


Source Code Dependencies:
```go
import  "github.com/gocraft/web"
```

Nutzung im Code über letzten Pfadteil:
```go
web.Router
```

Verzögerte Ausführung `defer`
=============================

defer statements werden am Ende der funktion ausgeführt (== finally{})

```go
file, err := os.Open(srcName)
if err != nil {
    return
}
defer src.Close()
```

Typsystem
=================
* Streng typisiert
* Structs
* Build in maps und slices
* Interfaces und Ducktyping

`structs`
=================

```go
type User struct {
	UserName  string     `json:"userName"`
	NickName  string     `json:"nickName"`
}

user := User{UserName: "Ben", NickName: "Utzer"}
```

Unterstützung für Delegation und Embedding.

Fehlerbehandlung
=================
Fehlerhandling läuft meist über Rückgabewert.
```go
if err := machEtwas(); err != nil {
    // handle error
}
```

Es gibt aber auch ein Equivalent zu Exceptions.
```go
func travel() {
	defer func() {
	    if r := recover(); r != nil {
            fmt.Println(r, "dont't panic!", )
        }
	}()
	panic("I lost my towel")
}
```

Objecte
=================

Funktionen können auf eigenen Datentypen definiert werden.

```go
func (user *User) CallUser(msg string) {
    fmt.Printf("Hallo %v: %v", user.NickName, msg)
}

user := User{UserName: "Ben", NickName: "Utzer"}
user.CallUser("hör gut zu!")
```

`go` routinen
=================
Leichtgewichtige co-routinen, die im Hintergrund laufen.
```go
go doSomething()
```

`channel`
=================

Channel sind Datenstrukturen zur sicheren Kommunikation bei paralleler Verarbeitung. 
```go
func waitForTermination(callback func()) {
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Got singal '%v' .. exit more or less greacefully now", <-sigc)
	callback()
	os.Exit(0)
}
```

Danke ...
================
... Slides auf github: [https://github.com/smancke/talks](https://github.com/smancke/talks/tree/master/2015_froscon_docker_in_production)
