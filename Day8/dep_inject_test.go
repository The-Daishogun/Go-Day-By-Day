package depinject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "The_Daishogun")

	got := buffer.String()
	want := "Hello, The_Daishogun"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
