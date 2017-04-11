package main

import "fmt"

func main() {

	farben := []string{"black", "red", "blue"}
	fmt.Println(cap(farben))
	farben = append(farben, "green")
	fmt.Println(cap(farben))
	farben = farben[1 : len(farben)-1]

	fmt.Println(len(farben))
	fmt.Println(cap(farben))
	fmt.Println(farben)
}
