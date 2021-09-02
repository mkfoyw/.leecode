/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
// dp[i][j]： 表示从字符串s[i,j]全部分割成回文串的最小分割次数
// dp[i][j] = max(dp[i][j], dp[i][i+1]+dp[i+1][j])
func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	flag := make([][]int8, 0)
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int8, n))
	}

	var isPalindrome func(int, int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}

		if flag[i][j] != 0 {
			return flag[i][j]
		}

		if s[i] == s[j] {
			flag[i][j] = isPalindrome(i+1, j-1)
		}
		return flag[i][j]
	}

	dp := make([]int, n)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < n; i++ {
		dp[i] = n
		for j := 0; j <= i; j++ {
			if isPalindrome(j, i) == 1 {
				if j == 0 {
					dp[i] = 1
				} else {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}
	return dp[n-1] - 1
}

/*
Accepted
33/33 cases passed (920 ms)
Your runtime beats 5.42 % of golang submissions
Your memory usage beats 75.9 % of golang submissions (6.4 MB)
*/

// @lc code=end

