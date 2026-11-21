package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoi5Y2aIiwiaWF0IjoxNzk1MTM5NTQ3LCJuYmYiOjE3OTUxMzk1NDd9.OEOh2YNDNrDCJdi05_TikAxk9OhBbOxyAR2H5TRN8dI")
	resp, err := req.SendRequest() //resp
	if err != nil {
		fmt.Println(err)
	}
	//resp.ShowBody()
	resp.ShowHeader()
	//下载
	/*err1 := resp.Save("D:/out.jpg")
	if err1 != nil {
		fmt.Println(err1)
	}*/
	//攻破
	req1, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"D:/out.jpg",
		httptool.FILE, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	req1.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoi5Y2aIiwiaWF0IjoxNzk1MTM5NTQ3LCJuYmYiOjE3OTUxMzk1NDd9.OEOh2YNDNrDCJdi05_TikAxk9OhBbOxyAR2H5TRN8dI")
	resp1, err := req1.SendRequest() //resp
	if err != nil {
		fmt.Println(err)
	}
	resp1.ShowBody()
	fmt.Println("\n\n\n\n")
	resp1.ShowHeader()

	/*req2, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	req2.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoi5Y2aIiwiaWF0IjoxNzk1MTM5NTQ3LCJuYmYiOjE3OTUxMzk1NDd9.OEOh2YNDNrDCJdi05_TikAxk9OhBbOxyAR2H5TRN8dI")
	resp2, err := req2.SendRequest() //resp
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n\n\n")
	resp2.ShowHeader()*/
	// write your code below
	// ...http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate
}
