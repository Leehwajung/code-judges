package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	} else if head.Next == head {
		return true
	}

	forward := head.Next
	for ; head != nil; head = head.Next {
		if forward == nil {
			return false
		}
		if head == forward {
			return true
		}
		if forward.Next == nil {
			return false
		}
		forward = forward.Next.Next
	}
	return false
}
