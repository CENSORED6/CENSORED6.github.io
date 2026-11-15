package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name   string
	Scores []int
}

func (s *Student) top(n int) []int {
	if n <= 0 || s == nil || len(s.Scores) == 0 {
		return nil
	}
	t := make([]int, len(s.Scores))
	copy(t, s.Scores)
	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})
	if n > len(s.Scores) {
		n = len(s.Scores)
	}
	return t[:n]
}

func main() {
	var s Student
	s.Scores = []int{2, 2, 3, 9, 5, 6}
	fmt.Println(s.top(3))
}
