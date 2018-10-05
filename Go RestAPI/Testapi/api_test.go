package main

import (
	"testing"
)

type CounterApiFromMock struct{}

func (c CounterApiFromMock) count() (int, error) {
	return 2, nil
}

func TestResultShouldBe20_WhenCounterIs_2(t *testing.T) {
	result := myProcess(CounterApiFromMock{})
	if result != 20 {
		t.Fatalf("Expected 20 but got %d", result)
	}
}
