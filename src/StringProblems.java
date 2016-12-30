package com.freedom.leetcode;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class StringProblems {

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
}
