package main

import (
	"fmt"
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Val   interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func PreOrderPrint(root *BinaryTree) {
	fmt.Print(root.Val, " ")
	if root.Left != nil {
		PreOrderPrint(root.Left)
	}
	if root.Right != nil {
		PreOrderPrint(root.Right)
	}
}

func InOrderPrint(root *BinaryTree) {
	if root.Left != nil {
		InOrderPrint(root.Left)
	}
	fmt.Print(root.Val, " ")
	if root.Right != nil {
		InOrderPrint(root.Right)
	}
}

func PostOrderPrint(root *BinaryTree) {
	if root.Left != nil {
		PostOrderPrint(root.Left)
	}
	if root.Right != nil {
		PostOrderPrint(root.Right)
	}
	fmt.Print(root.Val, " ")
}

func LevelOrderPrint(root *BinaryTree) {
	if root == nil {
		return
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		item := l.Front()
		node := item.Value.(*BinaryTree)
		fmt.Print(node.Val, " ")
		l.Remove(item)
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
}

func (root *BinaryTree) GetNodeNumber(node *BinaryTree) int {
	if node == nil {
		return 0
	}
	return root.GetNodeNumber(node.Left) + root.GetNodeNumber(node.Right) + 1;
}

func (root *BinaryTree) GetHeight(node *BinaryTree) int {
	if node == nil {
		return 0
	} else {
		return maxOfTwo(root.GetHeight(node.Left), root.GetHeight(node.Right)) + 1
	}
}

func maxOfTwo(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func CreateTree(preorder []interface{}, inorder []interface{}) *BinaryTree {
	return buildBinaryTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildBinaryTree(preorder []interface{}, inorder []interface{},
	preStart, preEnd, inStart, inEnd int) *BinaryTree {
	if preStart > preEnd || inStart > inEnd {
		return nil;
	}

	root := &BinaryTree{Val: preorder[preStart]}
	var flag int
	for i := inStart; i <= inEnd; i++ {
		if preorder[preStart] == inorder[i] {
			flag = i
			break
		}
	}
	root.Left = buildBinaryTree(preorder, inorder, preStart+1, preStart+flag-inStart, inStart, flag-1);
	root.Right = buildBinaryTree(preorder, inorder, preStart+flag-inStart+1, preEnd, flag+1, inEnd);
	return root
}
