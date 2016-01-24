# Workshop 3

# Goroutinen
Goroutinen sind eine leichtgewichtige Alternative zu Theads.
Sie werden von der Go Runtime auf echte Threads verteilt.

Test auf *i5-2520M*:

__1 Mio Goroutinen lassen sich in unter 2 Sekunden ausführen.__

(Vgl.: 1 Mio Java Threads: ~40 Sekunden)

## Goroutinen Syntax
```go
go doInBackground()
```

Oder inline definiert:
```go
go func() {
	for i := 0; i < 100; i++ {
		fmt.Println("in background inline")
	}
}()
```
    
## Maximale Threads begrenzen
Die Anzahl maximal gleichzeitig aktiver Threads kann über 
`runtime.GOMAXPROCS(n int)` gesetzt werden.

Default ist `runtime.NumCPU()`.

## Sheduling
Die Ausführung einer laufenden Goroutine wird vom Sheduler nur sehr selten unterbrochen.
Durch ein `runtime.Gosched()`, ein `time.Sleep()` oder blockierende IO Operationen
wechselt der Sheduler jedoch direkt auf eine andere Goroutine.

# Channel

*Do not communicate by sharing memory; instead, share memory by communicating.*

## Basics
* Ein Channel ist eine typisierte fifo-Queue mit fester Länge
* Der Channel kann Daten beliebigen Typs aufnehmen
* Alle Operationen auf Channel sind robust gegenüber paralleler Zugriffe

Erstellen eines Channels: `make (chan DataType, size)`
Schreiben in den Channel: `ch <- value`
Lesen vom Channel: `value <- ch`


## Unbuffered Channel
* Operationen auf einen Channel blockieren
* Lesen wartet, bis Daten vorliegen
* Schreiben wartet, bis daten abgegeben werden können

### Beispiel: Schreiben in separater Goroutine:
```go
ch := make(chan string)

go func() {
	ch <- "The Answer is "
	ch <- "42"
}()

fmt.Println(<-ch)
fmt.Println(<-ch)
```

### Beispiel: Timer
```go
timeoutChannel := time.After(time.Second)
<-timeoutChannel
fmt.Println("One second is elapsed")
```

## Buffered Channel
Ein channel kann eine Buffer-Size besitzen.
Schreiben blockiert nicht, wenn der Channel noch Platz hat

Beispiel:
```go
ch := make(chan string, 2)

ch <- "The Answer is "
ch <- "42"
// ch <- "one more write would block!"

fmt.Println(<-ch)
fmt.Println(<-ch)
```

## Schließen eines Channel
Ein Channel kann geschlossen werden.
```go
close(ch)
```

* Schreiben an einen geschlossenen Channel erzeugt ein __panic()__
* Lesen aus einem geschlossenen Channel kehrt sofort zurück (mit oder ohne Wert).
* Es kann abgefragt werden, ob ein Channel einen Wert zurück geliefert hat.

```go
value, channelWasOpen <- ch
```

Beispiel:
```go
ch := make(chan string)

go func() {
	ch <- "The Answer is "
	ch <- "42"
	close(ch)
}()

for {
	msg, channelOpen := <-ch
	if !channelOpen {
		break
	}
	fmt.Println(msg)
}
```

## Mit range über Channels iterieren
* Bei einem Channel liefert `range` nur einen Wert zurück.
* Range liest blockierend, bis der Channel geschlossen wird.

```go
ch := make(chan string, 2)

ch <- "The Answer is "
ch <- "42"
close(ch)

for v := range ch {
	fmt.Printf("v=%v\n", v)
}
```

## Select
* Sie `select` Anweisung kann verwendet werden um mehrere Channel Operationen in einem durch zu führen.
* Bei mehreren Case-Zweigen wird der Zweig ausgeführt, der als erster verfübar wird.
* Sind mehrere Zweise Verfügbar, so ist die Reihenfolge zufällig.
* Der `default`-Zweig wird ausgeführt, wenn kein weiterer Zweig verfügbar ist.


