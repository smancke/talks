package main

import "fmt"

func main() {

	colors1 := []string{"red", "blue"}

	colors2 := colors1
	colors2[0] = "black"
	colors2[1] = "white"

	fmt.Println(colors1) // black, white

}
