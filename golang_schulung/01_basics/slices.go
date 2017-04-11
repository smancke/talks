package main

import "fmt"

func main() {

	farben := []string{"black", "red", "blue"}
	farben = append(farben, "green", "white")
	farben = farben[1 : len(farben)-1]

	fmt.Println(len(farben))
	fmt.Println(farben)
}
