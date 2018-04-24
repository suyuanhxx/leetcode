package main

/**
leetcode problem 105: construct the binary tree from preorder and inorder list
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]

Return the following binary tree:
    3
   / \
  9  20
    /  \
   15   7
 */

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder []int, inorder []int, preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil;
	}

	root := &TreeNode{Val: preorder[preStart]}
	var flag int
	for i := inStart; i <= inEnd; i++ {
		if preorder[preStart] == inorder[i] {
			flag = i
			break
		}
	}
	root.Left = build(preorder, inorder, preStart+1, preStart+flag-inStart, inStart, flag-1);
	root.Right = build(preorder, inorder, preStart+flag-inStart+1, preEnd, flag+1, inEnd);
	return root
}
