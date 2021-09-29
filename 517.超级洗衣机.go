/*
 * @lc app=leetcode.cn id=517 lang=golang
 *
 * [517] 超级洗衣机
 */

// @lc code=start
func findMinMoves(machines []int) int {
	n := len(machines)
	if n <= 1 {
		return 0
	}

	tot := 0
	for i := 0; i < n; i++ {
		tot += machines[i]
	}
	if tot%n != 0 {
		return -1
	}

	avg := tot / n

	curSum := 0
	res := 0
	for i := 0; i < n; i++ {
		curSum += machines[i] - avg
		res = max(res, max(abs(curSum), machines[i]-avg))
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

