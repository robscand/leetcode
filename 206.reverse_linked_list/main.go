package main

import (
	"fmt"
)

func main() {
	ln := NewListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(ln)
	ln = ReverseListIterative(ln)
	PrintListNode(ln)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseOther(l *ListNode) *ListNode {
	var prev *ListNode
	for l != nil {
		l, l.Next, prev = l.Next, prev, l
	}
	return prev
}

func ReverseListIterative(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func NewListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	var head *ListNode
	node := new(ListNode)
	for i := range list {
		node.Val = list[i]
		if len(list)-1 != i {
			node.Next = new(ListNode)
		}
		if i == 0 {
			head = node
		}
		node = node.Next
	}
	return head
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node, " ")
		node = node.Next
	}
	fmt.Println()
}
