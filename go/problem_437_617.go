package main

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	pathSum(root, sum)
	count := 0
	num := root.Val
	//cur := &root
	if num == sum {
		count++
	}
	if root.Left != nil {
		root = root.Left
		num += root.Val
	}
	if root.Right != nil {
		root = root.Right
		num += root.Val
	}
	return count
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

//func main() {
//	left := TreeNode{1, nil, nil}
//	right := TreeNode{5, nil, nil}
//	left1 := TreeNode{-2, nil, nil}
//	right1 := TreeNode{4, &left, &right}
//	root := TreeNode{1, &left1, &right1}
//	//fmt.Println(pathSum(&root, 6))
//
//	left2 := TreeNode{-2, &left, &right}
//	root2 := TreeNode{1, &left2, &right1}
//
//	fmt.Println(mergeTrees(&root, &root2))
//
//}
