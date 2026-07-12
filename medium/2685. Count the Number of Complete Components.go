package medium

func countCompleteComponents260711(n int, edges [][]int) int {
	father := make([]int, n)
	cNodes := make([]int, n)
	cEdges := make([]int, n)
	for i := range n {
		father[i] = i
		cNodes[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if father[x] != x {
			father[x] = find(father[x])
		}
		return father[x]
	}
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			if fx < fy {
				father[fy] = fx
				cNodes[fx] += cNodes[fy]
				cEdges[fx] += cEdges[fy] + 1
			} else {
				father[fx] = fy
				cNodes[fy] += cNodes[fx]
				cEdges[fy] += cEdges[fx] + 1
			}
		} else {
			cEdges[fx] += 1
		}

	}

	ans := 0
	for _, edge := range edges {
		union(edge[0], edge[1])
	}
	for i := range n {
		if father[i] == i {
			e := cNodes[i] - 1
			e = (1 + e) * e / 2
			if cEdges[i] == e {
				ans += 1
			}
		}
	}
	return ans
}
