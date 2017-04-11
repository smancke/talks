package main

import "fmt"

type F func(string)

func main() {
	hello("Marvin")
	f := hello
	f("Marvin")
	executer("Marvin", f)
}

func hello(name string) {
	fmt.Printf("Hello %v\n", name)
}

func executer(name string, f F) {
	f(name)
}
