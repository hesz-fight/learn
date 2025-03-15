package numisland

import "container/list"

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	conut := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				bfs(i, j, grid, visited)
				conut++
			}
		}
	}

	return conut
}

func bfs(i, j int, grid [][]byte, visited [][]bool) {
	m := len(grid)
	n := len(grid[0])
	visited[i][j] = true
	queue := list.New()
	queue.PushBack([2]int{i, j})
	for queue.Len() != 0 {
		lenQue := queue.Len()
		for i := 0; i < lenQue; i++ {
			point := (queue.Remove(queue.Front())).([2]int)
			x := point[0]
			y := point[1]
			if x-1 >= 0 && grid[x-1][y] == '1' && !visited[x-1][y] {
				queue.PushBack([2]int{x - 1, y})
				visited[x-1][y] = true
			}
			if x+1 <= m-1 && grid[x+1][y] == '1' && !visited[x+1][y] {
				queue.PushBack([2]int{x + 1, y})
				visited[x+1][y] = true
			}
			if y-1 >= 0 && grid[x][y-1] == '1' && !visited[x][y-1] {
				queue.PushBack([2]int{x, y - 1})
				visited[x][y-1] = true
			}
			if y+1 <= n-1 && grid[x][y+1] == '1' && !visited[x][y+1] {
				queue.PushBack([2]int{x, y + 1})
				visited[x][y+1] = true
			}
		}
	}
}
