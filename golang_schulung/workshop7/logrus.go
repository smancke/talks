package main

import (
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	start := time.Now()
	for {
		time.Sleep(time.Millisecond * 300)

		log.WithFields(log.Fields{
			"some":   rand.Int(),
			"random": rand.Int(),
			"values": time.Now().Sub(start),
		}).Info("Im doing nothing")

		log.WithFields(log.Fields{
			"other": rand.Int(),
		}).Error("Ups, something went wrong ..")
	}
}
