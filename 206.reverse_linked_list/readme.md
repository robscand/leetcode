Reverse a singly linked list.

URL: https://leetcode.com/problems/reverse-linked-list/

Example:

    Input: 1->2->3->4->5->NULL
    Output: 5->4->3->2->1->NULL

Follow up:

    A linked list can be reversed either iteratively or recursively. Could you implement both?

Bench:

    goos: darwin
    goarch: amd64
    pkg: reverse-linked-list
    BenchmarkRevers2-4                      307097828                3.82 ns/op            0 B/op          0 allocs/op
    BenchmarkReverseListIterative-4         402748599                2.93 ns/op            0 B/op          0 allocs/op
    PASS
    ok      reverse-linked-list     3.107s
