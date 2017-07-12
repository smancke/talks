package foo

import (
	"testing"
	"time"
)

func Benchmark_Something(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Sleep(time.Millisecond)
	}
}
