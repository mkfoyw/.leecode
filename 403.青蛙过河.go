/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start

// dp[cur][step]： 表示在第 cur 个存在石头的格子步长为 step 时能否跳到终点
func canCross(stones []int) bool {
	n := len(stones)
	if n == 0 {
		return true
	}

	dp := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = map[int]bool{}
	}

	var dfs func(int, int) bool
	dfs = func(cur int, step int) bool {
		if cur == n-1 {
			return true
		}

		if res, ok := dp[cur][step]; ok {
			return res
		}

		res := false
		pos := 0
		//step
		if step != 0 {
			pos = sort.SearchInts(stones, stones[cur]+step)
			res = pos != n && stones[pos] == stones[cur]+step && dfs(pos, step)
		}

		//step-1
		if step-1 > 0 && !res {
			pos = sort.SearchInts(stones, stones[cur]+step-1)
			res = pos != n && stones[pos] == stones[cur]+step-1 && dfs(pos, step-1)
		}

		//step + 1
		if !res {
			pos = sort.SearchInts(stones, stones[cur]+step+1)
			res = pos != n && stones[pos] == stones[cur]+step+1 && dfs(pos, step+1)
		}

		dp[cur][step] = res
		return res
	}

	return dfs(0, 0)

}

// @lc code=end

