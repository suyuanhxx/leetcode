package main

import (
	"testing"
	"fmt"
)

func silceToInterface(s []string) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(s))
	for i, d := range s {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func createTree() *BinaryTree {
	preorder := []string{"A", "L", "B", "E", "C", "D", "W", "X"}
	inorder := []string{"B", "L", "E", "A", "C", "W", "X", "D"}
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}

	root := CreateTree(silceToInterface(preorder), silceToInterface(inorder))
	return root
}

func TestCreateTree(t *testing.T) {
	root := createTree()
	PreOrderPrint(root)
}

func TestTreeNode_GetHeight(t *testing.T) {
	root := createTree()
	fmt.Print(root.GetHeight(root))
}

func TestBinaryTree_GetNodeNumber(t *testing.T) {
	root := createTree()
	fmt.Print(root.GetNodeNumber(root))
}
