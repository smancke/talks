
# Testen von Go Code

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

## Data Driven Tests
```
func Test_Operations(t *testing.T) {
	tests := []struct {
		title    string
		op       func(int, int) int
		a        int
		b        int
		expected int
	}{
		{"add", add, 0, 0, 0},
		{"add", add, 0, 1, 1},
		{"add", add, 1, 2, 3},
		{"add", add, 2, 2, 4},
		{"add", add, 42, 42, 42}, // this one fails
		{"sub", sub, 0, 0, 0},
		{"sub", sub, 0, 1, -1},
		{"sub", sub, 42, 42, 42}, // this one fails
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v(%v, %v) == %v", test.title, test.a, test.b, test.expected)
		t.Run(testName, func(t *testing.T) {
			actual := test.op(test.a, test.b)
			if actual != test.expected {
				t.Logf("expected %v, but got %v", test.expected, actual)
				t.Fail()
			}
		})
	}
}
```

## Benchmarks

### Benchmarks: Grundidee
* Benchmarks sind wie auch die Tests Funktionen, die in den Test-Dateien stehen.
* Sie haben die Signatur: `func Benchmark_*(b *testing.B)`
* Der Testcode wird in einer Schleife `b.N` mal wiederholt.
* Abhängig von der Ausführungszeit führt Go mehrere Tests mit unterschiedlichen Stichproben durch (z.B. 100, 10000, 1000000).

### Benchmarks: Hilfsfunktionen
Damit Hilfscode nicht mit gemessen wird, kann der Timer explizit gesteuert werden:

* `b.ResetTimer()` - Setz the timer zurück
* `b.StartTimer()` - (Re)Start des Timers
* `b.StopTimer()` - Hält den Timer an

### Benchmarks: Beispiel

```go
func Benchmark_Creation_Of_Goroutines(b *testing.B) {
	fmt.Printf("testing with %v goroutines\n", b.N)
	doneChannel := make(chan bool)

	for i := 0; i < b.N; i++ {
		go doInBackground(doneChannel)
	}

	for i := 0; i < b.N; i++ {
		<-doneChannel
	}
}
```

### Benchmarks ausführen

Die Ausführung erfolgt über `go test -bench <regex>`

Beispiel:
```shell
go test -bench '.*' goroutine_lots_of_test.go
```

## Übung: Taschenrechner Programmes

* Schreibe einen einfachen Konsolen Taschenrechner 
* Teste die Funktionalität des Rechner-Programmes vollständig.

## Übung: Benchmarking des Key-Value Stores
Messe die Zeit, die Dein KV-Store braucht um einen Durchlauf von *Schreiben-Speichern-Lesen* durch zu führen.
