/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 */

// @lc code=start
// 状态压缩， 把下面vis数组用一个整数来表示， 更方便使用记忆化搜索记录下来。
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger > desiredTotal {
		return true
	}

	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	dp := make(map[int]bool)
	var dfs func(cur, state int) bool
	dfs = func(cur, state int) bool {
		if _, ok := dp[state]; ok {
			return dp[state]
		}

		res := false
		for i := 0; i < maxChoosableInteger; i++ {
			if state&(1<<i) != 0 {
				continue
			}

			if cur+i+1 >= desiredTotal {
				res = true
				break
			}

			if !dfs(cur+i+1, state|(1<<i)) {
				res = true
				break
			}

		}

		dp[state] = res
		return dp[state]
	}
	return dfs(0, 0)
}

// 暴力搜索
func canIWin2(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	vis := make([]bool, maxChoosableInteger+1)

	var dfs func(int) bool
	// dfs 表示当前累积和下， 选择是否能赢。
	dfs = func(cur int) bool {
		can := true
		for i := 1; i <= maxChoosableInteger; i++ {
			if vis[i] == true {
				continue
			}

			if cur+i >= desiredTotal {
				return true
			}

			vis[i] = true
			can = dfs(cur + i)
			vis[i] = false

			if !can {
				return true
			}
		}
		return false
	}

	return dfs(0)
}

// @lc code=end

