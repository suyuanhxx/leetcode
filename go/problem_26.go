package main

func removeDuplicates(nums []int) int {
	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			for j := i + 2; j < len(nums); j++ {
				if nums[j] > nums[i] {
					nums[i+1] = nums[j]
					break
				}
			}
		}
	}
	for i = 0; i < len(nums); i++ {
		if nums[len(nums)-1] == nums[i] {
			break
		}
	}
	return i + 1
}

//func main() {
//
//	nums := []int{0, 0, 0, 0, 0, 1, 2, 2, 3, 3, 4, 4}
//	fmt.Println(removeDuplicates(nums))
//	for _, num := range nums {
//		fmt.Print(num, " ")
//	}
//}
