package main

import (
	"regexp"
)

func main() {
	rx, err := regexp.Compile("java")
	if err != nil {
		panic(err)
	}

	terms := []string{"java", "is", "fun"}
	for _, value := range terms {
		println(rx.ReplaceAllString(value, "golang"))
	}
}
