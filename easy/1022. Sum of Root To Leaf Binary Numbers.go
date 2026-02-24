package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf260224(root *TreeNode) int {
	ans := 0
	curr := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		curr = curr<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += curr
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		curr = curr >> 1
	}

	dfs(root)
	return ans
}
