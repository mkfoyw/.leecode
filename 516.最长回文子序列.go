/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
// dp[i][j]: 表示字符串s[i,j] 最长回文串的长度
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return 1
	}

	var dfs func(int, int) int
	dp := make([]map[int]int, n)

	dfs = func(left, right int) int {
		if left > right {
			return 0
		}

		if left == right {
			return 1
		}

		if _, ok := dp[left][right]; ok {
			return dp[left][right]
		}

		t1, t2 := 0

		if s[left] == right {
			t1 = 1 + dfs(left+1, right-1)
		} else {
			t2 = max(dfs(left, right-1), dfs(left+1, right))
		}

		dp[left][right] = max(t1, t2)
		return dp[left][right]
	}
	return dfs(0, n-1)
}

func longestPalindromeSubseq2(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var dfs func(int, int) int

	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
	}
	dfs = func(left, right int) int {
		if left > right {
			return 0
		}

		if left == right {
			return 1
		}

		if _, ok := dp[left][right]; ok {
			return dp[left][right]
		}
		//选择left
		j := right
		for ; j > left; j-- {
			if s[left] == s[j] {
				break
			}
		}
		t1 := 0
		if left != j {
			t1 = 2
		} else {
			t1 = 1
		}
		t1 += dfs(left+1, j-1)

		//不选择left
		t2 := dfs(left+1, right)

		dp[left][right] = max(t1, t2)
		return dp[left][right]
	}
	return dfs(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

