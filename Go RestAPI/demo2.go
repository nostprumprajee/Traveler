package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Start main")
	// name := "Fake name"
	fmt.Scanf("%s", &name)
	result := make(chan string)
	go hello(name, result)

	fmt.Println("Finish main")

	fmt.Println(<-result)
}

func hello(name string, result chan<- string) {
	output := "Hello " + name
	fmt.Printf("In function = %s\n", output)
	result <- output
}
