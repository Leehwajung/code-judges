package main

func intsToLinkedList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	node := head
	for _, val := range vals[1:] {
		node.Next = &ListNode{Val: val}
		node = node.Next
	}
	return head
}

func linkedListToInts(head *ListNode) (ints []int) {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next {
		ints = append(ints, node.Val)
	}
	return ints
}
