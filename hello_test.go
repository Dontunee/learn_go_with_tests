package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got : %q and expected : %q", got, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Tunde")
		expected := "Hello,Tunde"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("empty string defaults to input 'World'", func(t *testing.T) {
		got := Hello("")
		expected := "Hello,World"

		assertCorrectMessage(t, got, expected)
	})
}
