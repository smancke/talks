package main

import "fmt"

func main() {

	color := "nothing"
	switch color {
	case "green":
		fmt.Printf("Green")
	case "red":
		fmt.Printf("Red")
	default:
		fmt.Printf("Black")
	}
}
