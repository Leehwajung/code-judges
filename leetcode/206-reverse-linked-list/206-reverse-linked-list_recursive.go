//go:build recursive && !iterative

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead, ok := reverseListRecursive(head.Next, head)
	if ok {
		head.Next = nil
		return newHead
	}
	return head
}

func reverseListRecursive(curr, prev *ListNode) (*ListNode, bool) {
	if curr == nil {
		return nil, false
	}
	head, ok := reverseListRecursive(curr.Next, curr)
	curr.Next = prev
	if ok {
		return head, true
	}
	return curr, true
}
