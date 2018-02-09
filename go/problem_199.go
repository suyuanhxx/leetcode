package main

import "fmt"
import (
	. "./tree"
	"container/list"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	array := levelOrderEnd(root)
	for i, item := range array {
		if item == nil {
			result = append(result, *array[i-1])
		}
	}
	return result
}

func levelOrderEnd(root *TreeNode) []*int {
	l := list.New()
	l.PushBack(root)
	l.PushBack(nil)
	var result []*int
	for ; l.Len() > 0; {
		a := l.Front()
		l.Remove(a)
		value := a.Value
		if value != nil {
			node := value.(*TreeNode)
			result = append(result, &node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		} else if l.Len() > 0 {
			l.PushBack(nil)
			result = append(result, nil)
		}
	}
	result = append(result, nil)
	return result
}

func main() {
	left := TreeNode{1, nil, nil}
	right := TreeNode{5, nil, nil}
	left1 := TreeNode{-2, nil, nil}
	right1 := TreeNode{4, &left, &right}
	root := TreeNode{1, &left1, &right1}

	array := rightSideView(&root)
	for _, num := range array {
		fmt.Println(num)
	}
}
