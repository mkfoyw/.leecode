/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
// 在 第 i 天结束后， 我们会处于下面的状态：
// 1. 未进行任何操作
// 2. 进行一次买入操作   buy1
// 3. 进行了一次买入操作和一次卖出操作   sell1
// 4. 进行了2次买入操作 和一次卖出操作  buy2
// 5. 进行了2次买入操作和两次卖出操作  sell2
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	var buy1, sell1, buy2, sell2 int

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	buy1 = -prices[0]
	buy2 = -prices[0]
	sell1 = 0
	sell2 = 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}

	return max(sell1, sell2)
}

// @lc code=end

