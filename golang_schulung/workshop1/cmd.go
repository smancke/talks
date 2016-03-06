package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: cmd command [args]\n")
	}
	c := exec.Command(os.Args[1], os.Args[2:]...)

	if out, err := c.Output(); err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("> %s", out)
	}
}
