package main

import (
	"regexp"
)

func main() {
	terms := []string{"java", "is", "fun"}
	for _, value := range terms {
		rx := regexp.MustCompile("java")
		value = rx.ReplaceAllString(value, "golang")
		println(value)
	}
}
