/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start

// 最小堆解法， 每次从堆中找出一个最小的数， 然后乘以2，3，5。 然后一次重复这个步骤。
// 动态规划方法：
// dp[i] 表示第 i 个丑数
func nthUglyNumber(n int) int {
	p2, p3, p5 := 0, 0, 0

	dp := make([]int, n)
	dp[0] = 1

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, min(x3, x5))
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n-1]

}

// @lc code=end

