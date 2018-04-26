package main

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	size, cur := 0, head
	for cur != nil {
		size++
		cur = cur.Next
	}
	cur = head
	i, pos := 0, size-n
	if pos == 0 {
		return head.Next
	}
	var pre *ListNode
	for cur != nil {
		if i == pos {
			pre.Next = cur.Next
			break
		}
		pre, cur = cur, cur.Next
		i++
	}
	return head
}

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	var pre *ListNode
	for p != nil {
		p = p.Next
		pre = q
		q = q.Next
	}
	if q == head {
		return head.Next
	} else {
		pre.Next = q.Next
	}
	return head
}
