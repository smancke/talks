package main

import (
	"fmt"
	"os"
)

func main() {
	var i = 42

	// direct output
	fmt.Print(i)
	fmt.Print("\n")

	// output with newline
	fmt.Println(i)

	// format values
	fmt.Printf("the answer is %v\n", i)

	// formating as string
	s := fmt.Sprintf("the answer is %v\n", i)
	fmt.Print(s)

	// write to writer, e.g. stderr
	fmt.Fprintf(os.Stderr, "the answer is %v\n", i)
}
