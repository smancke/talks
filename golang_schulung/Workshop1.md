
# Workshop 1

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

## Projektstruktur
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

### Datentypen & Variablen
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
* int und float Typen in verschiedenen größen
* bool
* arrays mit fester Länge `[4]int{42, 43, 44, 45}`
* Slices als Abstraktion über Arrays: `[]int{42, 43, 44, 45}`
* string (Entspricht einem byte Slice)
* Maps `map[string]int{"a": 42, "b": 43}`
* Structs

### Array & Slices
* Arrays haben eine feste Größe
* Slices sind eine Referenz auf ein Array mit einem Start- und Index

```go
farben := [5]string{"black", "red", "blue", "green", "white"}

bunt := farben[1 : len(farben)-1]
fmt.Println(bunt)

bunt = append(bunt, "orange")
fmt.Println(bunt)

fmt.Println(len(farben[0:1])) // length: 1
fmt.Println(cap(farben[0:1])) // capacity: 5
```

Explizite Schreibweise:
```go
slicename := make(type, len, cap)

// example:
farben := make([]string, 0, 5)
```

### Maps
```go
m := make(map[KeyType]StringType)
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

### Structs
Typdefinition:
```go
type Person struct {
	Name  string
	Given string
	Age   int
}
```

Verwendung:
```go
person := Person{
	Name:  "Mancke",
	Given: "Sebastian",
	Age:   42,
}
fmt.Println(person)

person.Given = "Felix"
fmt.Println(person)
```

### Pointer
* Parameterübergabe immer Call-by-Value (auch bei struct)
* Typsichere Pointer können verwendet werden
```go
a := 41

var b *int
b = &a

*b++

fmt.Println(a) // 42
```

### Pointer und Structs
Referenzen auf structs:
```go
// copy by value
person1 := Person{
	Name: "Mancke",
    }
    
person2 := person1
person2.Name = "Meyer"

fmt.Println(person1.Name) // Mancke

// copy by reference
person3 := &person1
person3.Name = "Meyer"

fmt.Println(person1.Name) // Meyer
```

### Pointer und Slices oder Maps
Mit make erstellte Datentypen sind immer Pointer-Typen.

```go
colors1 := []string{"red", "blue"}

colors2 := colors1
colors2[0] = "black"
colors2[1] = "white"
    
fmt.Println(colors1) // black, white
```

### Funktionen
Normale Funktionen:
```go
func name(parameter1 type, parameter2 type) (returnParam1 type,  returnParam2  type) {
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
}
```

## Kontrollstrukturen - if
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

## Kontrollstrukturen - for
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
    
## Kontrollstrukturen - switch/case
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

# Schreiben in eine Datei
------------------------------
* Verwendung des Packages os
* defer statement werden am Ende der funktion ausgeführt (== finally{})

```go
import "os"

func main() {
	file, err := os.Create("/tmp/hello")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello World\n")
}

## Übung 1: IDE
Probiere mehrere Editoren/IDEs aus und entscheide Dich für eine.


## Übung 2: Key-Value Store
Schreibe einen kleines Programm `kv`, mit dem Du Schlüssel-Werte
Paare in einer Datei speichern und abfragen kannst.

Setzen von Werten:
```shell
kv name=Mancke vorname=Sebastian alter=42
```

Abfragen bestimmter Werte
```shell
kv name vorname
> name=Mancke
> vorname=Sebastian
```

Abfragen aller Werte
```shell
kv
> name=Mancke
> vorname=Sebastian
> alter=42
```

