package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// copy by value
	person1 := Person{
		Name: "Mancke",
	}
	person2 := person1

	person2.Name = "Meyer"
	fmt.Println(person1.Name) // Mancke

	// copy by reference
	person3 := &person1

	person3.Name = "Meyer"
	fmt.Println(person1.Name) // Meyer
}
