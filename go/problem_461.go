package main

//func hammingDistance(x int, y int) int {
//	a := binaryToDecimal(uint(x))
//	b := binaryToDecimal(uint(y))
//	count := 0
//	for i := 0; i < 32; i++ {
//		if a[i] != b[i] {
//			count++
//		}
//	}
//	return count
//}

//func hammingDistance(x int, y int) int {
//	n := x ^ y
//	a := binaryToDecimal(uint(n))
//	count := 0
//	for i := 0; i < 32; i++ {
//		if a[i] == 1 {
//			count++
//		}
//	}
//	return count
//}

func hammingDistance(x int, y int) int {
	n := x ^ y
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

func binaryToDecimal(n uint) [32]uint {
	var result [32]uint
	for i := 31; i >= 0; i-- {
		a := n >> uint(i) & 1
		result[32-i-1] = a
		//fmt.Print(a & 1)
	}
	return result
}

//func main() {
//	x := 10
//	y := 4
//	//binaryToDecimal(uint(x))
//	fmt.Println(hammingDistance(x, y))
//}
