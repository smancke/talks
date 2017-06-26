package main

import "fmt"

func main() {
	a := 42
	var b int
	b = int(a)
	fmt.Println(b)

	var i interface{} = "42"

	j, ok := i.(int)
	if !ok {
		fmt.Println("no int value")
	}

	fmt.Println(j)
}
