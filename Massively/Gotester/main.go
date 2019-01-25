package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// DefaultSleeper default
type DefaultSleeper struct{}

// Sleep logic
func (r *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleeper should real funcftion call only
type Sleeper interface {
	Sleep()
}

// Countdown is function loopback count
func Countdown(word io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(word, i)
	}
	s.Sleep()
	fmt.Fprint(word, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
