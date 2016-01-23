# Workshop 3

# Benchmarking
# Goroutinen
# Das `sync` Package
## Mutex
## Wait Group

# Channel

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
