package com.freedom.leetcode.string;

import org.junit.Test;

import java.util.LinkedHashMap;
import java.util.Map;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class StringProblems {

    //字符串反转
    public String reversString(String s) {
        //StringBuilder str=new StringBuilder(s);
        //return str.reverse().toString();
        char[] str = s.toCharArray();
        int len = s.length() - 1;
        for (int i = (len - 1) >> 1; i >= 0; i--) {
            int j = len - i;
            char a = str[i];
            char b = str[j];
            str[j] = a;
            str[i] = b;
        }
        return String.valueOf(str);
    }

    /**
     * 387. First Unique Character in a String
     *
     * @param s
     * @return
     */
    public int firstUniqChar(String s) {
        Map<Character, Integer> map = new LinkedHashMap<>();
        for (int i = 0; i < s.length(); i++) {
            Integer count = map.get(s.charAt(i));
            map.put(s.charAt(i), (null == count ? 1 : count + 1));
        }
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue().equals(1)) {
                return s.indexOf(entry.getKey());
            }
        }
        return -1;
    }

    @Test
    public void testString() {
        String s = "leetcodewewrsl";
        firstUniqChar(s);
    }
}
