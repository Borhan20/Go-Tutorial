package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		output := message("shovon", "")
		want := "Hello shovon"
		assertCorrectMessage(t, output, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		output := message("", "")
		want := "Hello world"
		assertCorrectMessage(t, output, want)
	})

	t.Run("saying hello in Bangla", func(t *testing.T) {
		output := message("নওফা", "Bangla")
		want := "আসসালামুয়ালাইকুম নওফা"

		if output != want {
			t.Errorf("output is %q, want %q", output, want)
		}
	})
}

// Add the missing function `assertCorrectMessage`
func assertCorrectMessage(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
