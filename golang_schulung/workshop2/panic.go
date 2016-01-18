package main

import "fmt"

func main() {
	travel()
}

func travel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "..dont't panic!")
		}
	}()
	panic("I lost my towel")
}
