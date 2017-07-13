# Native Development

## CGO
CGO erlaubt es Go und C code zu verbinden.

* CGO Code wird automatisch im Build mit übersetze
* Go Typen müssen in C-Typen konvertiert werden

```
/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
```

[Weitere Dokumentation](https://github.com/golang/go/wiki/cgo)

## Shared Libraries verwenden
Go ist in der Lage shared libraries zu laden und Funktionen darin aufzurufen.

## Shared Libraries erstellen

Mit
```
-buildmode=c-shared
```

können auch shared libraries erstellt werden.


## Plugins
Seit Go 1.8 können auch dynamisch shared libraries mit go plugins geladen werden:

https://jeremywho.com/go-1.8---plugins/

## Cross Compilation


### Windows exe erzeugen:
```
GOOS=windows GOARCH=386 go build -o out.exe .
```

### Übersetzen für Arm7
```
GOARCH=arm GOARM=7 go build -o out .
```
