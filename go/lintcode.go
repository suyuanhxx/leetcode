package main

import (
	"fmt"
	"math/big"
	"math"
	"strconv"
	"strings"
)

func trailingZeros(n int64) int64 {
	count := int64(0)
	for i := 1; math.Pow(float64(5), float64(i)) < float64(n); i++ {
		count += n / int64(math.Pow(float64(5), float64(i)))
	}
	return count
}

func digitCounts(k int, n int) int {
	if k > n {
		return 0
	}
	count := 0
	for i := 0; i <= n; i++ {
		j := strings.Count(strconv.Itoa(i), strconv.Itoa(k))
		if j > 0 {
			count += j
		}
	}
	return count
}

func factorial(n *big.Int) (result *big.Int) {
	intOne := big.NewInt(1)
	if n.Cmp(intOne) <= 0 {
		result = big.NewInt(1)
	} else {
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, intOne)))
	}
	return result
}

func mergeSortedArray(A []int, B []int) []int {
	// write your code here
	if len(A) == 0 {
		return B
	}
	if len(B) == 0 {
		return A
	}

	if A[0] <= B[0] {
		A = mergeArray(B, A)
		return A
	} else {
		B = mergeArray(A, B)
		return B
	}
}

func mergeArray(insertArray []int, result []int) []int {
	for _, item := range insertArray {
		for j := 0; j < len(result); j++ {
			if j+1 < len(result) && item >= result[j] && item <= result[j+1] {
				result = append(result[:j+1], append([]int{item}, result[j+1:]...)...)
				break
			}
			if j+1 >= len(result) {
				result = append(result[:j+1], append([]int{item}, result[j+1:]...)...)
				break
			}
		}
	}
	return result
}

func main() {
	//n := big.NewInt(10)
	//fmt.Print(factorial(n))
	//
	//x := new(big.Int)
	//x.MulRange(1, 10)
	//fmt.Print(x)

	A := []int{1, 2, 3, 4}
	B := []int{2, 4, 5, 6}

	//B = append(B[:0], append([]int{1}, B[0:]...)...)
	//B = append(B[:1], append([]int{1}, B[1:]...)...)

	fmt.Println(mergeSortedArray(A, B))
	fmt.Println(digitCounts(1, 12))
	n := int64(20)
	fmt.Println(trailingZeros(n))
}
