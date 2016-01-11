package main

import "fmt"

func main() {

	var hello = func(name string) {
		fmt.Println("Hello " + name)
	}

	var executer = func(name string, f func(name string)) {
		f(name)
	}

	executer("Marvin", hello)
}
