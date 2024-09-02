package main

// https://leetcode.cn/problems/climbing-stairs/description
func main() {
	climbStairs(4)
}

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)

	// base case
	dp[1] = 1
	dp[2] = 2
	// 3 111
	// 3 12
	// 3 21
	// 3 = 3
	for i := 3; i <= n; i++ {
		// 3 1+1  2+1
		// 4 2+2
		dp[i] = max(dp[i], dp[i-2]+2, dp[i-1]+1)
	}

	return dp[n]
}
