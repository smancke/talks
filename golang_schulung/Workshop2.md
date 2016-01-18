# Workshop 2

# defer
------------------------------
* `defer`: Verzögert die Ausführung bis an das Ende der aktuellen Funktion
* Die Ausführung wird garantiert (== finally{})

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
```

# Panic, Recover
* `panic`: Löst ein Panic aus (ähnlich einer Exception)
* `recover`: Fängt ein panic ab

```go
func travel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "..dont't panic!")
		}
	}()
	panic("I lost my towel")
}
```
# Type switches
* `interface{}` entspricht dem generischen Typ (Object bzw. void*).
* Mit type switches lässt sich elegant in einem Schritt prüfen und konvertieren.

```go
untypedList := []interface{}{"Hallo", 42, false}

for _, item := range untypedList {
	switch i := item.(type) {
	case string:
		fmt.Println("String: " + i)
	case int:
		i++
		fmt.Println("Int: " + strconv.Itoa(i))
	default:
		fmt.Println(i)
	}
}
```

# Testen von Go code

## Tests
* Alle Dateien mit der Endung `_test.go` beinhalten testcode
* `go test <package>`
* Tests sind Funktionen mit der Signatur: `func Test_*(t *testing.T)`

```go
package foo

import "testing"

func Test_Simple(t *testing.T) {

	t.Logf("This Test fails")
	t.Fail()

}
```

## stretchr/testify
* Einfache Library mit assertion Funktionen
* Installieren von Test-Abhängigkeiten mit `go get -t`

```go
package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_With_Testify(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1)
}
```

## Test mehrerer Packages
Mehrere Packages können auf einmal getestet werden: `go test package/...`

Beispiel:
```shell
go test -cover github.com/smancke/guble/...
?       github.com/smancke/guble        [no test files]
ok      github.com/smancke/guble/client 0.125s  coverage: 80.3% of statements
ok      github.com/smancke/guble/gcm    0.129s  coverage: 33.7% of statements
ok      github.com/smancke/guble/guble  0.011s  coverage: 96.3% of statements
?       github.com/smancke/guble/guble-cli      [no test files]
ok      github.com/smancke/guble/gubled 0.194s  coverage: 58.5% of statements
ok      github.com/smancke/guble/server 0.151s  coverage: 83.2% of statements
ok      github.com/smancke/guble/store  1.295s  coverage: 78.1% of statements
```

## Coverage anschauen
Anschauen der Testcoverage mit `go tool cover`

```shell
go test -cover -coverprofile cover.out  github.com/smancke/guble/server
go tool cover -html=cover.out
```

# Objekte

## Methoden
* Methoden sind Funktionen, die eine Variable als *Receiver* haben.
* Als Receiver-Typen können nur eigene Typen des selben packages verwendet werden.

```go
type Point [2]int

func (p Point) Add(pointToAdd Point) Point {
	p[0] = p[0] + pointToAdd[0]
	p[1] = p[1] + pointToAdd[1]
	return p
}
```

## Objekte
* Damit Methoden die Daten eines Objektes verändern können, müssen sie den Pointer-Typ als Receiver haben.
* Meist werden Structs als Grundlage für *Klassen* verwendet.

```go
type Item struct {
	Name string
	pos  Point
}

func (item *Item) MoveTo(vector Point) {
	item.pos = item.pos.Add(vector)
}
```

## Konstruktoren
* In Go gibt es keine expliziten Konstruktoren.
* Konvention ist es, als Konstruktor eine Funktion `NewTypname() *Typename` bereit zu stellen.
* Häufig reichen jedoch auch die Default-Werte eines Type als Initialisierung aus (Doku lesen).


```go
func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (item *Item) MoveInDirection(vector Point, time int) {
	for i := 0; i < time; i++ {
		item.pos = item.pos.Add(vector)
	}
}
```

## Embedding
* Structs in Go können andere Structs einbetten.
* An das eingebettete Struct wird automatisch delegiert.
```go
type Starship struct {
	Item
	Speed int
}

func NewStarship(name string, speed int) *Starship {
	ship := &Starship{
		Speed: speed,
	}
	ship.Item.Name = name
	return ship
}

func travel() {
  herzAusGold := NewStarship("HerzAusGold", 42)
  herzAusGold.MoveTo(Point{1, 1})
}
```

## Überschreiben
* Methoden eines eingebetten Type können auch überschrieben werden

```go
func (ship *Starship) MoveInDirection(vector Point, time int) {
	for i := 0; i < ship.Speed; i++ {
		ship.Item.MoveInDirection(vector, time)
	}
}
```

# Interfaces
* Go kennt keine Vererbung, aber Interfaces.
* Interfaces folgen dem Duck-Typing Ansatz: Was aussieht wie eine Ente, ist auch eine Ente!
* Der Consumer legt das Interface fest, nicht der implementierer.

```go

type myInt int

func (i myInt) String() string {
	return strconv.Itoa(int(i))
}

type Printable interface {
	String() string
}

func Test_Stringer(t *testing.T) {
	printable := []Printable{
		NewItem("Atom"),
		NewStarship("HerzAusGold", 42),
		myInt(42),
		os.ModeAppend | os.ModeSocket,
	}

	for _, p := range printable {
		fmt.Println(p.String())
	}
}
```

## Übung 3: Key-Value Objekt Orientiert
Baue Deinen KV-Store so um, dass er intern eine Klasse Store verwendet,
die die Operationen auf den internen Storage abstrahiert.

## Übung 4: Testen
Teste die Klasse Store.
