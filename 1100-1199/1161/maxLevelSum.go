/*
1161 - https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/description
*/
package main

import . "leetcode-go/testutil"

// DFS
func maxLevelSum(root *TreeNode) (result int) {
	rowSum := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(rowSum) == level {
			rowSum = append(rowSum, node.Val)
		} else {
			rowSum[level] += node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	for i, s := range rowSum {
		if s > rowSum[result] {
			result = i
		}
	}
	return result + 1
}

func main() {

}
