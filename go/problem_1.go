package main

func TwoSum(nums []int, target int) []int {
	var m map[int]int = make(map[int]int)
	for i, value := range nums {
		m[value] = i
	}

	var result []int = make([]int, 2)
	for i, value := range nums {
		second := target - value
		if item, ok := m[second]; ok && item != i {
			result[0] = i
			result[1] = item
			return result
		}
	}
	return nil
}
