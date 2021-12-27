package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	prev := head
	head = head.Next
	prev.Next = nil
	for curr := head; head.Next != nil; prev, curr = curr, head {
		head = head.Next
		curr.Next = prev
	}
	head.Next = prev
	return head
}
