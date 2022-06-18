package main

import (
	"fmt"
	"sort"
)

type singer struct {
	first string
	age   int
}

type ByAge []singer

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByName []singer

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].first < bn[j].first
}

func main() {
	s1 := singer{"Momo", 25}
	s2 := singer{"Sana", 25}
	s3 := singer{"Dahyun", 24}
	s4 := singer{"Mina", 25}
	s5 := singer{"Nancy", 22}

	singers := []singer{s1, s2, s3, s4, s5}
	fmt.Println(singers)
	sort.Sort(ByAge(singers))
	fmt.Println(singers)

	sort.Sort(ByName(singers))
	fmt.Println(singers)
}
