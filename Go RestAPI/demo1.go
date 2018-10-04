package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start main")
	name := "Fake name"
	go hello(name)
	fmt.Println("Finish main")

	time.Sleep(time.Second)
}

func hello(name string) string {
	result := "Hello " + name
	fmt.Printf("In function = %s\n", result)
	return result
}
