package minstack

type MinStack struct {
	mainStack   []int
	assistStack []int
}

func Constructor() MinStack {
	return MinStack{
		mainStack:   []int{},
		assistStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.mainStack = append(this.mainStack, val)
	if len(this.assistStack) == 0 || val <= this.assistStack[len(this.assistStack)-1] {
		this.assistStack = append(this.assistStack, val)
	}
}

func (this *MinStack) Pop() {
	rtn := this.mainStack[len(this.mainStack)-1]
	if rtn == this.assistStack[len(this.assistStack)-1] {
		this.assistStack = this.assistStack[:len(this.assistStack)-1]
	}
	this.mainStack = this.mainStack[:len(this.mainStack)-1]
}

func (this *MinStack) Top() int {
	return this.mainStack[len(this.mainStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.assistStack[len(this.assistStack)-1]
}
