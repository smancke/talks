package main

import "fmt"

func main() {

	farben := [5]string{"black", "red", "blue", "green", "white"}

	bunt := farben[1 : len(farben)-1]
	fmt.Println(bunt)

	bunt = append(bunt, "orange")
	fmt.Println(farben)
	fmt.Println(bunt)

	miniSlice := farben[1:2]
	fmt.Println(len(miniSlice))
	fmt.Println(cap(miniSlice))

	farben2 := make([]string, 0, 5)
	fmt.Println(farben2)
}
