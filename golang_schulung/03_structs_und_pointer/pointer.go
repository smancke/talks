package main

import "fmt"

func main() {

	var a int = 41
	var b *int

	b = &a
	*b++

	fmt.Println(a) // 42
}
