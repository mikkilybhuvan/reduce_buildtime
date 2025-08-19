package compute

import "testing"

func TestFib(t *testing.T) {
	cases := []struct{ in, want int }{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {10, 55},
	}
	for _, c := range cases {
		if got := Fib(c.in); got != c.want {
			t.Fatalf("Fib(%d)=%d, want %d", c.in, got, c.want)
		}
	}
}
