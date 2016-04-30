package com.demo;

/**
 * Created by xxhuang on 2016/4/28.
 */

import java.util.LinkedList;
import java.util.Queue;

/**
 * Definition for a binary tree node.
 * */


public class Solution {


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

    public void moveZeroes(int[] nums) {
        int len = nums.length;
        int i=0;
        if(0==len)return;
        while(i<len){
            if(0==nums[i]){
                int j=i;
                for(;j+1<=len-1;j++)
                {
                    nums[j] = nums[j+1];
                }
                nums[j]=0;
                len--;
            }
            else{
                i++;
            }
        }
    }

    public void deleteNode(ListNode node) {
        if(null==node)return;
        ListNode next = node.next;
        while(null!=next)
        {
            node.val = next.val;
            node = next;
            next = next.next;
        }
    }

    public void deleteNode1(ListNode node) {
        if(null==node)return;
        if(null==node.next){
            node=null;
            return;
        }
        else{
            ListNode next = node.next;
            node.val = next.val;
            node.next = next;
            deleteNode(next);
        }
        node = null;
    }


    //字符串反转
    public String reversString(String s){
        //StringBuilder str=new StringBuilder(s);
        //return str.reverse().toString();
        char[] str = s.toCharArray();
        int len = s.length()-1;
        for(int i=(len-1)>>1;i>=0;i--){
            int j = len-i;
            char a = str[i];
            char b = str[j];
            str[j] = a;
            str[i] = b;
        }
        return String.valueOf(str);
    }

    public static void main(String[]args){
        Solution s = new Solution();
        TreeNode root = new TreeNode(4);
        root.left = new TreeNode(7);
        root.right = new TreeNode(9);
        root.left.left = new TreeNode(12);
        root.left.right = new TreeNode(45);
        //TreeNode result = s.invertTree(root);
        //s.printTreeNode(root);
        ListNode ls = new ListNode(0);
        ListNode ls1 = new ListNode(1);
        ListNode ls2 = new ListNode(6);
        ListNode ls3 = new ListNode(7);
        ls.next = ls1;
        ls1.next = ls2;
        ls2.next = ls3;
        s.deleteNode1(ls1);
    }


}
