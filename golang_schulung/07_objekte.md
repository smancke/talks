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

## Interfaces
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

## Mocking

Mocking kann in go nicht dynamisch erfolgen. Mock-Frameworks verlangen eine Mock-Generierung zur Entwicklungszeit. Die Mocks werden dann üblicherweise mit eingecheckt. Z.b. [GoMock](https://github.com/golang/mock)

```
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockMyInterface(ctrl)

	mock.EXPECT(). ...
```
    
Code Generierung in go:
```
//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -self_package objects -package objects -destination $GOPATH/src/objects/mocks_test.go objects Flyable
func ...
```

```
$ go generate .
```

## Übungen

### Übung: Key-Value Objekt Orientiert
Baue Deinen KV-Store so um, dass er intern eine Klasse Store verwendet,
die die Operationen auf den internen Storage abstrahiert.

### Übung: Testen mit Mocks
* Teste die Klasse Store
* Teste den aufrufenden Code der Klasse Store, gegen ein Mock-Objekt 
