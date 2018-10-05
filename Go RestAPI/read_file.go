package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}

func main() {
	readFile("data.txt")
	fmt.Println("Finish !!")
}
