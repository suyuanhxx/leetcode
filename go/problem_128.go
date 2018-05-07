package main

//Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
//
//Your algorithm should run in O(n) complexity.
//
//Example:
//
//Input: [100, 4, 200, 1, 3, 2]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

/**
solution 1: first sort
 */

//func LongestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	sort.Ints(nums)
//	max, current := 1, 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1] {
//			continue
//		}
//
//		if nums[i]-nums[i-1] == 1 {
//			current ++
//		} else {
//			if current > max {
//				max = current
//			}
//			current = 1
//		}
//	}
//
//	if current > max {
//		max = current
//	}
//	return max
//}

/**

 */
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] = v
	}

	length := 0

	for _, item := range nums {
		currLength := 1
		if _, ok := m[item-1]; !ok {
			currNum := item
			for {
				if _, ok := m[currNum+1]; ok {
					currLength ++
					currNum++
				} else {
					break
				}
			}
			if currLength > length {
				length = currLength
			}
		}
	}

	return length
}
