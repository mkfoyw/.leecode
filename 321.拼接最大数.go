

/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */

// @lc code=start

// x 表示从nums1选取元素数量
// y 表示从nums2选取元素数量
// 那么一定有 x + y = k
// 现在的问题就变为如何从数组中选取元素， 以及选取元素如何进行合并， 以及合并后， 然后判断大小

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	mm := min(len(nums1), k)
	for i := 0; i <= mm; i++ {
		t1 := selectX(nums1, i)
		t2 := selectX(nums2, k-i)
		t3 := merge(t1, t2)
		if len(t3) != k {
			continue
		}
		res = maxRes(res, t3)
	}
	return res
}

func selectX(nums []int, x int) []int {
	if len(nums) < x || x <= 0 {
		return []int{}
	}

	stack := []int{}
	for i := 0; i < len(nums); i++ {
		//不会添加的情况
		if len(stack) == x && stack[len(stack)-1] >= nums[i] {
			continue
		}

		//需要添加的情况
		for len(stack) != 0 && len(stack)+len(nums[i:]) > x && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])

	}
	return stack
}

func maxRes(a []int, b []int) []int {
	na, nb := len(a), len(b)
	if na == 0 {
		return b
	}
	if nb == 0 {
		return a
	}

	for i := 0; i < na; i++ {
		if a[i] > b[i] {
			return a
		}

		if b[i] > a[i] {
			return b
		}
	}
	return a

}

// 在进行合并的时候， 如果两个数组的第一个元素不同，则选择最大的。
// 如果相同， 就会比较下一个元素， 知道两个元素不同。
// 这也就是选择字典顺序最大的。
func lexicCompare(a []int, b []int) bool {
	na, nb := len(a), len(b)

	i, j := 0, 0
	for i < na && j < nb {
		if a[i] > b[j] {
			return true
		}
		if a[i] < b[j] {
			return false
		}
		i++
		j++
	}

	if na > nb {
		return true
	} else {
		return false
	}
}

func merge(a []int, b []int) []int {
	na, nb := len(a), len(b)
	if na == 0 {
		return b
	}
	if nb == 0 {
		return a
	}

	res := []int{}
	i, j := 0, 0
	for i < na || j < nb {
		if lexicCompare(a[i:], b[j:]) {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
