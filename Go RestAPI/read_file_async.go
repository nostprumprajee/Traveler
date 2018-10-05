package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

func readFile(filename string) {
	defer wg.Done()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}

func main() {
	wg.Add(1)
	go readFile("data.txt")
	fmt.Println("Finish !!")
	wg.Wait()
}
