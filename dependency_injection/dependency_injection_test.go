package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "John")

	got := b.String()
	want := "Hello, John"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
