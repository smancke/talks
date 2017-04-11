
## Testen von Go code

### Tests
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

### stretchr/testify
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

### Test mehrerer Packages
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

### Coverage anschauen
Anschauen der Testcoverage mit `go tool cover`

```shell
go test -cover -coverprofile cover.out  github.com/smancke/guble/server
go tool cover -html=cover.out
```

## Übung: Testing des Rechner-Programmes

Teste die Funktionalität des Rechner-Programmes vollständig.
