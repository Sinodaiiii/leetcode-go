package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree260607(descriptions [][]int) *TreeNode {
	nodeDict := map[int]*TreeNode{}
	inDict := map[*TreeNode]int{}
	for _, description := range descriptions {
		if nodeDict[description[0]] == nil {
			nodeDict[description[0]] = &TreeNode{description[0], nil, nil}
		}
		if nodeDict[description[1]] == nil {
			nodeDict[description[1]] = &TreeNode{description[1], nil, nil}
		}
		father, child := nodeDict[description[0]], nodeDict[description[1]]
		if description[2] == 1 {
			father.Left = child
		} else {
			father.Right = child
		}
		if _, exist := inDict[father]; !exist {
			inDict[father] = 0
		}
		inDict[child] += 1
	}
	for k, v := range inDict {
		if v == 0 {
			return k
		}
	}
	return nil
}
