package main

import "fmt"

func main(){

	x := 0 
	m := make(map[float64]int)
	for i := 2.0; i <= 1000 ; i++{
		j ,k := 1.0 , i-1
		if int(i) % 2 == 0{
			for j < i{
				var p float64
				p = k/j
				if m[p] == 0{
					m[p] = x
				}
				j++ 
				k--
				x++
				//fmt.Println(m[p],p,x)
			}
		}else{
			for j < i{
				var p float64
				p = j/k
				if m[p] == 0{
					m[p] = x
				}
				j++ 
				k--
				x++
				//fmt.Println(m[p],p,x)
			}
		}

	}

	//var n,r1,r2 float64
	//r1 = 1.0
	//r2 =3.0
	//n = r1 / r2

	//制表
	var a,b,c int

	fmt.Scan(&a,&b)

	c = m[float64(a)/float64(b)]
	fmt.Println(c/2)

}