/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

//dfs
func cloneGraph1(node *Node) *Node {
	visited := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}

		if _, ok := visited[node]; ok {
			return visited[node]
		}

		cloneNode := &Node{Val: node.Val, Neighbors: []*Node{}}
		visited[node] = cloneNode

		for _, neighber := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(neighber))
		}
		return cloneNode

	}
	return dfs(node)
}

//bfs
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := map[*Node]*Node{}
	queue := []*Node{node}

	visited[node] = &Node{Val: node.Val, Neighbors: []*Node{}}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, neighber := range cur.Neighbors {
			if _, ok := visited[neighber]; !ok {
				// 没有克隆
				visited[neighber] = &Node{Val: neighber.Val, Neighbors: []*Node{}}
				queue = append(queue, neighber)
			}
			visited[cur].Neighbors = append(visited[cur].Neighbors, visited[neighber])
		}
	}
	return visited[node]
}

// @lc code=end

