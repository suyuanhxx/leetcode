package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	// 换底公式
	// logaB = logcB/logcA
	a := int(math.Log(float64(n)) / math.Log(3))
	b := math.Pow(3, float64(a))
	if int(b) == n {
		return true
	}
	return false
}

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	a := int(math.Log(float64(num)) / math.Log(3))
	b := math.Pow(3, float64(a))
	if int(b) == num {
		return true
	}
	return false
}

func main() {
	n := 9
	fmt.Println(3)
	fmt.Println(math.Log(float64(n)))
	fmt.Println(isPowerOfThree(n))
}
