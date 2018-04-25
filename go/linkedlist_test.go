package main

import (
	"testing"
)

func buildLinkedList() *LinkedList {
	array := []int{3, 22, 5, 45, 23, 8}
	head := BuildLinkList(array)

	return &LinkedList{Head: head}
}

func TestLinkedList_Insert(t *testing.T) {
	link := buildLinkedList()

	node := &ListNode{Val: 7}
	link.Head = link.Insert(3, node)
	link.Print()
}

func TestLinkedList_InsertAfter(t *testing.T) {
	link := buildLinkedList()
	node := &ListNode{Val: 7}

	link.Head = link.InsertAfter(3, node)
	link.Print()
}

func TestLinkedList_Delete(t *testing.T) {
	link := buildLinkedList()
	link.Head = link.Delete(3)
	link.Print()
}
func TestLinkedList_DeleteValue(t *testing.T) {
	link := buildLinkedList()
	link.Head = link.DeleteValue(3)
	link.Print()
}
