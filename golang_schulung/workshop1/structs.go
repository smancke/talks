package main

import "fmt"

type Person struct {
	Name  string
	Given string
	Age   int
}

func main() {

	person := Person{
		Name:  "Mancke",
		Given: "Sebastian",
		Age:   42,
	}
	fmt.Println(person)

	person.Given = "Felix"
	fmt.Println(person)
}
