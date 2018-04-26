package main

import (
	"fmt"
)

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

func (head *DoubleListNode) Print() {
	cur := head
	var prev *DoubleListNode
	for cur != nil && prev != head {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
		prev = cur
	}
}

func (head *DoubleListNode) Insert(y int, x int) *DoubleListNode {
	cur := head
	var prev *DoubleListNode
	for cur != nil && prev != head {
		if cur.Val == y {
			tmp := &DoubleListNode{Val: x, Next: cur.Next, Prev: cur}
			cur.Next.Prev = tmp
			cur.Next = tmp
			break
		}
		cur = cur.Next
		prev = cur
	}
	return head
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
