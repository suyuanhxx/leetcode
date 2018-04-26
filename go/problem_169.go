package main

func MajorityElement(nums []int) int {
	m := make(map[int]int)
	for _, i := range nums {
		if _, ok := m[i]; ok {
			m[i] += 1
			if m[i] > len(nums)/2 {
				return i
			}
		} else {
			m[i] = 1
		}
	}
	return nums[0];
}
