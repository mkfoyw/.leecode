/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
// dp[i][j] 表示将word1 的前 i 个字符转换 word2 的前 j个字符所需的最小操作数
// 在此时有3种操作
// 添加一个字符 dp[i][j-1] + 1
// 删除一个字符 dp[i-1][j] + 1
// 替换一个字符 dp[i-1][j-1] + 1 此时 word1[i] != word2[j]
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}

	dp := make([][]int, m+1, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for j := 0; j <= m; j++ {
		dp[j][0] = j
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

