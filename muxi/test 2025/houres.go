package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0 ; i < n ; i ++{
		fmt.Scan(&a[i])
	}
	for i := 0 ; i < n ; i ++{
		fmt.Scan(&b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	m , t := 0 , n-1
	i , j := 0 , n-1
	money := 0


	for k := 0 ; k < n ; k++ {
		if a[t] > b[j]{
			money += 200
			t--
			j--
		}else if a[m] > b[i] {
			money += 200
			m++
			i++
		}else { 
			if a[m] < b[j] {
				money -= 200
			}
			m++
			j--
		}

	}
	fmt.Println(money)


}