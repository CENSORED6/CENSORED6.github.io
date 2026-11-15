package main

import (
	"fmt"
)

func main(){
	/*var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[5:7]
	//fmt.Println(a)
	for i := 0 ; i < len(a);i++{
		fmt.Println(a[i])
	}
	fmt.Println("\n")
	fmt.Println(cap(a))
	a=a[1:4]
	fmt.Println(cap(a))
	fmt.Println("\n")
	for i := 0 ; i < len(a);i++{
		fmt.Println(a[i])
	}*/
	//fmt.Println(a)
	/*for i := 0 ; i < len(items);i++{
		fmt.Println(items[i])
	}*/
	var n,m int
	var h string
	h = "hello,世界"
	n=len(h)
	for i := range h{
		m++
		i=i
	}
	fmt.Println(n,"\n",m)//12 8
	/*s := make([]byte, 5)
	var m,n int
	m=cap(s)
	n=len(s)
	fmt.Println(n,"\n",m)*///5 5 2 3
}