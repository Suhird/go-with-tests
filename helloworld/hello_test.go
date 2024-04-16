package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("spanish greeting", func(t *testing.T) {
		got := Hello("Suhird", "Spanish")
		want := "Hola, Suhird"
		assertCorrectMessage(t, got, want)
	})
	t.Run("french greeting", func(t *testing.T) {
		got := Hello("Suhird", "French")
		want := "Bonjour, Suhird"
		assertCorrectMessage(t, got, want)
	})
	t.Run("english greeting", func(t *testing.T) {
		got := Hello("Suhird", "")
		want := "Hello, Suhird"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
