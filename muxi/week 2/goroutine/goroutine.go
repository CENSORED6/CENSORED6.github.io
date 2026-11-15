package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"time"
)

type Item struct {
	id int
	n  int
}

const n = 20

func main() {
	ch := make(chan Item)

	for i := 1; i <= 20; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.IntN(1001)) * time.Millisecond)
			ch <- Item{id: i, n: rand.IntN(100)}
		}(i)
	}

	items := make([]Item, 0, 20)
	for i := 0; i < 20; i++ {
		items = append(items, <-ch)
	}

	for _, it := range items {
		fmt.Printf("%d->%d\n", it.id, it.n)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].id < items[j].id
	})

	fmt.Println("\n")
	for _, it := range items {
		fmt.Printf("%d->%d\n", it.id, it.n)
	}
}
