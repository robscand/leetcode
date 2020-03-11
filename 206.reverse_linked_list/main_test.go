package main

import (
	"testing"
)

var l = NewListNode([]int{1, 2, 3, 4, 5})

func BenchmarkRevers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Revers(l)
	}
}
func BenchmarkReverseListIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseListIterative(l)
	}
}
