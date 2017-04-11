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
## Type switches
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

