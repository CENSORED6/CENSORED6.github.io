package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6ImNoYWxsZW5nZSIsImV4cCI6MTc5NTI0NDI2OSwiaWF0IjoxNzk1MjQyNDY5LCJqdGkiOiJmMDI3MTQ2MGFhYjI5Mjk2YTI4NGM5YWI2YTU1NmI5OCJ9.zHMdA2SDBjhi7-9hwG47zYQGcMx2eJZOJsn6U-qAeXA"
const url = "http://117.72.42.214:8080/api/challenge"

func x(m int, wg *sync.WaitGroup, ch chan<- string) {
	//defer wg.Done()    //?

	req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf(`{"seq":%d}`, m)))
	if err != nil {
		ch <- fmt.Sprintf("seq=%d err=%v", m, err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("seq=%d err=%v", m, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%s", body)
	//ch <- fmt.Sprintf("seq=%d body=%s", m , body)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go x(i, &wg, ch)
	}

	go func() { wg.Wait(); close(ch) }()

	for s := range ch {
		fmt.Println(s)
	}
}
