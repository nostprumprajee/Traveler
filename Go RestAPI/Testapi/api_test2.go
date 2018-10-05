package main

import (
	"fmt"
	"testing"
)

type CounterErrorApiFromMock struct{}

func (c CounterErrorApiFromMock) count() (int, error) {
	return 0, fmt.Errorf("Have error")
}

func TestResultShouldBe0_WhenCounterHaveError(t *testing.T) {
	result := myProcess(CounterErrorApiFromMock{})
	if result != 0 {
		t.Fatalf("Expected 20 but got %d", result)
	}
}
