/*
 * @lc app=leetcode.cn id=1130 lang=golang
 *
 * [1130] 叶值的最小代价生成树
 */

// @lc code=start
// 解题思路：
// 我们可以看到数组可以分为两部分， 一边是左子树， 一边是右子树。
// 根节点的值为左右两边的最小值的乘积

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	// mn[i][j] 表示数组 arr[i, j] 最小值
	mn := make([][]int, n)
	// dp[i][j] 表示数组 arr[i, j] 构建二叉树的最小花费
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		mn[i] = make([]int, n)
		dp[i] = make([]int, n)
		dp[i][i] = 0
		mn[i][i] = arr[i]
		for j := i + 1; j < n; j++ {
			dp[i][j] = math.MaxInt32
			mn[i][j] = math.MaxInt32
		}
	}

	var dfs func(int, int) int

	dfs = func(l, r int) int {
		if dp[l][r] != math.MaxInt32 {
			return dp[l][r]
		}

		for k := l; k < r; k++ {

			dp[l][r] = min(dp[l][r], dfs(l, k)+dfs(k+1, r)+mn[l][k]*mn[k+1][r])
			mn[l][r] = max(mn[l][k], mn[k+1][r])

		}
		return dp[l][r]
	}

	return dfs(0, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

