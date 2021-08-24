/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

// dp[i][]string: 表示以字符s[i]结尾的字符串有的分割方案

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}

	//flag[i][j]:表示字符串s[i]到s[j]是否是回文串
	flag := make([][]int8, n)
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int8, n))
	}

	// 1 表示是回文串， -1 表示不是回文串， 0 表示未搜索
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if flag[i][j] != 0 {
			return flag[i][j]
		}

		flag[i][j] = -1
		if s[i] == s[j] {
			flag[i][j] = isPalindrome(i+1, j-1)
		}
		return flag[i][j]
	}

	dp := make([][][]string, n+1, n+1)
	for i := 0; i < n; i++ {
		dp[i] = [][]string{}
	}

	//dp[i]:前 i 个字符串可能的分割方案
	for i := 0; i < n; i++ {
		cur := [][]string{}
		for j := 0; j <= i; j++ {
			if isPalindrome(i, j) == 1 {
				if j == 0 {
					cur = append(cur, append([]string(nil), s[:i+1]))
				} else {
					for _, item := range dp[j-1] {
						cur = append(cur, append([]string(nil), item...))
						cur[len(cur)-1] = append(cur[len(cur)-1], s[j:i+1])
					}
				}
			}

		}
		dp = append(dp, cur)
	}
	return dp[n-1]
}

// @lc code=end
