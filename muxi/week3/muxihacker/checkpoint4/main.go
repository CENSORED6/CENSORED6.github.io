package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"D:/Code/go/muxihakcer/test.go",
		httptool.FILE, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoi5Y2aIiwiaWF0IjoxNzk1MTM5NTQ3LCJuYmYiOjE3OTUxMzk1NDd9.OEOh2YNDNrDCJdi05_TikAxk9OhBbOxyAR2H5TRN8dI")
	resp, err := req.SendRequest() //resp
	if err != nil {
		fmt.Println(err)
	}
	resp.ShowBody()
	//resp.ShowHeader()

	// write your code below
	// ...
}

/*
1Map-Fragments : muxi
2Map-Fragments : backend
3Map-Fragments : computer
4Map-Fragments : examination
muxi/backend/computer/examination

Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
END
我就知道你能成功！Backend组织欢迎你！
*/
