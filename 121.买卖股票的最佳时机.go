/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mm := prices[0]
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-mm > res {
			res = prices[i] - mm
		}

		if mm > prices[i] {
			mm = prices[i]
		}
	}
	return res
}

// @lc code=end

