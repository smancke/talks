package main

import (
	"fmt"
	"time"
)

type Measure struct {
	name     string
	start    time.Time
	lastSeen time.Time
	count    int
}

func StartMeasure(name string) *Measure {
	return &Measure{
		start:    time.Now(),
		lastSeen: time.Now(),
		name:     name,
	}
}

func (m *Measure) Action() {
	m.count++
	if time.Since(m.lastSeen) > time.Millisecond*100 {
		fmt.Printf("\r%v: %v", m.name, m.count)
		m.lastSeen = time.Now()
	}
}

func (m *Measure) End() {
	d := time.Since(m.start)
	if m.count > 0 {
		fmt.Printf("\r%v: %v in %v (%v/sec)\n", m.name, m.count, d, int(float64(m.count)/d.Seconds()))
	} else {
		fmt.Printf("%v: %v in %v\n", m.name, m.count, d)
	}
}
