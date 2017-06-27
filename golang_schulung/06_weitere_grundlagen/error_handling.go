package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	object, err := readJsonObject("/tmp/nonExistingFile")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("some data: %v", object)

}

func readJsonObject(file string) (map[string]string, error) {
	data, err := readFile(file)
	if err != nil {
		return nil, err
	}

	o := map[string]string{}
	err = json.Unmarshal(data, &o)
	return o, err
}

func readFile(file string) ([]byte, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
