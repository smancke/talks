package main

import "fmt"

func main() {

	color := "blue"
	switch color {
	case "green", "blue":
		fmt.Printf("Green")
	case "red":
		fmt.Printf("Red")
	default:
		fmt.Printf("Black")
	}

}
