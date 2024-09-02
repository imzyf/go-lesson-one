package main

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长
子串的长度。

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func main() {
	lengthOfLongestSubstring("abcb")

}

func lengthOfLongestSubstring(s string) int {
	mp := map[string]int{}
	ret := 0

	// 遍历一次就可以
	for left, right := 0, 0; right < len(s); right++ {
		// 当前单词
		word := string(s[right])
		// 是否重复
		if index, ok := mp[word]; ok && index >= left {
			// 滑动左边窗口
			left = index + 1
		}
		mp[word] = right

		ret = max(ret, right-left+1)
	}

	return ret

}
