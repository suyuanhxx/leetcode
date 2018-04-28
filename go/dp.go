package main

func LongestIncreaseList(list []int) int {
	n := len(list)
	dp := make([]int, n)
	dp[0] = 1
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if list[j] < list[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		res = Max(res, dp[i])
	}
	return dp[n-1]
}

func LongestList(list []int) int {
	n := len(list)
	dp := make([]int, n)
	dp[0] = 1
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if list[j]-list[i] == 1 || list[j]-list[i] == -1 {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		res = Max(res, dp[i])
	}
	return dp[n-1]
}
