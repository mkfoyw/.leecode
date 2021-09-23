/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */

// @lc code=start

//弱等差数列表示至少有两个数的等差数列
// dp[i][d]: 表示以nums[i]为等差数列的尾元素以d为公差的弱等差数列的总和
// 那么 dp[i][d] += dp[j][d] + 1 ((nums[i],nums[j]) 这一对元素也可以当作一个弱等差子序列)
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	dp := make([]map[int]int, n)

	res := 0
	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			res += dp[j][d]
			dp[i][d] += dp[j][d] + 1
		}
	}
	return res
}

// @lc code=end

