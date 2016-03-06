package main

import "fmt"

func main() {

	var hello = func(name string) {
		fmt.Println("Hello " + name)
	}

	var executer = func(name string, f func(string)) {
		f(name)
	}

	executer("Marvin", hello)
}
