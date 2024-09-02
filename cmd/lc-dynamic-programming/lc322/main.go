package main

import "math"

/*
https://leetcode.cn/problems/coin-change/
*/
func main() {
	coinChange([]int{1, 2, 5}, 11)
	coinChangeAdd([]int{1, 2, 5}, 11)
}

/*
*
回到凑零钱问题，为什么说它符合最优子结构呢

// 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币

	func dp(coins []int, n int) int {
	    // 初始化为最大值
	    res := math.MaxInt32

	    // 做选择，选择需要硬币最少的那个结果
	    for _, coin := range coins {
	        res = min(res, 1+subProblem)
	    }
	    return res
	}
*/
func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

// 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币
func dp(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt32

	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dp(coins, amount-coin)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		res = min(res, subProblem+1)
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
func coinChangeWithMemo(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		memo[i] = -2
	}
	return dpWithMemo(coins, amount, memo)
}
func dpWithMemo(coins []int, amount int, memo []int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if memo[amount] != -2 {
		return memo[amount]
	}

	res := math.MaxInt32
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dpWithMemo(coins, amount-coin, memo)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		res = min(res, subProblem+1)
	}

	if res == math.MaxInt32 {
		memo[amount] = -1
		return -1
	}

	memo[amount] = res

	return res
}

func coinChangeAdd(coins []int, amount int) int {
	// 数组大小为 amount + 1，初始值也为 amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])

		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChangeAddSelf(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		// 一定有 1 所以理论最大是 amount
		// [amount+1] 是个特殊的魔法值
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 自上往上
	//
	for subAmount := 1; subAmount <= amount; subAmount++ {
		// 遍历每个硬币
		for _, coin := range coins {
			if subAmount-coin < 0 {
				// 这颗币不能用
				continue
			}
			dp[subAmount] = min(dp[subAmount], 1+dp[subAmount-coin])

		}
	}

	// 可能有无法计算出来的值
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
