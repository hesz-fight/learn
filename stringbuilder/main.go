package main

import "fmt"

type People struct {
	Name string
	Age  int
	Job  string
}

func NewPeople() *People {
	return &People{}
}

func (p *People) SetName(name string) *People {
	p.Name = name
	return p
}

func (p *People) SetAge(age int) *People {
	p.Age = age
	return p
}

func (p *People) SetJob(job string) *People {
	p.Job = job
	return p
}

func main() {
	p := NewPeople()
	p.SetName("Mary").SetAge(20).SetJob("teacher")
	fmt.Println(p)
}
