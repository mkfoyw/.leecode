
/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	mn, mx, ans := nums[0], nums[0], nums[0]

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		mn, mx = min(nums[i], min(nums[i]*mn, nums[i]*mx)), max(nums[i], max(nums[i]*mn, nums[i]*mx))
		ans = max(mx, ans)
	}
	return ans
}

// @lc code=end
