package main

import "fmt"

type F func(Name)

type Name string

func main() {

	var hello = func(name Name) {
		fmt.Printf("Hello %v\n", name)
	}

	var executer = func(name Name, f F) {
		f(name)
	}

	executer(Name("Marvin"), hello)
}
