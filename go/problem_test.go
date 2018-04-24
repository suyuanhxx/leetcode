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
