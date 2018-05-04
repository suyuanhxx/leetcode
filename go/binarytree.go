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

func PreOrderPrint(root *TreeNode) {
	fmt.Print(root.Val, " ")
	if root.Left != nil {
		PreOrderPrint(root.Left)
	}
	if root.Right != nil {
		PreOrderPrint(root.Right)
	}
}

func InOrderPrint(root *TreeNode) {
	if root.Left != nil {
		InOrderPrint(root.Left)
	}
	fmt.Print(root.Val, " ")
	if root.Right != nil {
		InOrderPrint(root.Right)
	}
}

func PostOrderPrint(root *TreeNode) {
	if root.Left != nil {
		PostOrderPrint(root.Left)
	}
	if root.Right != nil {
		PostOrderPrint(root.Right)
	}
	fmt.Print(root.Val, " ")
}

func LevelOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		item := l.Front()
		node := item.Value.(*TreeNode)
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
