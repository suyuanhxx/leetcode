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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := TreeNode{Val: preorder[0]}
	return build(preorder, inorder, &root)
}

func build(preorder []int, inorder []int, root *TreeNode) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return root
	}
	for _, data := range preorder {
		for j, k := range inorder {
			if data == k {

			}
		}
	}
}

func main() {

}
