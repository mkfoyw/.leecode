/*
 * @lc app=leetcode.cn id=600 lang=golang
 *
 * [600] 不含连续1的非负整数
 */

// @lc code=start

func findIntegers(n int) int {

	//dp[t]： 表示高度为t-1，根节点为0满二叉树中不包含连续的1的路径个数。
	dp := [32]int{1, 1}
	for i := 2; i < 32; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	pre := 0
	ans := 0
	for i := 29; i >= 0; i-- {
		val := 1 << i
		//判断当前是否有右子树
		if n&val > 0 {
			//加上左子树
			ans += dp[i+1]
			if pre == 1 {
				break
			}
			// 标记前一位为1
			pre = 1
		} else {
			pre = 0
		}

		// 加上一条全为0的路径
		if i == 0 {
			ans++
		}

	}
	return ans
}

// @lc code=end

