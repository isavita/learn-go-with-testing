package concurrency

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}

func slowStubWebsiteChecker(string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{"https://www.githubstatus.com", "http://doesnotexisttt.com", "api.com"}
	got := CheckWebsites(mockWebsiteChecker, urls)
	want := map[string]bool{
		"https://www.githubstatus.com": true,
		"http://doesnotexisttt.com":    true,
		"api.com":                      false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "someurl.com"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
