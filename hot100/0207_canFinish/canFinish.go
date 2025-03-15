package canfinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	// 构建邻接表以及统计节点入度
	for _, info := range prerequisites {
		// [a, b] b-->a
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	// 入度为0代表没有先修课程 作为起始节点进入队列
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		// 队首出队
		u := q[0]
		q = q[1:]
		// 出队元素保存下来 代表能满足条件的课程
		result = append(result, u)
		// 节点u出队 则和u邻接的节点的度减少1
		for _, v := range edges[u] {
			indeg[v]--
			// 只有入度为0才能入队 队列中维护的是能满足先修条件的课程
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
