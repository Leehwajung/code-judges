package main

func linkedListWithCycle(pos int, vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var cycleNode *ListNode
	head := &ListNode{Val: vals[0]}
	node := head
	if pos == 0 {
		cycleNode = node
	}

	pos--
	for idx, val := range vals[1:] {
		node.Next = &ListNode{Val: val}
		node = node.Next
		if idx == pos {
			cycleNode = node
		}
	}

	node.Next = cycleNode
	return head
}
