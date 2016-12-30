package com.freedom.leetcode;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class LinkProblems {

    //2. Add Two Numbers
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int[] a = new int[4];
        if(l1==null && l2==null)
            return null;
        ListNode node = new ListNode(l1.val+l2.val);
        ListNode cur = node;
        ListNode pre = node;
        ListNode next = new ListNode(0);
        while(l1!=null || l2!=null){
            if(l1==null || l2==null) next.val += (l1==null)?l2.val:l1.val;
            else next.val += l1.val + l2.val;
            if(next.val>=10){
                cur.val = next.val-10;
                next = new ListNode(1);
            }
            else{
                cur.val = next.val;
                next = new ListNode(0);
            }
            pre = cur;
            cur.next = next;
            cur = next;
            if(l1!=null) l1 = l1.next;
            if(l2!=null) l2 = l2.next;
        }
        if(1!=cur.val)
            pre.next=null;
        return node;
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


}
