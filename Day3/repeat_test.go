package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	}

	t.Run("5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertEqual(t, repeated, expected)
	})

	t.Run("10 times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"
		assertEqual(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
