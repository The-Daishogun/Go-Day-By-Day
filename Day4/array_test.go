package array

import "testing"

func TestSum(t *testing.T) {

	assertEqual := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}

	t.Run("len 5 array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertEqual(t, got, want, numbers)
	})

}
