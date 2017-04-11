
# 01 Basics

## Hello World
```go
// file: hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

Direkte Ausführung mit:
```go
go run hello.go
```

oder im [Golang Playground](https://play.golang.org/)

## Datentypen & Variablen
* Dynamische Speicherverwaltung
* Streng typisiert
* Alle Variablen haben einen sicheren Initialisierungswert
Deklaration und Zuweisung:
```go
var i int
i = 42
```

Inferenz für die Deklaration:
```go
i := 42
```

### Build-in Datentypen
* int, uint und float Typen in verschiedenen größen
* bool
* complex, rune
* arrays mit fester Länge `[4]int{42, 43, 44, 45}`
* Slices als Abstraktion über Arrays: `[]int{42, 43, 44, 45}`
* string (Entspricht einem byte Slice)
* Maps `map[string]int{"a": 42, "b": 43}`
* Structs

Typ Konvertierung:
```go
a := 42
b := uint64(a)
```

### Arrays
Arrays haben eine feste Größe. Sie sind recht unflexibel und werden selten direkt verwendet.
```go
farben := [5]string{"black", "red", "blue", "green", "white"}
fmt.Println(len(farben))
fmt.Println(farben[0])
```

### Slices
Slices bieten einen flexiblen Listen-Typ. Intern werden die Daten in einem Array gespeichert,
so dass ein Slice eine eine Referenz auf ein Array mit einem Start- und Index

```go
	farben := []string{"black", "red", "blue"}
	farben = append(farben, "green", "white")
	farben = farben[1 : len(farben)-1]

	fmt.Println(len(farben))
	fmt.Println(farben)
```

Explizite Schreibweise:
```go
slicename := make(type, len, cap)

// example:
farben := make([]string, 0, 5)
```

### Maps
```go
m := make(map[KeyType]ValueType)
```

Beispiel:
```go
person := map[string]string{
	"name":    "Mancke",
	"given": "Sebastian",
}
fmt.Println(person)
fmt.Println(len(person))

person["given"] = "Felix"
fmt.Println(person)
delete(person, "given")

_, exist := person["given"]
fmt.Println(exist) // false
```


## Text Ausgaben
```go
import (
	"fmt"
	"os"
)

func main() {
	var i = 42

	// direct output
	fmt.Print(i)
	fmt.Print("\n")

	// output with newline
	fmt.Println(i)

	// format values
	fmt.Printf("the answer is %v\n", i)

	// formating as string
	s := fmt.Sprintf("the answer is %v\n", i)
	fmt.Print(s)

	// write to writer, e.g. stderr
	fmt.Fprintf(os.Stderr, "the answer is %v\n", i)
}
```

## Strings
Strings in go sind slices von bytes:
```go
	b := []byte{72, 97, 108, 108, 111, 32, 87, 101, 108, 116}
	fmt.Println(string(b))
```

Zugriff aug Teilstrings
```go
	fmt.Println(s[0:5])
```

Das Package [strings](https://golang.org/pkg/strings/) enthält praktische Funktionen für den Umgang mit Zeichenketten.
```go
    fmt.Println(strings.HasPrefix(s, "Hallo"))
	fmt.Println(strings.ToLower(s))
```


Das Package [strconv](https://golang.org/pkg/strconv/) enthält Funktionen zur Konvertierung von oder nach Strings.
```go
	var i int
	i, err := strconv.Atoi("42")
	if err != nil {
		panic("not an integer")
	}
	fmt.Println(i)
```
    

## Funktionen
Normale Funktionen:
```go
func name(parameter1 type, parameter2 type) (returnParam1 type, returnParam2 type) {
 ..
}
```

Funktionen können wie andere daten zugewiesen werden:
```go
var hello = func(name string) {
	fmt.Println("Hello " + name)
}

var executer = func(name string, f func(name string)) {
	f(name)
}

executer("Marvin", hello)
```

## Kontrollstrukturen

### if {}
```go
if 2 > 1 {
	fmt.Println("1>2")
}

if data, err := readFromDatabase(); err != nil {
	fmt.Println("error reading data")
} else {
	fmt.Println(data)
}
```

### for {}
```go
colors := []string{"black", "red", "blue", "green", "white"}

// clasic for
for i := 0; i < len(colors); i++ {
	fmt.Printf("%v: %v\n", i, colors[i])
}

// iterate
for i, color := range colors {
	fmt.Printf("%v: %v\n", i, color)
}

// while
i := 0
for i < len(colors) {
	fmt.Printf("%v: %v\n", i, colors[i])
	i++
}

// while true
j := 0
for {
	if j >= len(colors) {
		break
	}
	fmt.Printf("%v: %v\n", j, colors[j])
	j++
}
```
    
### switch {}
```go
color := "nothing"
switch color {
case "green":
	fmt.Printf("Green")
case "red":
	fmt.Printf("Red")
default:
	fmt.Printf("Black")
}
```

## Übung 1: IDE
Probiere mehrere Editoren/IDEs aus und entscheide Dich für eine.

## Übung 2: Multiplikationstabelle

__Eingabe__: n, maximale Zahl als Faktor als Parameter

__Aufgabe__: Gibt eine formatierte Tabelle mit den Produkten aus

__Beispiel__: go run mult.go 4

```
      1   2   3   4
  1   1   2   3   4
  2   2   4   6   8
  3   3   6   9  12
  4   4   8  12  16
```


__Tipp__: `os.Args` enthält die Aufrufparameter


