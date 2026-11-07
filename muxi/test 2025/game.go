package main

import (
    "fmt"
    "strings"
)

const q_SIZE = 8

type ParseResult struct {
    Currentq string   // 当前输入队列
    Output       []string // 输出技能
}

// 技能列表
var skillPatterns = map[string]string{
    "SDJ":   "波动拳",
    "DSDJ":  "升龙拳",
    "ASDL":  "防御反击",
    "SDJK":  "波动连击",
    "DSDJJ": "升龙连破",
    "WASJ":  "旋风拳",
    "ASDLK": "反击爆破",
    "SDJKL": "终极波动拳",
    "DJJJ":  "三连击",
}

func parseCommand(input string) ParseResult {
    result := ParseResult{
        Output: []string{},
    }

	var q,z []string
	var t int
	for i:=0; i < len(input); i++ {

		var c int
		c = int(input[i])

        if c == 'T'{
            q = z
            continue
        }

		if c=='W'||c=='A'||c=='S'||c=='D'{
			if c != t{
				q = append(q , string(c))
			}
		}else {q = append(q , string(c))}
		t = c
        if len(q) > q_SIZE {
			q = q[1:]
		}
	}

    var qw string
    qw = strings.Join(q,"")

    result.Currentq = qw

    var pat string

    for i:=0;i < len(qw); i++{
        for j,_ := range  skillPatterns{
            if strings.Contains(qw, j) && len(pat)<=len(j){
                pat = j
                qw = qw[i+1:]
            }
        }
    }

    r := []rune(skillPatterns[pat])
    for i := 0; i < len(r); i++ {
        result.Output = append( result.Output, string(r[i]))
    }


    return result
}

func main() {

    var input string

	fmt.Scan(&input)

    res := parseCommand(input)
    for i := 0 ; i < len(res.Output);i++{
	    fmt.Print(res.Output[i])
    }
    fmt.Println(res.Currentq)
    
}