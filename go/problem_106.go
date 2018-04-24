package main

/**
leetcode problem 105: construct the binary tree from inorder and postorder list
inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]

Return the following binary tree:
    3
   / \
  9  20
    /  \
   15   7
 */

func BuildTree_2(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 {
		return nil;
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	postorder = postorder[:len(postorder)-1]
	var leftInorder, leftPostorder, rightInorder, rightPostorder []int
	for index, value := range inorder {
		if value == root.Val {
			leftPostorder = postorder[:index]
			rightPostorder = postorder[index:]
			leftInorder = inorder[:index]
			rightInorder = inorder[index+1:]
			break
		}
	}
	root.Right = BuildTree_2(rightInorder, rightPostorder);
	root.Left = BuildTree_2(leftInorder, leftPostorder);
	return root
}
