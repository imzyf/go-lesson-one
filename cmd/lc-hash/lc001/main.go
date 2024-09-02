package main

// hash
func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		if first, ok := mp[target-v]; ok {
			// 差值 存在
			return []int{first, k}
		}

		// 不存在
		mp[v] = k

	}

	return nil

}
