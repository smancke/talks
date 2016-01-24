
Allgemeines
------------

* Statisch gelinkte binaries
* Besonderers gut für Server
* Sehr schneller Build
* Dynamische Speicherverwaltung
* Build + Tests tool included

Angenehme Syntax
-----------------
* Sehr gewohnt, aber vereinfacht
* Sparsam: Keine Klammern, kein `;`
* Public Bezeichner werden groß geschrieben
* `:=` initialisiert und deklariert eine Variable
* Mächtiger `for` loop für alle Schleifentypen
* `if` mit Initialisierung
* Mehrere Rückgabewerte

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	terms := []string{"bli", "bla", "blub"}
	for _, value := range terms {
		rx := regexp.MustCompile("bl")
		value = rx.ReplaceAllString(value, "di")
		fmt.Println(value)
	}
}
```

Packages
----------
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
------------------------------
defer statements werden am Ende der funktion ausgeführt (== finally{})

```go
file, err := os.Open(srcName)
if err != nil {
    return
}
defer src.Close()
```

Typsystem
-------------
* Streng typisiert
* Structs
* Build in maps und slices
* Interfaces und Ducktyping

`structs`
----------
```go
type User struct {
	UserName  string     `json:"userName"`
	NickName  string     `json:"nickName"`
}

user := User{UserName: "Ben", NickName: "Utzer"}
```

Unterstützung für Delegation und Embedding.

Fehlerbehandlung
-----------------
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
---------
Funktionen können auf eigenen Datentypen definiert werden.

```go
func (user *User) CallUser(msg string) {
    fmt.Printf("Hallo %v: %v", user.NickName, msg)
}

user := User{UserName: "Ben", NickName: "Utzer"}
user.CallUser("hör gut zu!")
```

`go` routinen
-------------
Leichtgewichtige co-routinen, die im Hintergrund laufen.
```go
go doSomething()
```

`channel`
---------
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
