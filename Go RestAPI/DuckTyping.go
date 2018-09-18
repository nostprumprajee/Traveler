package main

import "fmt"

type human struct {
	name string
	age  int
}

type parrot struct {
	name string
	age  int
}

type talker interface {
	talk()
}

func (h human) talk() {
	fmt.Println("Human - I'm talking.")
}

func (p parrot) talk() {
	fmt.Println("Parrot  - I'm talking.")
}

func main() {
	talkers := [2]talker{
		human{name: "Somchai", age: 23},
		parrot{name: "Dum", age: 2},
	}

	for _, talker := range talkers {
		talker.talk()
	}

	// ผลลัพธ์
	// Human - I'm talking.
	// Parrot  - I'm talking.
}
