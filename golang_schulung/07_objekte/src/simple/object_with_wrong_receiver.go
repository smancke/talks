package main

import (
	"fmt"
)

type I int

func (i I) Add(b int) {
	i += I(b)
}

func main() {

	i := I(20)
	i.Add(22)
	fmt.Println(i)

}
