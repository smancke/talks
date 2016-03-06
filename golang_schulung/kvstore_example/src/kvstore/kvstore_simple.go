package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file := "store_simple.db"
	store, err := ReadStore(file)
	if err != nil {
		fmt.Printf("error opening store: %v", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		WriteTo(store, os.Stdout)
		os.Exit(0)
	}

	shouldSave := false
	for _, arg := range os.Args[1:] {
		if modified := handleCommand(store, arg); modified {
			shouldSave = modified
		}
	}

	if shouldSave {
		if err := WriteStore(store, file); err != nil {
			fmt.Printf("error writing store: %v", err.Error())
			os.Exit(2)
		}
	}
}

func handleCommand(store map[string]string, arg string) (modified bool) {
	if strings.Contains(arg, "=") {

		kv := strings.SplitN(arg, "=", 2)
		store[kv[0]] = kv[1]
		return true

	} else {

		if v, exist := store[arg]; exist {
			fmt.Printf("%v=%v\n", arg, v)
		} else {
			fmt.Printf("%v NOT FOUND\n", arg)
		}
		return false
	}
}

func ReadStore(file string) (map[string]string, error) {
	store := make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		handleCommand(store, scanner.Text())
	}

	return store, scanner.Err()
}

func WriteStore(store map[string]string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return WriteTo(store, f)
}

func WriteTo(store map[string]string, w io.Writer) error {
	for k, v := range store {
		if _, err := fmt.Fprintf(w, "%v=%v\n", k, v); err != nil {
			return err
		}
	}
	return nil
}
