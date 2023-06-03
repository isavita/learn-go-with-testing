package main

import (
	"learn_go_with_testing/mocking"
	"os"
	"time"
)

func main() {
	sleeper := mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
