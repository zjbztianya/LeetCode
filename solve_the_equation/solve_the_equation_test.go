package solve_the_equation

import "testing"

func TestSolveEquation(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"x+5-3+x=6+x-2", "x=2"},
		{"x=x", "Infinite solutions"},
		{"2x=x", "x=0"},
		{"2x+3x-6x=x+2", "x=-1"},
		{"x=x+2", "No solution"},
		{"1-x+x-x+x=-99-x+x-x+x", "No solution"},
		{"0x=0", "Infinite solutions"},
	}
	for i := range tests {
		if got := solveEquation(tests[i].input); got != tests[i].want {
			t.Errorf("solveEquation(%v)=%v", tests[i].input, got)
		}
	}
}
