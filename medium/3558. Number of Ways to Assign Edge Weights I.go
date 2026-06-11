package medium

func assignEdgeWeights260611(edges [][]int) int {
	n := len(edges) + 1
	father := make([]int, n)
	up, down := make([]int, n), make([]int, n)
	for i := range n {
		father[i] = i
		up[i] = 1
		down[i] = 1
	}
	var find func(index int) int
	find = func(index int) int {
		if father[index] != index {
			tmp := father[index]
			father[index] = find(father[index])
			up[index] += up[tmp] - 1
		}
		return father[index]
	}
	union := func(x, y int) {
		root := find(y)
		father[x] = root
		down[y] = max(down[y], 2)
		up[x] = up[y] + 1
		down[root] = max(down[root], up[y]+down[x])
	}

	for _, edge := range edges {
		union(edge[1]-1, edge[0]-1)
	}
	maxDepth := 0
	for _, depth := range down {
		maxDepth = max(maxDepth, depth)
	}
	ans := 1
	for _ = range maxDepth - 2 {
		ans = ans * 2 % 1000000007
	}
	return ans
}
