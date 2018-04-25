package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type ILinkedList interface {
	Print()
	Find(x int) *ListNode
	FindIndex(x int) *ListNode
}

type LinkedList struct {
	Head    *ListNode
	Current *ListNode
	Size    int
}

func (linkedList *LinkedList) Print() {
	cur := linkedList.Head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func (linkedList *LinkedList) Find(x int) *ListNode {
	cur := linkedList.Head
	for cur != nil && cur.Val != x {
		cur = cur.Next
	}
	return cur
}

func (linkedList *LinkedList) FindIndex(x int) *ListNode {
	pos := 0
	cur := linkedList.Head
	for cur != nil && pos != x {
		cur = cur.Next
		pos++
	}
	return cur
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

func main() {
	array := []int{3, 22, 5, 45, 23, 8}

	head := BuildLinkList(array)

	link := &LinkedList{Head: head}
	fmt.Println(link.Find(100))
	fmt.Println(link.FindIndex(9))
	//link.Print()
}
