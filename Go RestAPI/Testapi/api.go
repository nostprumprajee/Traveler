package main

import (
	"fmt"
)

type Counter interface {
	count() (int, error)
}

type CounterApiFromNetwork struct{}

func (c CounterApiFromNetwork) count() (int, error) {
	// Call data from API
	return -1, fmt.Errorf("TODO NEXT")
}

func myProcess(counter Counter) int {
	result, error := counter.count()
	if error == nil {
		return result * 10
	}
	return 0
}

func main() {
	result := myProcess(CounterApiFromNetwork{})
	fmt.Printf("Result = %d", result)
}
