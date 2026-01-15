package main

type MinStask struct {
	MainSlice []int
	MinSlice  []int
}

func (m *MinStask) push(val int) {
	m.MainSlice = append(m.MainSlice, val)
	if len(m.MinSlice) <= 0 {
		m.MinSlice = append(m.MinSlice, val)
		return
	}
	if val <= m.MinSlice[len(m.MinSlice)-1] {
		m.MinSlice = append(m.MinSlice, val)
	}
}

func (m *MinStask) pop() {
	topNum := m.MainSlice[len(m.MainSlice)-1]
	// 栈顶元素出栈
	m.MainSlice = m.MainSlice[:len(m.MainSlice)-1]
	if topNum == m.MinSlice[len(m.MinSlice)-1] {
		// 最小元素出栈
		m.MinSlice = m.MinSlice[:len(m.MinSlice)-1]
	}
}

func (m *MinStask) top() int {
	return m.MainSlice[len(m.MainSlice)-1]
}

func (m *MinStask) getMin() int {
	return m.MinSlice[len(m.MinSlice)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
