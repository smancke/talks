package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	resp, err := http.Get("http://tarent.de/")
	if err != nil {
		panic(err)
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(doc.Find("title").Text())
}
