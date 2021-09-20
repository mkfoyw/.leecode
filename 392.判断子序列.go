/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start

//在使用双指针法， 我们会花费大量的时间在寻在与 s[i] 相匹配的字符串t[j]
// 因此我们定义这样一个表格 f[i][j] 表示在字符串t[i]以后第一个出现字符 j(a <=j<=z)的位置

func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	if ns == 0 || nt == 0 {
		if ns == 0 {
			return true
		}

		if ns != 0 {
			return false
		}
	}

	f := make([][26]int, nt+1)

	for i := 0; i < 26; i++ {
		f[nt][i] = nt
	}

	for i := nt - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte('a'+j) {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}

	add := 0
	for i := 0; i < ns; i++ {
		if f[add][int(s[i]-'a')] == nt {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}
	return true
}

// 双指针法
func isSubsequence2(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 || m == 0 {
		if n == 0 {
			return true
		}

		if m == 0 {
			return false
		}
	}

	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == n {
		return true
	}
	return false
}

// dp[i][j]： 表示字符串s前i个字符串是否是字符串t前j个字符串的子序列
func isSubsequence1(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 || m == 0 {
		if n == 0 {
			return true
		}
		if m == 0 {
			return false
		}
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= n; i++ {
		for j := i; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] || dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

// @lc code=end
