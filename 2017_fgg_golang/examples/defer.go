package main

import (
	"os"
)

func main() {

	file, err := os.Create("/tmp/foo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
