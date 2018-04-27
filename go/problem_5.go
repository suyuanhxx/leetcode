package main

import "fmt"

func LongestPalindrome(s string) string {
	return ""
}

// Longest common sequence length
func LongestCommonSubLen(A string, B string) int {
	LCS := make([][]int, len(A)+1)
	for x := 0; x <= len(A); x++ {
		LCS[x] = make([]int, len(B)+1)
	}
	x, y := 0, 0
	for ; x <= len(A); x++ {
		for y = 0; y <= len(B); y++ {
			if (x == 0 || y == 0) {
				LCS[x][y] = 0
			} else if A[x-1] == B[y-1] {
				LCS[x][y] = LCS[x-1][y-1] + 1
			} else {
				LCS[x][y] = max(LCS[x-1][y], LCS[x][y-1])
			}
		}
	}
	fmt.Println(reverseString(LongestCommonSub(LCS, A, B)))
	return LCS[len(A)][len(B)]
}

// Longest common sequence string
func LongestCommonSub(s [][]int, A string, B string) (result string) {
	i := len(A)
	j := len(B)
	for i > 0 && j > 0 {
		if A[i-1] == B[j-1] {
			i, j = i-1, j-1
			result = result + string(A[i])
		} else {
			if s[i][j-1] == s[i-1][j] {
				i -= 1
			} else if s[i][j-1] > s[i-1][j] {
				j -= 1
			} else {
				i -= 1
			}
		}
	}
	return result
}

// reverse string
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// return max(a,b)
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
