package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	s := Repeat("abc", 2)
	fmt.Println(s)
	// Output: abcabc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkStandardRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 10)
	}
}
