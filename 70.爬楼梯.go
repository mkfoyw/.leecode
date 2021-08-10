/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	f1 := 1
	f2 := 1

	for i := 2; i <= n; i++ {
		f := f1 + f2
		f1 = f2
		f2 = f
	}
	return f2
}

// @lc code=end

