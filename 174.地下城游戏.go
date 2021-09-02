
/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start

// 这有个问题，

// dp[i][j]: 表示走到该位置所需要的最少体力

func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	if n == 0 {
		return 0
	}
	m := len(dungeon[0])
	if m == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp[n-1][m], dp[n][m-1] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

// func calculateMinimumHP(dungeon [][]int) int {
// 	if len(dungeon) == 0 {
// 		return 0
// 	}
// 	n, m := len(dungeon), len(dungeon[0])

// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}

// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}

// 	var dfs func(int, int, int, int)
// 	res := -11111111
// 	dfs = func(i, j, need, left int) {
// 		if i == n-1 && j == m-1 {
// 			res = max(need, res)
// 			return
// 		}
// 		//往右走
// 		if j+1 < m {
// 			dfs(i, j+1, min(need, left+dungeon[i][j+1]), left+dungeon[i][j+1])
// 		}
// 		//往下走
// 		if i+1 < n {
// 			dfs(i+1, j, min(need, left+dungeon[i+1][j]), left+dungeon[i+1][j])
// 		}
// 	}

// 	dfs(0, 0, min(0, dungeon[0][0]), dungeon[0][0])
// 	return -res + 1
// }

// @lc code=end
