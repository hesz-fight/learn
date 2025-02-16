package main

import "fmt"

func main() {
	a := []int{1}
	fmt.Println(a)
	xxx(a)
	fmt.Println(a)
}

func xxx(a []int) {
	a = append(a, 2)
	fmt.Println(a)
}
