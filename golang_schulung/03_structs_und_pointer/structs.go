package main

import "fmt"

type Adresse struct {
	Street string
}

type Person struct {
	Name     string
	Given    string
	Age      int
	Location Adresse
}

func main() {

	person := Person{
		Name:     "Mancke",
		Given:    "Sebastian",
		Age:      42,
		Location: Adresse{Street: "Rochusstraße"},
	}
	fmt.Println(person)

	person.Given = "Felix"
	fmt.Println(person)

}
