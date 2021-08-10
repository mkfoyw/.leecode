/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res, current := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if current < 0 {
			current = nums[i]
		} else {
			current += nums[i]
		}

		res = max(current, res)
	}
	return res
}

// @lc code=end

