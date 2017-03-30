package com.freedom.leetcode.tree;

import org.junit.Test;

import java.util.LinkedList;
import java.util.Queue;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class BinaryTreeProblems {

    /**
     * LeetCodeOJ 105前序遍历和中序遍历构造二叉树
     *
     * @param preOrder
     * @param inOrder
     * @return
     */
    public TreeNode buildTree(int[] preOrder, int[] inOrder) {
        return helper(preOrder, 0, preOrder.length - 1, inOrder, 0, inOrder.length - 1);
    }

    public TreeNode helper(int[] preOrder, int i, int j, int[] inOrder, int k, int v) {
        if (i > j) return null;
        if (i == j) return new TreeNode(preOrder[i]);
        int index = k;
        for (; index <= v; index++) {
            if (preOrder[i] == inOrder[index])
                break;
        }
        TreeNode root = new TreeNode(preOrder[i]);
        root.left = helper(preOrder, i + 1, index - k + i, inOrder, k, index - 1);
        root.right = helper(preOrder, index - k + i + 1, j, inOrder, index + 1, v);
        return root;
    }


    /**
     * 二叉树倒置
     *
     * @param root
     * @return
     */
    public TreeNode invertTree(TreeNode root) {
        if (null == root) return null;
        TreeNode result = new TreeNode(root.val);
        if (null != root.left && null != root.right) {
            result.left = invertTree(root.right);
            result.right = invertTree(root.left);
        }
        if (null != root.left && null == root.right) {
            result.right = invertTree(root.left);
        }
        if (null == root.left && null != root.right) {
            result.left = invertTree(root.right);
        }
        return result;
    }

    /**
     * 二叉树层次遍历
     *
     * @param root
     */
    public void levelOrder(TreeNode root) {
        if (null == root) return;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        while (!q.isEmpty()) {
            TreeNode temp = q.remove();
            System.out.println(temp.val);
            if (null != temp.left && null != temp.right) {
                q.add(temp.left);
                q.add(temp.right);
            }
            if (null != temp.left && null == temp.right) {
                q.add(root.left);
            }
            if (null == temp.left && null != temp.right) {
                q.add(temp.right);
            }
        }
    }

    /**
     * 前序遍历
     *
     * @param root
     */
    public void preOrder(TreeNode root) {
        if (null == root) return;
        System.out.print(root.val + " ");
        if (null != root.left)
            preOrder(root.left);
        if (null != root.right)
            preOrder(root.right);
    }

    /**
     * 中序遍历
     *
     * @param root
     */
    public void inOrder(TreeNode root) {
        if (null == root) return;
        if (null != root.left)
            inOrder(root.left);
        System.out.print(root.val + " ");
        if (null != root.right)
            inOrder(root.right);
    }

    /**
     * 后序遍历
     *
     * @param root
     */
    public void postOrder(TreeNode root) {
        if (null == root) return;
        if (null != root.left)
            postOrder(root.left);
        if (null != root.right)
            postOrder(root.right);
        System.out.print(root.val + " ");
    }


    @Test
    public void testInteger() {
        Integer a = 127;
        Integer b = 127;
        System.out.print(a == b);
    }

}
