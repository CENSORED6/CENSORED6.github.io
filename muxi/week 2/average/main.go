package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Scores []int
}

func (s *Student) average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	sum := 0
	for _, score := range s.Scores {
		sum += score
	}
	return float64(sum) / float64(len(s.Scores))
}

func main() {
	var s Student
	s.Scores = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s.average())
}
