/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
// 也是0-1 背包
// dp[i][j][k]： 表示前 i 个字符串中0的个数为j， 1的个数为k， 最多能够选取多少个字符串。
func findMaxForm(strs []string, m int, n int) int {
	sn := len(strs)
	if sn == 0 {
		return 0
	}

	zeros := make([]int, sn)
	ones := make([]int, sn)

	for i := 0; i < sn; i++ {
		ones[i] = strings.Count(strs[i], "1")
		zeros[i] = len(strs[i]) - ones[i]
	}

	dp := make([][][]int, sn+1)
	for i := 0; i <= sn; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i <= sn; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if zeros[i-1] > j || ones[i-1] > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeros[i-1]][k-ones[i-1]]+1)
				}
			}
		}
	}

	return dp[sn][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

