package main

import (
	"fmt"
)

func main() {

	if 2 > 1 {
		fmt.Println("1>2")
	}

	if data, err := readFromDatabase(); err != nil {
		fmt.Println("error reading data")
	} else {
		fmt.Println(data)
	}
}

func readFromDatabase() (string, error) {
	return "some data", nil
}
