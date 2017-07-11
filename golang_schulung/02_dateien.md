
# 02 Dateien

## Das Pacakge `OS`
[https://golang.org/pkg/os/](https://golang.org/pkg/os/)
```go
var Args []string     // Commandline arguments

func Create(name string) (*File, error) // create a file
func Open(name string) (*File, error)   // open a file
func Getenv(key string) string          // get environment variable
func Exit(code int)                     // exit with return code
```

### Schreiben in eine Datei
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
```

### Exec Beispiel
```go
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: cmd command [args]\n")
	}
	c := exec.Command(os.Args[1], os.Args[2:]...)

	if out, err := c.Output(); err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("> %v", string(out))
	}
}
```

## Übung: cat Befehl
Implementiere die Basis-Funktionalität des `cat` Befehles unter Linux

## Übung: tac Befehl
Implementiere die Basis-Funktionalität des `tac` Befehles unter Linux

## Übung: wc Befehl
Implementiere die Basis-Funktionalität des `wc` Befehles unter Linux

__Tipp:__ Schau Dir mal `bufio.Scanner` an zum Einlesen von Daten an.

