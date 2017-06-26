# Weitere Grundlagen

## defer
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

## Error Handling in Go
* Error als Rückgabewert

```go
	if err != nil {
		return nil, err
	}
```

### Besser: Errors wrappen

* Anhängen des Kontextes in dem ein Fehler aufgetreten ist
* Erhalten der Stack Trace Informationen

```go
package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
}
```

### Spezifische Errors

In manchen fällen kann es sein, bestimmte Fehler zurück zu geben.
Am häufigsten erfolgt dies über globale Variablen vom Typ error.

Sie https://golang.org/pkg/io/#pkg-variables

Alternativen:

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

## Panic, Recover
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

## Type conversion

Typen in Go können in kompatible Typen konvertiert werden.

```go
i := 42
f := float64(i)
u := uint(f)
```

Dies geht auch mit selbst definierten Typen.
```go
	s := "string"
	type myString string
	var ms myString
	ms = myString(s)
```


## Empty interface: `interface{}`

* Entspricht `Object` in Java oder `void *` in C.
* Variablen vom Typ `interface{}` können beliebige Werte aufnehmen


### Type assertion

Auf den konkreten Wert einer `interface{}` Variablen kann über eine Type Assertion zugegriffen werden.

```go
var o interface{} = 42
i := o.(int)
i++
```

Wenn eine Type Assertion fehl schlägt wir ein `panic()`ausgelöst.

Um ein Panic zum umgehen kann eine Type Assertion inkl. Test ausgeführt werden:

```go
i, ok := o.(int)
```

### Type switches
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

