package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const FILENAME = "budget.pb"

func main() {
	fmt.Println("----------------------")
	fmt.Println("------Plan Better-----")
	fmt.Println("----------------------")

	a := ReadFile()
	a.Print()

	Transfer(a.Envelopes["Groceries"], a.Envelopes["Utilities"], 600)

	Save(a)
}

func Save(a Account) {
	bytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ioutil.WriteFile(FILENAME, bytes, 0777)
}

func ReadFile() (a Account) {
	bytes, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = json.Unmarshal(bytes, &a)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return a
}

func Transfer(src, dest Envelope, amount Money) {
	src.Total -= amount
	dest.Total += amount
}
