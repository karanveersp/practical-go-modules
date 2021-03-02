package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Person is a person
type Person struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

func main() {
	t, _ := ioutil.ReadFile("../data/person.json")
	var p Person
	json.Unmarshal(t, &p)
	fmt.Printf("%v", p)
}
