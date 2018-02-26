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

func main() {
	//n := big.NewInt(10)
	//fmt.Print(factorial(n))
	//
	//x := new(big.Int)
	//x.MulRange(1, 10)
	//fmt.Print(x)

	fmt.Println(digitCounts(1, 12))
	n := int64(20)
	fmt.Println(trailingZeros(n))
}
