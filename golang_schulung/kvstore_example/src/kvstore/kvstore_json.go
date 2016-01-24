package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := "store.json"
	store, err := ReadStore(file)
	if err != nil {
		fmt.Printf("error opening store: %v", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		for k, v := range store {
			fmt.Printf("%v=%v\n", k, v)
		}
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

func handleCommand(store map[string]string, arg string) bool {
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

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return store, nil
	}

	bytes, err := ioutil.ReadFile(file)
	if err == nil {
		err = json.Unmarshal(bytes, &store)
	}
	return store, err
}

func WriteStore(store map[string]string, file string) error {
	bytes, err := json.Marshal(&store)
	if err == nil {
		err = ioutil.WriteFile(file, bytes, 0644)
	}
	return err
}
