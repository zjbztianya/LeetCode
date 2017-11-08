package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		if got := trap(test.input); got != test.output {
			t.Errorf("trap(%v)=%d want:%d", test.input, got, test.output)
		}
	}
}
