package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handle)
	panic(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]

	switch r.Method {
	case "POST":
		bytes, _ := ioutil.ReadAll(r.Body)
		if err := save(filePath, bytes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}

	case "GET":
		if err := copyTo(filePath, w); err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	default:
		http.Error(w, fmt.Sprintf("Method not supported %v\n", r.Method), http.StatusBadRequest)
	}
}

func copyTo(path string, w io.Writer) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func save(path string, bytes []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0644)
}
