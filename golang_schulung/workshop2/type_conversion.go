package main

import "fmt"
import "strconv"

func main() {
	untypedList := []interface{}{"Hallo", 42, false}

	for _, item := range untypedList {
		switch i := item.(type) {
		case string:
			fmt.Println("String: " + i)
		case int:
			i++
			fmt.Println("Int: " + strconv.Itoa(i))
		default:
			fmt.Println(i)
		}
	}
}
