package com.freedom.leetcode;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class LinkProblems {

	//2. Add Two Numbers
	public LinkNode addTwoNumbers(LinkNode l1, LinkNode l2) {
		int[] a = new int[4];
		if (l1 == null && l2 == null)
			return null;
		LinkNode node = new LinkNode(l1.val + l2.val);
		LinkNode cur = node;
		LinkNode pre = node;
		LinkNode next = new LinkNode(0);
		while (l1 != null || l2 != null) {
			if (l1 == null || l2 == null)
				next.val += (l1 == null) ? l2.val : l1.val;
			else
				next.val += l1.val + l2.val;
			if (next.val >= 10) {
				cur.val = next.val - 10;
				next = new LinkNode(1);
			} else {
				cur.val = next.val;
				next = new LinkNode(0);
			}
			pre = cur;
			cur.next = next;
			cur = next;
			if (l1 != null)
				l1 = l1.next;
			if (l2 != null)
				l2 = l2.next;
		}
		if (1 != cur.val)
			pre.next = null;
		return node;
	}

	//单链表反转LeetCode OJ 206

	/**
	 * 单链表反转
	 * step1. head ---> ...... ---> q
	 *      	^
	 *      	|
	 *      	h
	 * step2. head ---> next ---> ...... ---> q
	 *     			 	^
	 *      		 	|
	 *    	  head <--- h
	 * pre 上一个节点
	 * h 中间节点，反转之后头部
	 * next 下一个节点
	 * @param head
	 * @return
	 */
	public LinkNode reverseSingleLink(LinkNode head) {
		if (null == head)
			return null;
		LinkNode h = null;
		LinkNode pre = null;
		while (null != head) {
			LinkNode next = head.next;
			h = new LinkNode(head.val);
			h.next = pre;
			pre = h;
			head = next;
		}
		return h;
	}

	public void deleteNode(LinkNode node) {
		if (null == node)
			return;
		LinkNode next = node.next;
		while (null != next) {
			node.val = next.val;
			node = next;
			next = next.next;
		}
	}

	public void print(LinkNode head) {
		while (null != head) {
			System.out.print(head.val + "--->");
			head = head.next;
		}
		System.out.println("null");
	}

}
