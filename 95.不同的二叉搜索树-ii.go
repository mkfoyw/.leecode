/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 二叉搜索树的性质是跟节点的值大于所有左子树的值， 小于右子树的值。
// 因此我们可以枚举根节点， 然后就把节点分为两部分， 然后递归解决
func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(left, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if left > right {
		res = append(res, nil)
		return res
	}

	if left == right {
		root := new(TreeNode)
		root.Val = left
		root.Left = nil
		root.Right = nil
		res = append(res, root)
		return res
	}

	for i := left; i <= right; i++ {
		lroot := dfs(left, i-1)
		rroot := dfs(i+1, right)
		for _, litem := range lroot {
			for _, ritem := range rroot {
				root := new(TreeNode)
				root.Val = i
				root.Left = litem
				root.Right = ritem
				res = append(res, root)
			}
		}
	}
	return res
}

// @lc code=end

