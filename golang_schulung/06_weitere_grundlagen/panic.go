package main

import "fmt"

func main() {
	travel()
}

func travel() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p, "..dont't panic!")
		}
	}()

	panic("I lost my towel")
	fmt.Println("never reached")
}
