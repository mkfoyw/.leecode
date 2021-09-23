/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
// 子数组是数组中的一个连续序列
// 因此我们可以枚举等差数列的前两个元素， 如果后面元素与前一个元素满足等差数列的关系。

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}

	d := nums[1] - nums[0]
	ans, t := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == d {
			t++
		} else {
			d = nums[i] - nums[i-1]
			t = 0
		}

		ans += t
	}

	return ans

}

// @lc code=end

