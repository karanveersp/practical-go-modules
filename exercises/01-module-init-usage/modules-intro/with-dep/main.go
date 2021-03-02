package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
)

func main() {
	data, err := ioutil.ReadFile("../data/person.json")
	if err != nil {
		panic(err)
	}

	value, _, _, err := jsonparser.Get(data, "city")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(value))
}
