package main

import "fmt"

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

func (head *DoubleListNode) Print() {
	fmt.Print(head.Val, " ")
	cur := head.Next
	for cur != head {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func BuildDoubleLinkList(array []int) *DoubleListNode {
	head := &DoubleListNode{Val: array[0]}
	cur := head
	for i := 1; i < len(array); i++ {
		temp := &DoubleListNode{Val: array[i]}
		temp.Prev = cur
		cur.Next = temp
		cur = cur.Next
	}
	cur.Next = head
	head.Prev = cur
	return head
}
