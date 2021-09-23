/*
 * @lc app=leetcode.cn id=472 lang=golang
 *
 * [472] 连接词
 */

// @lc code=start

const TREE_WIDTH int = 26

type Trie struct {
	Next  [TREE_WIDTH]*Trie_Node
	IsEnd bool
}

func Intert(root *Trie, s string) {
	if root == nil {
		root = new(Trie)
	}

	cur := root
	for i := 0; i < n; i++ {
		if cur.Nect[s[i]-'a'] == nil {
			t := new(Trie)
			cur.Next[s[i]-'a'] = t
			cur = t
		} else {
			cur = cur.Next[s[i]-'a']
		}
	}
	cur.IsEnd = true
}

func findAllConcatenatedWordsInADict(words []string) []string {

}

// @lc code=end

