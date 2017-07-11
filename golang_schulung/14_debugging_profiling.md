# Debugging und Profiling

## Race detection

Golang hat ein tool zur Suche nach race conditions direkt eingebaut.
```shell
go test -race mypkg    // test the package
go run -race mysrc.go  // compile and run the program
go build -race mycmd   // build the command
go install -race mypkg // install the package
```

## Debugging mit delve

Delve ist ein debugger extra für go.

### Installation unter Linux

Entweder:
```shell
git clone git@github.com:derekparker/delve.git
make install
```

oder:
```shell
go get -v -u github.com/derekparker/delve/cmd/dlv
```

### Installation OSX, Windows

https://github.com/derekparker/delve/wiki/Building

https://github.com/derekparker/delve/wiki/Tips-&-Troubleshooting


## Profiling

Gute Quelle: https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs

Go hat einen einfachen Profiler eingebaut.

```go
import _ "net/http/pprof"
..
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

```shell
### heap usage snapshot
go tool pprof http://localhost:6060/debug/pprof/heap

### 10sec cpu profiling snapshot
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

### show top cpu consumer ordered by cumulative consumption
(pprof) top -cum

### create an svg graph
(pprof) web

### create an svg graph for the calltree of doSomething()
(pprof) web doSomething

### Display callers and callees of doSomething()
(pprof) peek doSomething
```

### Blocking
Blockieren von Go Routinen

Aktivieren mit ..
```go
runtime.SetBlockProfileRate(50)
```

Analyse mit:
```shell
go tool pprof http://localhost:6060/debug/pprof/heap
```

### Zustand der Goroutinen

Analyse mit:
```shell
go tool pprof  http://localhost:6060/debug/pprof/goroutine
```

## GODEBUG
Die Umgebungsvariable `GODEBUG` kann gesetzt werden um verschiedene Ausgaben zu schreiben.

### Garbage collector 
Aktivierung über
```shell
GODEBUG=gctrace=1
```

Ausgabe Beispiel:
```
gc 7 @0.444s 9%: 0.12+0.24+0.017+69+0.37 ms clock, 0.49+0.24+0+7.5/68/117+1.4 ms cpu, 74->78->65 MB, 79 MB goal, 4 P
```

Felder:
```
gc # @#s #%: #+...+# ms clock, #+...+# ms cpu, #->#-># MB, # MB goal, # P

gc #        the GC number, incremented at each GC
@#s         time in seconds since program start
#%          percentage of time spent in GC since program start
#+...+#     wall-clock/CPU times for the phases of the GC
#->#-># MB  heap size at GC start, at GC end, and live heap
# MB goal   goal heap size
# P         number of processors used
```

### Scheduler tracen
```shell
GODEBUG=schedtrace=1000
```
Ausgabe Beispiel:
```
SCHED 1004ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=4 runqueue=8 [0 1 0 3]
SCHED 2005ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=5 runqueue=6 [1 5 4 0]
SCHED 3008ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=4 runqueue=10 [2 2 2 1]
```

# runtime/trace
Erzeugen eines Tracefiles:
```go
f, err := os.Create("trace_example.trace")
if err != nil {
	panic(err.Error())
}
defer f.Close()
trace.Start(f)
defer trace.Stop()
```

Ausführen:
```shell
go build  trace_example.go 
./trace_example 
go tool trace trace_example trace_example.trace
```
