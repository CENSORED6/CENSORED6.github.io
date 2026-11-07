package main

import (
	"fmt"
)

// 定义状态结构体
type xy struct {
	p int // 当前房间号
	mode int // 当前模式 (1, 2, 3)
}

// 翻转数字的函数 (例如 120 -> 21)
func reverseNum(x int) int {
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev
}

func main() {
	var M, N int
	if _, err := fmt.Scan(&M, &N); err != nil {
		return
	}

	// m[p][mode] 记录到达该状态的最少步数
	// 初始化为 -1 表示未访问过
	m := make([][]int, M+1)
	for i := range m {
		m[i] = make([]int, 4) // 模式只有 1, 2, 3
		for j := range m[i] {
			m[i][j] = -1
		}
	}

	// 队列用于 BFS
	q := []xy{{1, 1}}
	m[1][1] = 0 // 初始状态：房间 1，模式 1，步数 0

	for len(q) > 0 {
		// 出队
		curr := q[0]
		q = q[1:]

		x := curr.p
		mode := curr.mode
		steps := m[x][mode]

		// 如果到达了目标房间 N，直接输出步数并结束程序
		if x == N {
			fmt.Println(steps)
			return
		}

		// 1. 切换模式 (切换到其他两种模式中的任意一种)
		for next := 1; next <= 3; next++ {
			if next != mode && m[x][next] == -1 {
				m[x][next] = steps + 1
				q = append(q, xy{x, next})
			}
		}

		// 2. 空间迁移操作 (根据当前模式)
		switch mode {
		case 1:
			// 模式一：迁移 (x -> x + 1)
			if x+1 <= M && m[x+1][1] == -1 {
				m[x+1][1] = steps + 1
				q = append(q, xy{x + 1, 1})
			}
		case 2:
			// 模式二：跃迁
			// 前往 2x
			if 2*x <= M && m[2*x][2] == -1 {
				m[2*x][2] = steps + 1
				q = append(q, xy{2 * x, 2})
			}
			// 前往 x/2 (前提是 x 是偶数且 >= 2)
			if x%2 == 0 && x/2 >= 1 && m[x/2][2] == -1 {
				m[x/2][2] = steps + 1
				q = append(q, xy{x / 2, 2})
			}
		case 3:
			// 模式三：逆转
			rev := reverseNum(x)
			// 逆转后的数字必须在 1 到 M 之间
			if rev >= 1 && rev <= M && m[rev][3] == -1 {
				m[rev][3] = steps + 1
				q = append(q, xy{rev, 3})
			}
		}
	}
}
