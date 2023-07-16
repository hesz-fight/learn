package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := NewQueue()
	queue.EnQueue(0)
	queue.EnQueue(1)
	queue.EnQueue(2)
	fmt.Println(queue.getSize())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.getSize())
}

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	return &Queue{data: list.New()}
}

func (q *Queue) EnQueue(v interface{}) {
	q.data.PushBack(v)
}

func (q *Queue) DeQueue() interface{} {
	if q.data.Front() == nil {
		return nil
	}
	return q.data.Remove(q.data.Front())
}

func (q *Queue) getSize() int {
	return q.data.Len()
}

func (q *Queue) isEmpty() bool {
	return q.getSize() == 0
}
