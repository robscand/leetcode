package intersect

import "testing"

func TestIntersectSimpleFin(t *testing.T) {
	var tests = []struct{
		input1 []int
		input2 []int
		output []int
	}{
		{[]int{1,2}, []int{2,2}, []int{2}},
		{[]int{2,1}, []int{2,2}, []int{2}},
		{[]int{1,2,2,1}, []int{2,2}, []int{2,2}},
		{[]int{1,2,2,1}, []int{2}, []int{2}},
		{[]int{}, []int{1,2,3}, []int{}},
		{[]int{1,2,3}, []int{}, []int{}},
		{[]int{1,2}, []int{2,1}, []int{1,2}},
		{[]int{4,9,5}, []int{9,4,9,8,4}, []int{4,9}},
	}

	for i, test := range tests {
		res := IntersectSimple(test.input1, test.input2)
		if !equalSlices(res, test.output) {
			t.Error("Test", i, ": want", test.output, "- got", res)
		}
	}
}

func equalSlices(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
