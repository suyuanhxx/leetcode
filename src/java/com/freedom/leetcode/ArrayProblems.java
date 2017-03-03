package com.freedom.leetcode;

import java.util.Arrays;

/**
 * Created by xxhuang on 2016/5/17.
 */
public class ArrayProblems {

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
        return Arrays.copyOfRange(result, 0, index);
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

    //4. Median of Two Sorted Arrays
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        int[] total = new int[m+n];
        int i=0,j=0,k=0;
        while(i<m || j<n){
            if(i>=m&&j<n){
                total[k++]=nums2[j];
                j++;
            }
            else if(i<m&&j>=n){//else if和if使用结果会不同
                total[k++]=nums1[i];
                i++;
            }
            else{
                if(nums1[i]<nums2[j]){
                    total[k++]=nums1[i];
                    i++;
                }
                else if(nums1[i]>nums2[j]){
                    total[k++]=nums2[j];
                    j++;
                }
                else{
                    total[k++]=nums1[i];
                    total[k++]=nums2[j];
                    i++;
                    j++;
                }
            }

        }
        double result = (m+n)%2!=0?total[(m+n)/2]:((double)total[(m+n)/2-1]+(double)total[(m+n)/2])/2;
        return result;
    }
}
