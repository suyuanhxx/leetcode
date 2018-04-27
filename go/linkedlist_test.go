package main

import (
	"testing"
)

func buildLinkedList() *ListNode {
	array := []int{3, 22, 5, 45, 23, 8}
	head := BuildLinkList(array)

	return head
}

func buildDoubleLinkList() *DoubleListNode {
	array := []int{3, 22, 5, 45, 23, 8}
	head := BuildDoubleLinkList(array)
	return head
}

func TestLinkedList_Insert(t *testing.T) {
	link := buildLinkedList()

	node := &ListNode{Val: 7}
	link = link.Insert(3, node)
	link.Print()
}

func TestLinkedList_InsertAfter(t *testing.T) {
	link := buildLinkedList()
	node := &ListNode{Val: 7}

	link = link.InsertAfter(3, node)
	link.Print()
}

func TestLinkedList_Delete(t *testing.T) {
	link := buildLinkedList()
	link = link.Delete(0)
	link.Print()
}
func TestLinkedList_DeleteValue(t *testing.T) {
	link := buildLinkedList()
	link = link.DeleteValue(3)
	link.Print()
}

func TestBuildDoubleLinkList(t *testing.T) {
	link := buildDoubleLinkList()
	link.Print()
}

func TestDoubleListNode_Insert(t *testing.T) {
	link := buildDoubleLinkList()
	link.Insert(22, 15)
	link.Print()
}
