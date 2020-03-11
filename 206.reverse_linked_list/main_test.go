package main

import (
	"testing"
)

var bl = NewListNode([]int{1, 2, 3, 4, 5})

func TestReverseListIterative1(t *testing.T) {
	var res = []int{5, 4, 3, 2, 1}
	l := NewListNode([]int{1, 2, 3, 4, 5})
	l = ReverseListIterative(l)
	for i := range res {
		if res[i] != l.Val {
			t.Errorf("Expect %d, but got %d", res[i], l.Val)
		}
		l = l.Next
	}
}

func BenchmarkReversOther(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseOther(bl)
	}
}

func BenchmarkReverseListIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseListIterative(bl)
	}
}
