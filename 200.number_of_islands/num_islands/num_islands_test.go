package num_islands

import (
	"testing"
)

func TestNumIslandsChange(t *testing.T) {
	var tests = []struct {
		grid   [][]byte
		result int
	}{
		{
			grid:   [][]byte{{'1'}},
			result: 1,
		}, {
			grid:   [][]byte{{'0'}},
			result: 0,
		}, {
			grid:   [][]byte{{'0', '1'}},
			result: 1,
		}, {
			grid:   [][]byte{{'0'}, {'1'}},
			result: 1,
		}, {
			grid:   [][]byte{{}},
			result: 0,
		}, {
			grid:   [][]byte{{'1', '1', '1', '1', '0'}},
			result: 1,
		}, {
			grid:   [][]byte{{'1', '1', '0', '1', '0'}},
			result: 2,
		}, {
			grid: [][]byte{{'1'},
				{'1'},
				{'1'},
				{'0'}},
			result: 1,
		}, {
			grid: [][]byte{{'0'},
				{'1'},
				{'0'},
				{'1'},
				{'0'}},
			result: 2,
		}, {
			grid: [][]byte{{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'}},
			result: 1,
		}, {
			grid: [][]byte{{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'}},
			result: 3,
		}, {
			grid: [][]byte{{'1', '1', '0', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'0', '0', '0', '0', '0'},
				{'1', '1', '1', '1', '1'},
				{'0', '1', '0', '0', '0'}},
			result: 2,
		}, {
			grid: [][]byte{{'0', '0', '0', '1', '0'},
				{'0', '0', '1', '1', '1'},
				{'0', '1', '0', '0', '1'},
				{'0', '0', '0', '1', '1'}},
			result: 2,
		}, {
			grid: [][]byte{{'0', '0', '0', '1', '0'},
				{'0', '1', '1', '1', '1'},
				{'0', '1', '0', '0', '1'},
				{'0', '0', '0', '0', '1'},
				{'0', '0', '0', '1', '1'}},
			result: 1,
		},
	}

	for _, test := range tests {
		if got := NumIslands(test.grid); got != test.result {
			t.Errorf("Expect NumIslandsChange(grid) = %d, but in result = %d", test.result, got)
		}
	}
}

func BenchmarkNumIslands(b *testing.B) {
	grid := [][]byte{{'0', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1'},
		{'0', '1', '0', '0', '1'},
		{'0', '0', '0', '0', '1'},
		{'0', '0', '0', '1', '1'}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NumIslands(grid)
	}
}
