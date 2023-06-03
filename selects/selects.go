package selects

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("httpclient timeout")

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

type indexDuration struct {
	index    int
	duration time.Duration
}

func Racer(a, b string, timeout time.Duration) (string, error) {
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racerr(urls ...string) string {
	if len(urls) == 0 {
		return ""
	}

	times := make([]time.Duration, len(urls))
	timesChannel := make(chan indexDuration)
	for i, url := range urls {
		go func(ind int, u string) {
			timesChannel <- indexDuration{ind, measureResponseTime(u)}
		}(i, url)
	}

	for i := 0; i < len(times); i++ {
		r := <-timesChannel
		times[r.index] = r.duration
	}

	var winnerIndex int
	for i := 1; i < len(times); i++ {
		if times[winnerIndex] > times[i] {
			winnerIndex = i
		}
	}
	return urls[winnerIndex]
}
