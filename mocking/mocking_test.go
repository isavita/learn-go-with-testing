package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		b := &bytes.Buffer{}
		spy := &SpySleeper{}
		Countdown(b, spy)

		got := b.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spy.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spy.Calls)
		}
	})

	t.Run("sleeps before every print", func(t *testing.T) {
		spy := &SpyCountdownOperations{}
		Countdown(spy, spy)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spy.Calls, want) {
			t.Errorf("got %v wanted calls %v", spy.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func BenchmarkCountdown(b *testing.B) {
	buff := &bytes.Buffer{}
	spy := &SpySleeper{}
	for i := 0; i < b.N; i++ {
		Countdown(buff, spy)
	}
}
