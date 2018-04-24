package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderPrint(root *TreeNode) {
	fmt.Print(root.Val, " ")
	if root.Left != nil {
		InOrderPrint(root.Left)
	}
	if root.Right != nil {
		InOrderPrint(root.Right)
	}
}
