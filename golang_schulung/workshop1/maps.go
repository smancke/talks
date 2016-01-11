package main

import "fmt"

func main() {

	person := map[string]string{
		"name":  "Mancke",
		"given": "Sebastian",
	}
	fmt.Println(person)
	fmt.Println(len(person))

	person["given"] = "Felix"
	fmt.Println(person)
	delete(person, "given")

	_, exist := person["given"]
	fmt.Println(exist) // false
}
