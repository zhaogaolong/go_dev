package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type student struct {
	name string
	age  int
	id   string
}

type studentArray []student

func (p studentArray) Len() int {
	return len(p)
}

func (p studentArray) Less(i, j int) bool {
	// return p[i].age < p[j].age
	return p[i].id < p[j].id
}

func (p studentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// sort 接口排序
func main() {

	var stus studentArray

	for i := 0; i < 10; i++ {
		stu := student{
			name: fmt.Sprintf("stud%d", rand.Intn(100)),
			id:   fmt.Sprintf("110%d", rand.Int()),
			age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println("\n\n")
	sort.Sort(stus)

	for _, v := range stus {
		fmt.Println(v)
	}
}
