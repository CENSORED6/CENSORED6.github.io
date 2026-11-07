package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
)

type Request struct {
    Op     string `json:"op"`
    Client string `json:"client"`
    Seq    int    `json:"seq"`
    Data   string `json:"data"`
}

type mail struct {
	count string
	seq []int
	in map[int]string
}
// TODO
var m mail

type Server struct{
   /* Code     string `json:"code"`
    Count string `json:"count"`
    Seq    int    `json:"seq"`
    Next_seq   string `json:"next_seq"`
	State	string  `json:"state"`*/
}

func NewServer() *Server {
	m.in = make(map[int]string)
    return &Server{
            
    }
}
var i int
var state bool
// TODO 实现小柠猫对四种指令的响应
// 返回的 map 会自动被编码为 JSON
func (s *Server) Handle(req Request) map[string]any {
    if req.Op == "MEOW" {
		if m.count == ""{
			m.count = req.Client
			return map[string]any{"code":"PURR"}
		}else {return map[string]any{"code":"ERR_ALREADY_EXISTS"}}
	}
	if req.Op == "SEND" && state == false{
		if m.count != ""{
			if (m.in[req.Seq] == ""&&i+1 == req.Seq)||m.in[req.Seq] == req.Data {
				if m.in[req.Seq] == ""{
					i++
				}
				m.in[req.Seq] = req.Data
				return map[string]any{"code":"ACK","seq":req.Seq}
				

			}else if m.in[req.Seq] != req.Data&&i == req.Seq{
				return map[string]any{"code":"ERR_CONFLICT"}
			}else{
				return map[string]any{"code":"ERR_OUT_OF_ORDER"}
			}
		}else {return map[string]any{"code":"ERR_NO_SESSION"}}
	}
	
	if req.Op == "PEEK"{
		if state == false{
			return map[string]any{"code":"PURR","state":"AWAKE","count":i,"next_seq":i+1}
		}else{return map[string]any{"code":"PURR","state":"ASLEEP","count":i,"next_seq":i+1}}
	}
		
	if req.Op == "SEAL"{
		state = true
    	return map[string]any{"code":"PURR"}
	}
	return map[string]any{"code":"ERR_ASLEEP"}
}
//只写了一个邮箱的T_T
func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    server := NewServer()

    for {
        var req Request
        if err := dec.Decode(&req); err != nil {
            if err == io.EOF {
                return
            }
            fmt.Fprintln(os.Stderr, "读取请求失败:", err)
            return
        }
        if err := enc.Encode(server.Handle(req)); err != nil {
            fmt.Fprintln(os.Stderr, "输出响应失败:", err)
            return
        }
    }
}