package java;

import java.util.LinkedList;
import java.util.Queue;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class BinaryTreeProblems {

    // LeetCodeOJ 105前序遍历和中序遍历构造二叉树
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return helper(preorder,0,preorder.length-1,inorder,0,inorder.length-1);
    }
    public TreeNode helper(int[] preorder,int i,int j,int[] inorder,int k,int v){
        if(i>j)return null;
        if(i==j)return new TreeNode(preorder[i]);
        int index=k;
        for(;index<=v;index++){
            if(preorder[i]==inorder[index])
                break;
        }
        TreeNode root = new TreeNode(preorder[i]);
        root.left = helper(preorder,i+1,index-k+i,inorder,k,index-1);
        root.right = helper(preorder,index-k+i+1,j,inorder,index+1,v);
        return root;
    }


    //二叉树倒置
    public TreeNode invertTree(TreeNode root) {
        if(null==root)return null;
        TreeNode result = new TreeNode(root.val);
        if(null!=root.left && null!=root.right){
            result.left=invertTree(root.right);
            result.right=invertTree(root.left);
        }
        if(null!=root.left && null==root.right){
            result.right=invertTree(root.left);
        }
        if(null==root.left && null!=root.right){
            result.left=invertTree(root.right);
        }
        return result;
    }

    //二叉树层次遍历
    public void printTreeNode(TreeNode root){
        if(null==root)return;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        while (!q.isEmpty()){
            TreeNode temp = q.remove();
            System.out.println(temp.val);
            if(null!=temp.left && null!=temp.right){
                q.add(temp.left);
                q.add(temp.right);
            }
            if(null!=temp.left && null==temp.right){
                q.add(root.left);
            }
            if(null==temp.left && null!=temp.right){
                q.add(temp.right);
            }
        }
    }

    public void preOrder(TreeNode root){
        if(null==root)return;
        System.out.print(root.val+" ");
        if(null!=root.left)
            preOrder(root.left);
        if(null!=root.right)
            preOrder(root.right);
    }
    public void inOrder(TreeNode root){
        if(null==root)return;
        if(null!=root.left)
            inOrder(root.left);
        System.out.print(root.val+" ");
        if(null!=root.right)
            inOrder(root.right);
    }
    public void postOrder(TreeNode root){
        if(null==root)return;
        if(null!=root.left)
            postOrder(root.left);
        if(null!=root.right)
            postOrder(root.right);
        System.out.print(root.val+" ");
    }

}
