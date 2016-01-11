package main

import "fmt"

func main() {

	a := 41
	var b *int

	b = &a
	*b++

	fmt.Println(a) // 42
}
