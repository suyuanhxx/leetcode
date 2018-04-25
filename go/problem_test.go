package main

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTree(preorder, inorder)
	InOrderPrint(root)
}

func TestBuildTreeFromInAndPostOrder(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTree_2(inorder, postorder)
	InOrderPrint(root)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildLinkList([]int{2, 4, 3})
	l2 := BuildLinkList([]int{5, 6, 4})
	link := &LinkedList{Head: AddTwoNumbers(l1, l2)}
	link.Print()
}
