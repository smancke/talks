package main

import "fmt"

type Name string

type F func(Name)

func main() {

	hello := func(name Name) {
		fmt.Printf("Hello %v\n", name)
	}

	executer := func(name Name, f F) {
		f(name)
	}

	executer(Name("Marvin"), hello)
}
