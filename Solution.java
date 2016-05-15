package com.leetcode;
/**
 * Created by xxhuang on 2016/4/28.
 */
import java.util.*;

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

    //单链表反转LeetCode OJ 206
    public ListNode reverseList(ListNode head){
        if(null==head)return null;
        ListNode result =null;
        ListNode p=null;
        ListNode q=head;
        while(null!=head){
            q = head.next;
            result =new ListNode(head.val);
            result.next = p;
            p=result;
            head = q;
        }
        return result;
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

    public boolean isPowOfTwo(int n){
        //if(0==n)return true;
        //int count=0;
        //while (n>0){
        //    count+=n&0x01;
        //    n=n>>1;
        //}
        //return count==1;

        return n>0 && ((n&(n-1))==0);

    }

    public boolean isPowOfThree(int n){
        if(0==n)return true;
        boolean p=false;
        int sum=1;
        for (int i=0;i<n;i++){
            sum*=3;
            if(n==sum) {
                p = true;
                break;
            }
        }
        return p;
    }

    //前序遍历和中序遍历构造二叉树 LeetCodeOJ 105
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

    //LeetCode 242 Valid Anagram
    public boolean isAnagram(String s, String t) {
        if(s.length()!=t.length())return false;
        else{
            Map<Character,Integer> s1=new HashMap<Character,Integer>();
            for(int i=0;i<s.length();i++){
                if(s1.containsKey(s.charAt(i))){
                    s1.put(s.charAt(i),s1.get(s.charAt(i))+1);
                }
                else{
                    s1.put(s.charAt(i),1);
                }
            }
            Map<Character,Integer> s2=new HashMap<Character,Integer>();
            for(int i=0;i<t.length();i++){
                if(s2.containsKey(t.charAt(i))){
                    s2.put(t.charAt(i),s2.get(t.charAt(i))+1);
                }
                else{
                    s2.put(t.charAt(i),1);
                }
            }
            if(s1.size()!=s2.size())return false;
            for(Character k:s1.keySet()){
                int v1 = s1.get(k);
                if(!s2.containsKey(k))return false;
                int v2 = s2.get(k);
                if(v1!=v2){
                    return false;
                }
            }
            return true;
        }
    }

    public boolean isAnagram2(String s, String t) {
        if(s.length()!=t.length())return false;
        char[] s1 = s.toCharArray();
        char[] s2 = t.toCharArray();
        Arrays.sort(s1);
        Arrays.sort(s2);
        for(int i=0;i<s.length();i++){
            if(s1[i]!=s2[i])return false;
        }
        return true;
    }

    //171. Excel Sheet Column Number
    public int titleToNumber(String s) {
        int num = 0;
        int len = s.length();
        for(int i=len-1;i>=0;i--){
            int c = (int)s.charAt(i)-64;
            if(len-1-i>0) num += c*Math.pow(26,len-1-i);
            else num+=c;
        }
        return num;
    }

    //217. Contains Duplicate
    public boolean containsDuplicate(int[] nums) {
        Set set = new HashSet();
        for(int i=0;i<nums.length;i++)
            set.add(nums[i]);
        if(set.size()==nums.length)return false;
        else return true;
    }

    //169. Majority Element
    public int majorityElement(int[] nums){
        int n = nums.length;
        if(0==n)return n;
        else
        {
            Arrays.sort(nums);
            int result=0,count=0;
            int k=nums[0],v=0;
            for(int i=0;i<n;i++){
                if(k==nums[i]){
                    v++;
                    if(v>count){
                        count=v;
                        result = nums[i];
                    }
                    else if(v==count){
                        result = nums[i]>result?nums[i]:result;
                    }
                }
                else{
                    k=nums[i];
                    v=1;
                }
            }
            if(count>=n/2)
                return result;
            else return 0;
        }
    }
    //HashMap方法没有一点优势，时间复杂度太高
    public int majorityElement2(int[] nums){
        int n = nums.length;
        if(0==n)return n;
        else
        {
            int re=0,c=0;
            HashMap<Integer,Integer> map = new HashMap<Integer,Integer>();
            for(int i=0;i<n;i++){
                if(map.containsKey(nums[i]))
                    map.put(nums[i], map.get(nums[i]) +1 );
                else
                    map.put(nums[i],1);
            }
            for(Integer k:map.keySet()){
                int v = map.get(k);
                if(v>c){
                    c=v;
                    re = k;
                }
                if(v==c){
                    re = k;
                }
            }
            if(c>=n/2)
                return re;
            else return 0;
        }
    }

    //13.Roman to Integer
    public int romantoInteger(String s){
        int result = 0,n=0;
        if(s.length()<=0)return result;
        Map<Character,Integer> map = new HashMap();
        map.put('I',1);
        map.put('V',5);
        map.put('X',10);
        map.put('L',50);
        map.put('C',100);
        map.put('D',500);
        map.put('M',1000);
        for(int k=s.length()-1;k>=0;k--) {
            char c=s.charAt(k);
            n = map.get(c);
            if(k==0){
                char a = s.charAt(k);
                char b = k+1<s.length()?s.charAt(k+1):' ';
                if(map.get(a)<map.get(b))
                    result -=n;
                else
                    result +=n;
            }
            else
                result +=n;
        }
        return result;
    }

    //260 Single Number
    public int[] singleNumber(int[] nums){
        int[] result=new int[nums.length];
        if(nums.length<=0)return null;
        int index=0;
        for(int i=0;i<nums.length;i++){
            int j=0;
            int count=0;
            while(j<nums.length){
                if(nums[j]==nums[i])
                    count++;
                j++;
            }
            if(count==1) {
                result[index++]=nums[i];
            }
        }
        return Arrays.copyOfRange(result,0,index);
    }
    public int[] singleNumber2(int[] nums){
        int[] result=new int[nums.length];
        if(nums.length<=0)return null;
        Arrays.sort(nums);
        int index=0;
        int i=0,j=0;
        while(i<nums.length) {
            j = i + 1 < nums.length ? i + 1 : i - 1;
            if (nums[j] != nums[i]) {
                if (i == 0) result[index++] = nums[i];
                else {
                    if (nums[i] != nums[i - 1]) {
                        result[index++] = nums[i];
                    }
                }
            }
            i++;
        }
        return Arrays.copyOfRange(result,0,index);
    }

    public static void main(String[]args){
        Solution s = new Solution();
        TreeNode root = new TreeNode(4);
        root.left = new TreeNode(7);
        root.right = new TreeNode(9);
        root.left.left = new TreeNode(12);
        root.left.right = new TreeNode(45);
        //s.inOrder(root);
        int[] a={2,3,8,7,7,9,10};
        int[] b={8,3,7,2,9,5,10};
        //TreeNode se=s.buildTree(a, b);
        //s.preOrder(se);
        //TreeNode result = s.invertTree(root);
        //s.printTreeNode(root);
        //---------------------------------------------
        ListNode ls = new ListNode(0);
        ListNode ls1 = new ListNode(1);
        ListNode ls2 = new ListNode(6);
        ListNode ls3 = new ListNode(7);
        ls.next = ls1;
        ls1.next = ls2;
        ls2.next = ls3;
        //ListNode temp = s.reverseList(ls1);
        int n=13;
        //System.out.print(b);
        //System.out.print(s.isPowOfThree(27));
        System.out.println(s.titleToNumber("AC"));
        System.out.println(s.romantoInteger("IX"));
        int[] aa = s.singleNumber(new int[]{1, 2, 3, 5, 3,4,2,1});
        System.out.println(4^3);
    }


}
