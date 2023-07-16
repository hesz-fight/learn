package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := list.New()
	list.PushBack(0)
	list.PushFront(-1) // 链表头部插入元素
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(-1)

	fmt.Println(list.Remove(list.Front())) // 移除链表头部的元素
	fmt.Println(list.Remove(list.Back()))  // 移除链表末尾的元素

	// 从前往后遍历
	for it := list.Front(); it != nil; it = it.Next() {
		fmt.Println(it.Value)
	}
	// 从后往前遍历
	for it := list.Back(); it != nil; it = it.Prev() {
		fmt.Println(it.Value)
	}

	list.Init() // 清空数据
}
