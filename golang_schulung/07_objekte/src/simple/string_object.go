package main

import (
	"fmt"
)

type S string

func (s S) Println() {
	fmt.Println(s)
}

func main() {

	s := S("foo")
	s.Println()

}
