package main

//func main() {
//	array := []int{2, 1, 4, 7, 5, 3, 6, 8, 9, 0}
//	print(array)
//	result := BubbleSort(array)
//	print(result)
//	result = QuickSort(array)
//	print(result)
//}

func BubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j] > s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func SelectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		max := s[i]
		k := i
		for j := i; j < len(s); j++ {
			if s[j] > max {
				max = s[j]
				k = j
			}
		}
		s[i], s[k] = s[k], s[i]
	}
}

func QuickSort(array []int) []int {
	return nil
}
