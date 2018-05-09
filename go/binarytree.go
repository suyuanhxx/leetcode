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

func (binaryTree *BinaryTree) CreateHuffman(element []int) *BinaryTree {
	forest := make([]*BinaryTree, len(element))
	for i, item := range element {
		node := &BinaryTree{Val: item}
		forest[i] = node
	}
	var result *BinaryTree
	for i := 1; i < len(element); i++ {
		var k1 = -1
		var k2 int // k1表示森林中具有最小权值的树根结点的下标，k2为次最小的下标
		for j := 0; j < len(element); j++ { // 让k1初始指向森林中第一棵树，k2指向第二棵
			if forest[j] != nil && k1 == -1 {
				k1 = j
				continue
			}
			if forest[j] != nil {
				k2 = j
				break
			}
		}
		for j := k2; j < len(element); j++ { // 从当前森林中求出最小权值树和次最小
			if forest[j] != nil {
				if forest[j].Val.(int) < forest[k1].Val.(int) {
					k2, k1 = k1, j
				} else if forest[j].Val.(int) < forest[k2].Val.(int) {
					k2 = j
				}
			}
		}
		result = &BinaryTree{Val: forest[k1].Val.(int) + forest[k2].Val.(int),
			Left: forest[k1], Right: forest[k2]}

		forest[k1] = result
		forest[k2] = nil
	}
	return result
}

var a [10]int

func (binaryTree *BinaryTree) HuffManCoding(node *BinaryTree, len int) {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			fmt.Printf("结点权值为%d的编码：", node.Val);
			for i := 0; i < len; i++ {
				fmt.Print(a[i])
			}
			fmt.Println()
		} else {
			a[len] = 0
			binaryTree.HuffManCoding(node.Left, len+1)
			a[len] = 1
			binaryTree.HuffManCoding(node.Right, len+1)
		}
	}
}
