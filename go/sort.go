package main

import "fmt"

func main() {
	array := []int{2, 1, 4, 7, 5, 3, 6, 8, 9, 0}
	print(array)
	result := bubbleSort(array)
	print(result)
	result = quickSort(array)
	print(result)
}

func bubbleSort(array []int) []int {
	result := copy(array)
	for i := range result {
		j := i + 1
		for ; j < len(result); j++ {
			if result[j] < result[i] {
				temp := result[i]
				result[i] = result[j]
				result[j] = temp
			}
		}
	}
	return result
}

func quickSort(array []int) []int {
	return nil
}

// deep copy
func copy(array []int) []int {
	return append(make([]int, 0, len(array)), array...)
}

func print(array []int) {
	for _, num := range array {
		fmt.Print(num, "  ")
	}
	fmt.Println()
}
