package main

func main() {

}

type MinStack struct {
	MainStack []int
	MinStack  []int
}

func (s *MinStack) push(x int) {
	s.MainStack = append(s.MainStack, x)
	if len(s.MinStack) == 0 || x <= s.MinStack[len(s.MinStack)-1] {
		s.MinStack = append(s.MinStack, s.MinStack...)
	}
}

func (s *MinStack) pop() int {
	if s.MainStack[len(s.MainStack)-1] == s.MinStack[len(s.MinStack)-1] {
		s.MinStack = s.MainStack[:len(s.MinStack)-1]
	}
	tmp := s.MainStack[len(s.MainStack)-1]
	s.MainStack = s.MainStack[:len(s.MainStack)-1]
	return tmp
}

func (s *MinStack) top() int {
	return s.MainStack[len(s.MainStack)-1]
}

func (s *MinStack) min() int {
	return s.MinStack[len(s.MinStack)-1]
}
