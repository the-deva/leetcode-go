/*
2501 : https://leetcode.com/problems/longest-square-streak-in-an-array/description
1. make set of numbers and dfs the nums in set
2. If already exists return it else dfs(x * x)
*/
package main

func longestSquareStreak(nums []int) (result int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}

	var dfs func(x int) int
	dfs = func(x int) int {
		// If not exists return 0
		if !set[x] {
			return 0
		}
		dp := map[int]int{}
		if v, ok := dp[x]; ok {
			// if v already exists return
			return v
		}
		dp[x] = 1 + dfs(x*x)
		return dp[x]
	}
	for x := range set {
		result = max(result, dfs(x))
	}
	if result == 1 {
		return -1
	}
	return
}

func main() {
	nums := []int{4, 3, 6, 16, 8, 2}
	longestSquareStreak(nums)
}
