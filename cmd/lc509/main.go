package main

func main() {
	fib2(3)
}

// O(2^n)。
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// 只不过这种解法是「自顶向下」进行「递归」求解，我们更常见的动态规划代码是「自底向上」进行「递推」求解。
func fib2(n int) int {
	// 备忘录全初始化为 0
	// n+1 是因为下标是从 0 开始
	memo := make([]int, n+1)
	// 进行带备忘录的递归
	return dp(memo, n)
}

func dp(memo []int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {

		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = dp(memo, n-1) + dp(memo, n-2)
	return memo[n]

}

func fib3(n int) int {
	if n == 0 || n == 1 {
		return 0
	}

	fibRet := 0
	fib1, fib2 := 1, 0
	for i := 2; i <= n; i++ {
		fibRet = fib1 + fib2
		fib1 = fibRet
		fib2 = fib1
	}

	return fibRet
}
