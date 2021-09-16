/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */
// n = 0   不同的位数 1
// n =  1  不同的位数 10
//  n位数， 有多少个位数不同的数字呢
//   = 9 *(10-1)*(10-2)*...*(10-n+1)
// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n >= 10 {
		n = 10
	}

	dp := []int{1, 9}

	for i := 2; i <= n; i++ {
		dp = append(dp, 9)
		for j := 1; j < i; j++ {
			dp[i] = dp[i] * (10 - j)
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		ans += dp[i]
	}
	return ans
}

// @lc code=end



