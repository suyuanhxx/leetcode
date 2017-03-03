package com.freedom.leetcode; /**
 * Created by xxhuang on 2016/4/28.
 */

import java.util.*;

public class Solution {

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
        HashMap<Character,Integer> map = new HashMap();
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

    //9. Palindrome Number
    public boolean isPalindrome(int x) {
        String str = String.valueOf(x);
        int len = str.length();
        if(0==len)return false;
        int i=0,j=len-1;
        for(;i<len/2;i++,j--){
            if(str.charAt(i)!=str.charAt(j))
                return false;
        }
        return true;
    }

    public boolean isPalindrome2(int x) {
        //HashMap a = new HashMap<>();
        HashSet aa = new HashSet();
        int i=x,j=0;
        while(i>=10){
            i=i/10;
            j++;
        }
        //bits=j
        int bits = (int)Math.log10((double)x);
        while(j>0){
            int k = x%10;
            int temp = (int)Math.pow(10,j);
            i = (x-x%temp) / temp;
            if(k!=i)
                return false;
            x = (x - x%10*temp)/10;
            j-=2;
        }
        return true;
    }

	public static void main(String[] args) {
		HashMap<Integer,Integer> map  = new HashMap();
        map.put(1,1);
        map.put(1,2);
        int q = map.get(1);
        String sre= "s";
        StringBuilder sb = new StringBuilder("aaa");
        StringBuffer sf = new StringBuffer("aaa");
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
		LinkNode ls = new LinkNode(0);
		LinkNode ls1 = new LinkNode(1);
		LinkNode ls2 = new LinkNode(6);
		LinkNode ls3 = new LinkNode(7);
		ls.next = ls1;
        ls1.next = ls2;
        ls2.next = ls3;
		//LinkNode temp = s.reverseList(ls1);
		//System.out.print(b);
        //System.out.print(s.isPowOfThree(27));
        System.out.println(s.titleToNumber("AC"));
        System.out.println(s.romantoInteger("IX"));
        ArrayProblems ap = new ArrayProblems();
        int[] aa = ap.singleNumber(new int[]{1, 2, 3, 5, 3,4,2,1});
        System.out.println(4^3);
        LinkProblems lk = new LinkProblems();
        //lk.addTwoNumbers(ls, ls1);
        //ap.findMedianSortedArrays(new int[]{},new int[]{2,3});
        System.out.println(s.isPalindrome2(1221));
        System.out.println(Math.pow(2,7)-1);
        LinkedList linkedList = new LinkedList();
        List al = new ArrayList();
        Vector v = new Vector();
    }

}
