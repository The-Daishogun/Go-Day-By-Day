package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("The_Daishogun")
	want := "Hello, The_Daishogun"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
