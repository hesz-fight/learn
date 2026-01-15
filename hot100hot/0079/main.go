package main

var (
	board0 [][]byte
	word0  string
	m, n   int
)

func exist(board [][]byte, word string) bool {
	board0 = board
	word0 = word
	m, n = len(board), len(board[0])
	visited := [][]bool{}
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helper(i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

func helper(i int, j int, k int, visited [][]bool) bool {
	if board0[i][j] != word0[k] {
		return false
	}
	if k == len(word0)-1 {
		return true
	}

	if i-1 >= 0 && !visited[i-1][j] {
		visited[i][j] = true
		if helper(i-1, j, k+1, visited) {
			return true
		}
		visited[i][j] = false
	}
	if i+1 < m && !visited[i+1][j] {
		visited[i][j] = true
		if helper(i+1, j, k+1, visited) {
			return true
		}
		visited[i][j] = false
	}
	if j-1 >= 0 && !visited[i][j-1] {
		visited[i][j] = true
		if helper(i, j-1, k+1, visited) {
			return true
		}
		visited[i][j] = false
	}
	if j+1 < n && !visited[i][j+1] {
		visited[i][j] = true
		if helper(i, j+1, k+1, visited) {
			return true
		}
		visited[i][j] = false
	}

	return false
}
