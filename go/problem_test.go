package main

import (
	"testing"
	"fmt"
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
	l1 := BuildLinkList([]int{2, 4, 8, 7, 4, 7})
	l2 := BuildLinkList([]int{5, 6, 4, 6})
	link := &LinkedList{Head: AddTwoNumbers(l1, l2)}
	link.Print()
}

func TestAdd(t *testing.T) {
	array := []int{3, 22, 5, 45, 23, 8}

	head := BuildLinkList(array)

	link := &LinkedList{Head: head}
	fmt.Println(link.Find(100))
	fmt.Println(link.FindIndex(3))
	//link.Print()

}
