

import org.junit.Before;
import org.junit.Test;

import java.LinkNode;
import java.LinkProblems;

/**
 * REVIEW
 * @Description:
 * @author xiaoxu.huang@baidao.com xiaoxu.huang
 * @date 2017/3/3  10:27
 *
 */
public class LinkTest {

	private LinkProblems linkProblems;

	@Before
	public void linkTest() {
		linkProblems = new LinkProblems();
	}

	@Test
	public void singleLinkTest() {
		LinkNode linkNode = new LinkNode(11);
		linkProblems.reverseSingleLink(linkNode);
		linkProblems.print(linkNode);

	}

}
