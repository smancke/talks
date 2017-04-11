package main

import (
	"fmt"
)

func main() {

	if 2 > 1 {
		fmt.Println("1>2")
	}

	if data, err := readFromDatabase(); err != nil {
		fmt.Println("error reading data " + err.Error())
	} else {
		fmt.Println(data)
	}

}

func readFromDatabase() (string, error) {
	return "", fmt.Errorf("foo %v", "bar")
}
