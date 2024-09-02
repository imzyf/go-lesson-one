package main

/*
438. 找到字符串中所有字母异位词

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/
func main() {
	findAnagrams("cbaebabacd", "abc")
}
func findAnagrams(s string, t string) []int {
	need, window := make(map[string]int), make(map[string]int)
	for _, c := range t {
		need[string(c)]++
	}

	left, right := 0, 0
	valid := 0
	// 记录结果
	var res []int
	for right < len(s) {
		c := string(s[right])
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(t) {
			// 当窗口符合条件时，把起始索引加入 res
			if valid == len(need) {
				res = append(res, left)
			}
			d := string(s[left])
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
