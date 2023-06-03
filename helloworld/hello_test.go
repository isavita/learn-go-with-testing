package helloworld

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World' and English", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'Welt' when German", func(t *testing.T) {
		got := Hello("", "de")
		want := "Hallo, Welt"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'monde' when French", func(t *testing.T) {
		got := Hello("", "fr")
		want := "Bonjour, monde"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hallo in German", func(t *testing.T) {
		got := Hello("Scholz", "de")
		want := "Hallo, Scholz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying bonjour in French", func(t *testing.T) {
		got := Hello("Nicolas", "fr")
		want := "Bonjour, Nicolas"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Hello, %s", "John")
	}
}

func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello, " + "John"
	}
}
