package main

import "fmt"

func main() {

	colors := []string{"black", "red", "blue", "green", "white"}

	// clasic for
	for i := 0; i < len(colors); i++ {
		fmt.Printf("%v: %v\n", i, colors[i])
	}

	// iterate
	for i, color := range colors {
		fmt.Printf("%v: %v\n", i, color)
	}

	// while
	i := 0
	for i < len(colors) {
		fmt.Printf("%v: %v\n", i, colors[i])
		i++
	}

	// while true
	j := 0
	for {
		if j >= len(colors) {
			break
		}
		fmt.Printf("%v: %v\n", j, colors[j])
		j++
	}

}
