package com.freedom.leetcode.link;

/**
 * REVIEW
 * @Description:
 * @author xiaoxu.huang@baidao.com xiaoxu.huang
 * @date 2017/3/3  11:18
 *
 */
public class SingleLink {

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

	public void print(LinkNode head) {
		while (null != head) {
			System.out.print(head.val + " ---> ");
			head = head.next;
		}
		System.out.println("null");
	}
}
