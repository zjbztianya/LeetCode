package valid_sudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	var tests = []struct {
		input [][]byte
		want  bool
	}{
		{[][]byte{
			[]byte("53..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		}, true,
		},
	}

	for i, test := range tests {
		if got := isValidSudoku(test.input); got != test.want {
			t.Errorf("test case %d: output: %v want: %v", i, got, test.want)
		}
	}
}
