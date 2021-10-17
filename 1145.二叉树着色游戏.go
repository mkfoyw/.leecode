/*
 * @lc app=leetcode.cn id=1145 lang=golang
 *
 * [1145] 二叉树着色游戏
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
// 解题思路：
// 二号玩家想要图尽可能图更多的节点， 那么它选择图色的节点可能有 x 左右子树以及x的父亲节点。

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	sum := make([]int, n+1)

	// 节点 x 的父亲节点
	px := root
	// 统计每棵子树的节点个数
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}

		//找到节点x的父亲节点
		if (cur.Left != nil && cur.Left.Val == x) || (cur.Right != nil && cur.Right.Val == x) {
			px = cur
		}

		sum[cur.Val] = 1

		sum[cur.Val] += dfs(cur.Left)
		sum[cur.Val] += dfs(cur.Right)

		return sum[cur.Val]
	}

	dfs(root)

	// 如果 x 节点为根节点， 那么不存在父亲节点
	if px.Val != x {
		if n-sum[x] > sum[x] {
			return true
		}
	}

	// 找到 x 节点
	var curx *TreeNode
	if px.Val == x {
		curx = px
	} else if px.Left != nil && px.Left.Val == x {
		curx = px.Left
	} else if px.Right != nil && px.Right.Val == x {
		curx = px.Right
	}

	//选择 x 的左节点
	if curx.Left != nil {
		if sum[curx.Left.Val] > n-sum[curx.Left.Val] {
			return true
		}
	}

	//选择 x 的右节点
	if curx.Right != nil {
		if sum[curx.Right.Val] > n-sum[curx.Right.Val] {
			return true
		}
	}
	return false
}

// @lc code=end

