package main

import (
	"fmt"
)

func main(){
	var n,m int
	var h string
	h = "hello,世界"
	n=len(h)
	for i := range h{
		m++
		i=i
	}
	fmt.Println(n,"\n",m)//len按字节算“世界”这两个玩意每个三字节，其他每个一个 12 8
	/*s := make([]byte, 5)
	var m,n int
	m=cap(s)
	n=len(s)
	s = s[2:4]
	fmt.Println(n,"\n",m)*///5 5 2 3 cap算的总内存 len当前占用  
}