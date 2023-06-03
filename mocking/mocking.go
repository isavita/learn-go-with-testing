package mocking

import (
	"fmt"
	"io"
	"time"
)

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleepp   func(time.Duration)
}

func (s ConfigurableSleeper) Sleep() {
	s.Sleepp(s.Duration)
}

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprintln(w, "Go!")
}
