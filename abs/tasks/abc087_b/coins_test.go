package abc087b

import (
	"testing"
)

var cases = []struct {
	a, b, c, x, want int
}{
	{a: 2, b: 2, c: 2, x: 100, want: 2},
	{a: 5, b: 1, c: 0, x: 150, want: 0},
	{a: 30, b: 40, c: 50, x: 6000, want: 213},
}

func TestCoin(t *testing.T) {
	for _, tt := range cases {
		actual := coins(tt.a, tt.b, tt.c, tt.x)
		if actual != tt.want {
			t.Errorf("%d, want: %d", actual, tt.want)
		}

	}
}
