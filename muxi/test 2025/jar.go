package main
 
import "fmt"
 
type node struct {
    pos    int
    energy int
    jumps  int
}
 
func main() {
    var n, E, r int
    var road string
    fmt.Scan(&n, &E, &r)
    fmt.Scan(&road)
 
    if n <= 1 {
        fmt.Println(0)
        return
    }
 
    best := make([]int, n)
    for i := range best {
        best[i] = -1
    }
    best[0] = E
 
    q := []node{{0, E, 0}}
    for head := 0; head < len(q); head++ {
        cur := q[head]
        if cur.pos == n-1 {
            fmt.Println(cur.jumps)
            return
        }
 
        maxStep := cur.energy
        if maxStep > n-1-cur.pos {
            maxStep = n - 1 - cur.pos
        }
        for k := 1; k <= maxStep; k++ {
            np := cur.pos + k
            if road[np] == '-' {
                continue
            }
            ne := cur.energy
            if road[np] == '+' {
                ne += r
            }
            if ne > best[np] {
                best[np] = ne
                q = append(q, node{np, ne, cur.jumps + 1})
            }
        }
    }
    fmt.Println(-1)
}