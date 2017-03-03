

import org.junit.Before;
import org.junit.Test;

import com.freedom.leetcode.LinkNode;
import com.freedom.leetcode.LinkProblems;

/**
 * REVIEW
 * @Description:
 * @author xiaoxu.huang@baidao.com xiaoxu.huang
 * @date 2017/3/3  10:27
 *
 */
public class LinkTest {

	private LinkProblems linkProblems;
	private LinkNode head = new LinkNode(11);

	@Before
	public void linkTest() {
		linkProblems = new LinkProblems();
		LinkNode a = new LinkNode(2);
		LinkNode b = new LinkNode(3);
		LinkNode c = new LinkNode(4);
		head.next = a;
		a.next = b;
		b.next = c;
	}

	@Test
	public void singleLinkTest() {
		linkProblems.print(head);
		LinkNode result = linkProblems.reverseSingleLink(head);
		linkProblems.print(result);
	}

}
