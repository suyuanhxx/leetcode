

import com.freedom.leetcode.link.SingleLink;
import org.junit.Before;
import org.junit.Test;

import com.freedom.leetcode.link.LinkNode;

/**
 * REVIEW
 * @Description:
 * @author xiaoxu.huang
 * @date 2017/3/3  10:27
 *
 */
public class LinkTest {

	private SingleLink singleLink;
	private LinkNode head = new LinkNode(11);

	@Before
	public void linkTest() {
		singleLink = new SingleLink();
		LinkNode a = new LinkNode(2);
		LinkNode b = new LinkNode(3);
		LinkNode c = new LinkNode(4);
		head.next = a;
		a.next = b;
		b.next = c;
	}

	@Test
	public void singleLinkTest() {
		singleLink.print(head);
		LinkNode result = singleLink.reverseSingleLink(head);
		singleLink.print(result);
	}

}
