package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("The_Daishogun", "")
		want := "Hello, The_Daishogun"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello, world! when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jeff", "French")
		want := "Bonjour, Jeff"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Persian", func(t *testing.T) {
		got := Hello("جواد", "Persian")
		want := "سلام، جواد"
		assertCorrectMessage(t, got, want)
	})
}
