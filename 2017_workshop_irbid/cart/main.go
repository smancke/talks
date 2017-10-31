package main

import (
	"fmt"
	"net/http"
)

func main() {

	cartData := map[string]string{}

	http.Handle("/cart", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			fmt.Fprintf(w, "<html><body>\n")
			fmt.Fprintf(w, "This is you cart")
			fmt.Fprintf(w, "  <ul>\n")

			for id, title := range cartData {

				fmt.Fprintf(w, " <li>%v (%v)</li>", title, id)

			}

			fmt.Fprintf(w, "  </ul>\n")
			fmt.Fprintf(w, "</body></html>\n")
		}

		if r.Method == "POST" {
			r.ParseForm()
			articleId := r.FormValue("articleId")
			articleTitle := r.FormValue("articleTitle")

			cartData[articleId] = articleTitle

			fmt.Fprintf(w, "You have added a %v to the cart", articleTitle)
		}
	}))
	panic(http.ListenAndServe(":80", nil))
}
