package main

import (
	"github.com/PuerkitoBio/goquery"
)

func main() {

	doc, err := goquery.NewDocument("https://www.google.de/search?q=golang")
	if err != nil {
		panic(err)
	}

	println("\nGolang top search matches:\n")
	doc.Find(".r a").Each(func(i int, s *goquery.Selection) {
		println(" - " + s.Text())
	})
}
