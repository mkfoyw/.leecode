/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */

// @lc code=start
func findRotateSteps(ring string, key string) int {
	nr, nk := len(ring), len(key)
	if nr == 0 || nk == 0 {
		return 0
	}

	var dfs func(int, int) int

	dp := make([]map[int]int, nr)
	for i := 0; i < nr; i++ {
		dp[i] = map[int]int{}
	}

	dfs = func(pr, pk int) int {

		if _, ok := dp[pr][pk]; ok {
			return dp[pr][pk]
		}

		step := 0
		if ring[pr] == key[pk] {
			step += 1
			pk++
		}

		if pk == nk {
			return step
		}

		t1, t2 := 0, 0
		if pr+1 >= nr {
			t1 = dfs(0, pk)
		} else {
			t1 = dfs(pr+1, pk)
		}

		if pr-1 < 0 {
			t2 = dfs(nr-1, pk)
		} else {
			t2 = dfs(pr-1, pk)
		}

		step += min(t1, t2) + 1

		dp[pr][pk] = step
		return dp[pr][pk]
	}

	return dfs(0, 0)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

