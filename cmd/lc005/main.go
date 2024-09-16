package main

func main() {
	longestPalindrome("abcb")
	longestPalindromeDp("abcb")
}

func longestPalindrome(s string) string {
	res := ""
	max := 0

	for k := 0; k < len(s); k++ {
		left, right := k-1, k+1
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left+1 > max {
					max = right - left + 1
					res = s[left : right+1]
				}

				left--
				right++
			} else {
				break
			}
		}

		left, right = k-1, k
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left+1 > max {
					max = right - left + 1
					res = s[left : right+1]
				}

				left--
				right++
			} else {
				break
			}
		}
	}

	if max == 0 && len(s) > 0 {
		return string(s[0])
	}

	return res
}

func longestPalindromeDp(s string) string {
	if s == "" {
		return ""
	}

	dp := make([][]bool, len(s))
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 根据case 1 初始数据
	}

	// case 2
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1

			if s[start] != s[end] {
				dp[start][end] = false
			} else if length < 3 {
				// 首尾相同
				dp[start][end] = true

			} else {
				// 首尾相同 看斜角
				dp[start][end] = dp[start+1][end-1]
			}

			if dp[start][end] && length > len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}

	return result

}
