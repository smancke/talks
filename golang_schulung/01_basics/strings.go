package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	b := []byte{72, 97, 108, 108, 111, 32, 87, 101, 108, 116}
	s := string(b)
	fmt.Println(s)

	fmt.Println(s[0:5])

	fmt.Println(strings.HasPrefix(s, "Hallo"))
	fmt.Println(strings.ToLower(s))

	var i int
	i, err := strconv.Atoi("42x")
	if err != nil {
		panic("not an integer")
	}
	fmt.Println(i)
}
