/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
// 解题思路
// 我们要戳n次气球
// 每次我们怎么选择我们要戳的气球呢？
// 有好的贪心算法吗？ 或者遍历所有的方案。
// 遍历所有的方案我们可以考虑记忆化搜索
// 或者有什么转化的方案吗？

/*
问题转化：

我们观察戳气球的操作，发现这会导致两个气球从不相邻变成相邻，使得后续操作难以处理。于是我们倒过来看这些操作，将全过程看作是每次添加一个气球。
dp[i][j]表示 表示将开区间(i,j)内的位置全部填满气球能够得到的最多硬币数
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(nums []int) int {
	if n := len(nums); n == 0 {
		return 0
	}

	val := []int{1}
	val = append(val, nums...)
	val = append(val, 1)
	n := len(val)

	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n, n)
	}

	for k := 1; k < n; k++ {
		for left := 0; left+k < n; left++ {
			right := left + k
			for mid := left + 1; mid < right; mid++ {
				dp[left][right] = max(dp[left][right], val[left]*val[mid]*val[right]+dp[left][mid]+dp[mid][right])
			}
		}
	}
	return dp[0][n-1]
}

func maxCoins2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	val := []int{1}
	val = append(val, nums...)
	val = append(val, 1)

	dp := make([][]int, len(val))
	for i := 0; i < len(val); i++ {
		dp[i] = make([]int, len(val))
	}

	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if left >= right {
			return 0
		}

		if dp[left][right] != 0 {
			return dp[left][right]
		}

		res := 0
		for i := left + 1; i < right; i++ {
			res = max(res, val[left]*val[i]*val[right]+dfs(left, i)+dfs(i, right))
		}

		dp[left][right] = res
		return res
	}

	dfs(0, len(val)-1)
	return dp[0][len(val)-1]
}

/*
暴力遍历：
	Time Limit Exceeded
	23/70 cases passed (N/A)
	Testcase
	[7,9,8,0,7,1,3,5,5,2,3]
	Expected Answer
	1654
*/

func maxCoins1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	flag := make([]bool, n)
	res := 0

	var dfs func(int, int)
	dfs = func(cur int, sum int) {
		if cur == n {
			res = max(res, sum)
		}

		for i := 0; i < n; i++ {
			if flag[i] {
				continue
			}

			flag[i] = true
			left := 1
			for j := i - 1; j >= 0; j-- {
				if !flag[j] {
					left = nums[j]
					break
				}
			}
			right := 1
			for j := i + 1; j < n; j++ {
				if !flag[j] {
					right = nums[j]
					break
				}
			}
			dfs(cur+1, sum+left*right*nums[i])

			flag[i] = false
		}
	}

	dfs(0, 0)
	return res

}

// @lc code=end


