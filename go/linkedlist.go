package main

import (
	"fmt"
)

type ILinkedList interface {
	Print()
	Find(x int) *ListNode
	FindIndex(x int) *ListNode
	Insert(i int, node *ListNode) *ListNode
	InsertAfter(y int, x *ListNode)
	Delete(i int) *ListNode
	DeleteValue(x int) *ListNode
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Print() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func (head *ListNode) Find(x int) *ListNode {
	cur := head
	for cur != nil && cur.Val != x {
		cur = cur.Next
	}
	return cur
}

func (head *ListNode) FindIndex(x int) *ListNode {
	pos := 0
	cur := head
	for cur != nil && pos != x {
		cur = cur.Next
		pos++
	}
	return cur
}

func (head *ListNode) Insert(i int, node *ListNode) *ListNode {
	pos, cur := 0, head
	for cur != nil {
		if pos == i {
			temp := cur.Next
			cur.Next, node.Next = node, temp
			break
		} else {
			cur = cur.Next
		}
		pos++
	}
	return head
}

func (head *ListNode) InsertAfter(y int, x *ListNode) *ListNode {
	cur := head
	for ; cur != nil; cur = cur.Next {
		if cur.Val == y {
			temp := cur.Next
			cur.Next, x.Next = x, temp
			break
		}
	}
	return head
}

func (head *ListNode) Delete(i int) *ListNode {
	pos, cur := 0, head
	var pre *ListNode
	if pos == 0 {
		head = cur.Next
	} else {
		for cur != nil {
			if pos == i {
				pre.Next = cur.Next
				break
			}
			pre, cur = cur, cur.Next
			pos++
		}
	}
	return head
}

func (head *ListNode) DeleteValue(x int) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		if cur.Val == x {
			if cur == head {
				head = cur.Next //第一个元素
			} else {
				pre.Next = cur.Next
			}
			break
		}
		pre, cur = cur, cur.Next
	}
	return head
}

func BuildLinkList(array []int) *ListNode {
	head := &ListNode{Val: array[0]}
	cur := head
	for i := 1; i < len(array); i++ {
		temp := &ListNode{Val: array[i]}
		cur.Next = temp
		cur = cur.Next
	}
	return head
}

//func main() {
//	array := []int{3, 22, 5, 45, 23, 8}
//
//	head := BuildLinkList(array)
//
//	link := &LinkedList{Head: head}
//	fmt.Println(link.Find(100))
//	fmt.Println(link.FindIndex(9))
//	//link.Print()
//}
