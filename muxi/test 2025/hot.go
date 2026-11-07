package main

import "fmt"

func main(){
	var k int
	var num [20]int
	var t int
	var sum int
	sum = 0
	t = 0

	fmt.Scan(&k)

	for i := 1; i <= k;i ++ {
		fmt.Scan(&num[i])
	} 
	for i := 1; i+2 <= k;i ++ {
		if num[i] % 2 == 0 && num[i+1] % 2 == 1 && num[i+2] % 2 == 0 {
			t = num[i] + num[i+1] + num[i+2]
				if t > sum {
					sum = t
				}
		} 
	}

	if sum == 0 {
		sum = -1
	}

	fmt.Println(sum)

}