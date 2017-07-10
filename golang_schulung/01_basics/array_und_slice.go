package main

import "fmt"

func main() {

	farben := [5]string{"black", "red", "blue", "green", "white"}

	bunt := farben[1 : len(farben)-1]
	fmt.Println(bunt)

	bunt = append(bunt, "orange")
	fmt.Println(bunt)
	fmt.Println(farben)
}
