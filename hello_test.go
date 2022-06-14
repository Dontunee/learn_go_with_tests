package code

import "testing"

/*
write a test
make the compiler pass
run the test, see that it fails, and check the error message is meaningful
write enough code to make the test pass
refactor
*/

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got : %q and expected : %q", got, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Tunde", "English")
		expected := "Hello,Tunde"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("empty string defaults to input 'World'", func(t *testing.T) {
		got := Hello("", "English")
		expected := "Hello,World"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Tunde", "Spanish")
		expected := "Hola,Tunde"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("says hello in french", func(t *testing.T) {
		got := Hello("Tunde", "French")
		expected := "Bonjour,Tunde"

		assertCorrectMessage(t, got, expected)
	})
}