```go
func readNonBlocking() {
	ch := make(chan string, 2)

	ch <- "The Answer is "
	ch <- "42"

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Println("no input available.")
			return
		}
	}
}
```

```go
func sendWithTimeout() {
	ch := make(chan string)

	select {
	case ch <- "a message":
	case <-time.After(time.Millisecond):
		fmt.Println("timeout on writing to channel.")
		return
	}
}
```

## Channel Tricks: Exit on Signal
```go
func waitForTermination(callback func()) {
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Got singal '%v' .. exit now", <-sigc)
	callback()
	os.Exit(0)
}
```

## Channel Tricks: Close als Wait-Broadcast
Da lesen von einem Geschlossenen Channel direct zurück kehrt kann dies als
Broadcast Signal verwendet werden.

```go
func startWorker(name string, startSignal chan bool) {
	<-startSignal
	fmt.Printf("Worker %v got start singal\n", name)
}

func main() {

	ch := make(chan bool)

	go startWorker("Worker 1", ch)
	go startWorker("Worker 2", ch)
	go startWorker("Worker 3", ch)

	close(ch)

	time.Sleep(time.Millisecond)
}
```

## Channel Tricks: Channel mit Callback
Wenn auf eine Anfrage an einen Channel eine Antwort erwartet wird kann es praktisch sein,
einen Callback Channel mit zu übergeben.

```go
type request struct {
	msg      string
	callback chan string
}

func echoRoutine(requestChannel chan request) {
	for {
		request := <-requestChannel
		request.callback <- request.msg
	}
}
```

# Das `sync` Package
Das `sync` Package enthält kleine Helper Concurrency Aufgaben.

## Mutexes
Zum expliziten Locken des Zugriffes auf Daten kann ein  Mutex verwendet werden.

Exclusiver Lock:
```go
var mutex = &sync.Mutex{}
mutex.Lock()
mutex.Unlock()
```

Read Write Lock:
```go
var rwmutex = &sync.RWMutex}
rwmutex.Lock()
rwmutex.RLock()
rwmutex.RUnlock()
rwmutex.Unlock()
```

## Wait Group
Eine Wait Group kann verwendet werden um auf eine Menge von Jobs zu warten:
```go
func doInBackground(waitGroup *sync.WaitGroup) {
	fmt.Println("do in backgroud")
	waitGroup.Done()
}

func main() {
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go doInBackground(waitGroup)
	}

	waitGroup.Wait()

	fmt.Printf("done.")
}
```

# Benchmarks

## Benchmarks: Grundidee
* Benchmarks sind Funktionen in den Test Dateien mit der Signatur: `func Benchmark_*(b *testing.B)`
* Der Testcode wird in einer Schleife `b.N` mal wiederholt.
* Abhängig von der Ausführungszeit für Go mehrere Tests mit unterschiedlichen Stichproben durch (z.B. 100, 10000, 1000000).

## Benchmarks: Hilfsfunktionen
Damit Hilfscode nicht mit gemessen wird, kann der Timer explizit gesteuert werden:

* `b.ResetTimer()` - Setz the timer zurück
* `b.StartTimer()` - (Re)Start des Timers
* `b.StopTimer()` - Hält den Timer an

## Benchmarks: Beispiel

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

## Benchmarks ausführen

Die Ausführung erfolgt über `go test -bench <regex>`

Beispiel:
```shell
go test -bench '.*' goroutine_lots_of_test.go
```

# Übungen

## Übung 5a: Concurrent Key-Value Store Access
Sichere Deine Key-Value Store Klasse so ab, dass sie parallelem Zugriff Stand hält.

## Übung 5b: Concurrency Test
Teste die Klasse mit vielen parallelen Reads und Writes.

## Übung 6: Benchmarking des Key-Value Stores
Messe die Zeit, die Dein KV-Store braucht um einen Durchlauf von *Schreiben-Speichern-Lesen* durch zu führen.
