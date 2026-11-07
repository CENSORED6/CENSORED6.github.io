package main

import (
	"fmt"
)

type xy struct {
	p int 
	mode int 
}

func sw(x int) int {
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev
}

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	m := make([][]int, M+1)
	for i := range m {
		m[i] = make([]int, 4)
		for j := range m[i] {
			m[i][j] = -1
		}
	}

	q := []xy{{1, 1}}
	m[1][1] = 0 

	for len(q) > 0 {
		now := q[0]
		q = q[1:]

		x := now.p
		mode := now.mode
		steps := m[x][mode]

		if x == N {
			fmt.Println(steps)
			return
		}

		for next := 1; next <= 3; next++ {
			if next != mode && m[x][next] == -1 {
				m[x][next] = steps + 1
				q = append(q, xy{x, next})
			}
		}

		if mode == 1 {
			if x+1 <= M && m[x+1][1] == -1 {
				m[x+1][1] = steps + 1
				q = append(q, xy{x + 1, 1})
			}
		}else if mode == 2{
			if 2*x <= M && m[2*x][2] == -1 {
				m[2*x][2] = steps + 1
				q = append(q, xy{2 * x, 2})
			}
			if x%2 == 0 && x/2 >= 1 && m[x/2][2] == -1 {
				m[x/2][2] = steps + 1
				q = append(q, xy{x / 2, 2})
			}
		}else{
			rev := sw(x)
			if rev >= 1 && rev <= M && m[rev][3] == -1 {
				m[rev][3] = steps + 1
				q = append(q, xy{rev, 3})
			}
		}
	}
}

