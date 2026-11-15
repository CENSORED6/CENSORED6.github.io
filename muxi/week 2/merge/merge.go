package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Scores []int
}

func (s1 *Student) merge(s2 Student) Student {
	s := len(s1.Scores) + len(s2.Scores)
	merged := make([]int, 0, s)
	merged = append(merged, s1.Scores...)
	merged = append(merged, s2.Scores...)

	return Student{Name: s1.Name, Scores: merged}
}

func main() {
	var s1 Student
	var s2 Student
	var s Student
	s1.Scores = []int{1, 2, 3, 4, 5, 6}
	s2.Scores = []int{1, 2, 3, 4, 5, 6}
	s = s1.merge(s2)
	fmt.Println(s)
}
