/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
// 如何判断一串字符串是否匹配呢？ 我们可以使用栈来判断
// 如果字符是 '(' 那么我们就把这个字符放入栈中。  如果 ')' 那么就从栈中弹出一个 '('
// 当栈中没有 ‘(' 说明字符串不匹配， 或者当字符串已经使用完了， 但是栈中还存在字符， 那么该字符串也无法匹配。
//
// dp[i]: 表示以第i个字符结尾的最长括号匹配的长度。
//       s[i] ==‘）’ 且 s[i-1] == '(', 那么该字符串的形式为 '....()'
//       s[i] ==')' 且 s[i-1] == ')', 那么字符串的形式为 '.....((...))'

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n, n)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-1-dp[i-1] >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i] >= 0 {
						dp[i] = dp[i] + dp[i-dp[i]]
					}
				}

			}
		}

		res = max(res, dp[i])
	}
	return res
}

// @lc code=end

