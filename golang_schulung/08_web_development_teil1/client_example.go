package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("http://tarent.de/")
	if err != nil {
		panic(err)
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(doc.Find("title").Text())

	// POST example
	values := url.Values{}
	values.Add("foo", "bar") // foo=bar&..
	r, _ := http.NewRequest("POST", "http://tarent.de/", bytes.NewBufferString(values.Encode()))
	r.Header.Set("Content-Tyoe", "application/x-www-form-urlencoded")

	resp, err = http.DefaultClient.Do(r)
	fmt.Printf("code: %v, err: %v\n", resp.StatusCode, err)
}
