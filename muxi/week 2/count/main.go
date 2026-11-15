package main

import (
	"fmt"
)

func add() func() int {
	n := 0
	return func() int {
		n++
		return n
	}

}
func main() {
	m := add()
	fmt.Println(m())
}
