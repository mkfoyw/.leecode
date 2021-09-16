/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start

//最低有效位
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

// 奇偶性
func countBits4(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}

// 最高有效位
func countBits3(n int) []int {
	res := make([]int, n+1)
	hightBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			hightBit = i
		}
		res[i] = res[i-hightBit] + 1
	}
	return res
}

func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = bitCount(i)
	}
	return res
}

// 统计一个整数中1的位数
func bitCount(n int) int {
	res := 0
	for n > 0 {
		n = n & (n - 1)
		res++
	}
	return res
}

// @lc code=end

