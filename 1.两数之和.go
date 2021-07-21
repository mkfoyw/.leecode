/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mm := map[int]int{}

	for i, num := range nums {
		j, ok := mm[target-num]
		if ok {
			return []int{j, i}
		}
		mm[num] = i
	}

	return []int{}
}

// @lc code=end

