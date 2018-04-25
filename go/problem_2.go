package main

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur1, cur2, carry := l1, l2, 0
	var head, cur *ListNode
	for cur1 != nil || cur2 != nil || carry > 0 {
		result := carry
		if cur1 != nil {
			result += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			result += cur2.Val
			cur2 = cur2.Next
		}
		result, carry = Add(result)
		if head == nil {
			head = &ListNode{Val: result}
			cur = head
		} else {
			temp := &ListNode{Val: result}
			cur.Next = temp
			cur = cur.Next
		}
	}
	return head
}

func Add(a int) (result, carry int) {
	if a >= 10 {
		carry = a / 10
		result = a % 10
		return result, carry
	}
	return a, 0
}
