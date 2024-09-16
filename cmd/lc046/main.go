package main

/*
*
https://leetcode.cn/problems/permutations/description/
46. 全排列
*/
func main() {
	permute([]int{1, 2, 3})
}

// [1,2,3]
func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			// 深度
			dfs(path)
			// 回退
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}
