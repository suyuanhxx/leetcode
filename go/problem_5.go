package main

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var res string
	for i := 1; i < len(s); i++ {
		checkPalindromeExpand(s, i, i)
		checkPalindromeExpand(s, i, i+1)
	}
	return res
}

func checkPalindromeExpand(s string, low, high int) bool {
	for low > 0 && high < len(s) {
		if s[low] == s[high] {
			low--
			high++
		}
	}
	if low == 0 && high == len(s) {
		return true
	}
	return false
}