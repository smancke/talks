package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func main() {

	object, err := readJsonObject("/tmp/nonExistingFile")
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	fmt.Printf("some data: %v\n", object)

}

func readJsonObject(file string) (map[string]string, error) {
	data, err := readFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "read file %s", file)
	}

	o := map[string]string{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return nil, errors.Wrapf(err, "json parse file %s", file)
	}
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
