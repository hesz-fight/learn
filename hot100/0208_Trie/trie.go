package trie

type Trie struct {
	Root *Node
}

type Node struct {
	EndFlag bool
	Next    [26]*Node
}

func Constructor() Trie {
	return Trie{Root: &Node{}}
}

func (this *Trie) Insert(word string) {
	chs := []byte(word)
	root := this.Root
	for _, ch := range chs {
		idx := ch - 'a'
		if root.Next[idx] == nil {
			root.Next[idx] = &Node{}
		}
		root = root.Next[idx]
	}
	root.EndFlag = true
}

func (this *Trie) Search(word string) bool {
	chs := []byte(word)
	root := this.Root
	for _, ch := range chs {
		idx := ch - 'a'
		if root.Next[idx] == nil {
			return false
		}
		root = root.Next[idx]
	}
	return root.EndFlag
}

func (this *Trie) StartsWith(prefix string) bool {
	chs := []byte(prefix)
	root := this.Root
	for _, ch := range chs {
		idx := ch - 'a'
		if root.Next[idx] == nil {
			return false
		}
		root = root.Next[idx]
	}
	return true
}
