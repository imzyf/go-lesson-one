package main

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	longestConsecutive(nums)

}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet {
		if !numSet[num-1] {
			// 前一个值 不存在 - 这个值 可以作最小值
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}

		}
	}

	return longestStreak
}
