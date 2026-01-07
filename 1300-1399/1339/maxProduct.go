/*
1339 - https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/description
*/
package main

import . "leetcode-go/testutil"

// DFS
func maxProduct(root *TreeNode) (result int) {
	var dfs1 func(*TreeNode) int
	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + dfs1(node.Left) + dfs1(node.Right)
	}
	total := dfs1(root)

	var dfs2 func(*TreeNode) int
	dfs2 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		s := node.Val + dfs2(node.Left) + dfs2(node.Right)
		result = max(result, s*(total-s))
		return s
	}
	dfs2(root)
	return result % 100_000_007
}

func main() {

}
