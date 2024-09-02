package main

import "sort"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
	groupAnagramsCount(strs)
}

// hash sort
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

		mp[string(runes)] = append(mp[string(runes)], str)
	}

	ret := make([][]string, 0, len(mp))
	for _, v := range mp {
		ret = append(ret, v)
	}

	return ret
}

// hash 计数
func groupAnagramsCount(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, v := range strs {
		cnt := [26]int{}
		for _, w := range v {
			cnt[w-'a']++
		}

		mp[cnt] = append(mp[cnt], v)
	}

	ret := make([][]string, 0, len(mp))

	for _, v := range mp {
		ret = append(ret, v)
	}

	return ret

}
