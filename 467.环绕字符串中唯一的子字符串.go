/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 */

// @lc code=start
// 刚开始没看明白题意。 题意是p的不同子串出现在环绕子串的数目。
// 现在的方法是对p的子串进行去重， 然后在判断子串是否在环绕子串中。

// 由于环绕子串的特殊性， 我们可以统计p中以某个字符结尾最长出现在环绕子串的数目

func findSubstringInWraproundString(p string) int {
	n := len(p)
	if n == 0 {
		return 0
	}

	dp := make([]int, 26)
	cnt := 0
	for i := 0; i < n; i++ {
		if i > 0 && (p[i-1]+1 == p[i] || (p[i] == 'a' && p[i-1] == 'z')) {
			cnt++
		} else {
			cnt = 1
		}

		dp[p[i]-'a'] = max(dp[p[i]-'a'], cnt)
	}
	res := 0
	for i := 0; i < 26; i++ {
		res += dp[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

