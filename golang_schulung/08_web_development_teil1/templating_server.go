package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var theTemplate = `<html>
<body>
  <h1>Hello {{.}}</h1>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowedBadRequest)
			fmt.Fprintf(w, "Only GET is allowed\n")
			return
		}

		t := template.New("template")
		template.Must(t.Parse(theTemplate))

		name := r.FormValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Parameter name is mising\n")
			return
		}

		err := t.Execute(w, name)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error on executing template\n")
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
