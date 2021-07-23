/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return float64(0)
	}
	target := (len(nums1) + len(nums2) + 2) / 2

	i := -1
	j := -1

	for current := 0; current < target; current++ {
		if i+1 < len(nums1) && j+1 < len(nums2) {
			if nums1[i+1] < nums2[j+1] {
				i++
			} else {
				j++
			}
		} else if i+1 < len(nums1) {
			i++
		} else {
			j++
		}
	}

	// 只有一个中位数的情况下
	if (len(nums1)+len(nums2))%2 == 1 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				return float64(nums1[i])
			} else {
				return float64(nums2[j])
			}
		} else if i >= 0 {
			return float64(nums1[i])
		} else {
			return float64(nums2[j])
		}
	} else {
		//在有两个中位数的情况下
		res := 0

		//找到第一个中位数
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				res += nums1[i]
				i--
			} else {
				res += nums2[j]
				j--
			}
		} else if i >= 0 {
			res += nums1[i]
			i--
		} else {
			res += nums2[j]
			j--
		}

		//找到第二个中位数
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				res += nums1[i]
			} else {
				res += nums2[j]
			}
		} else if i >= 0 {
			res += nums1[i]
		} else {
			res += nums2[j]
		}

		return float64(res) / 2
	}

}

// @lc code=end
