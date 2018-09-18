package main

import "fmt"

type human struct {
	name    string
	age     int
	isAdult bool
}

// ใช้ * แทนการ dereference หรือการถอดเอาค่าที่แท้จริงออกมา
func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func main() {
	somchai := human{name: "Somchai", age: 23}
	// ใช้ & แทนการอ้างถึง reference
	setAdult(&somchai)
	fmt.Println(somchai) // {Somchai 23 true}
}
