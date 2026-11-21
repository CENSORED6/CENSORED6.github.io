package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	//defer req.Body.Close()
	fmt.Println(req)
	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}
	resp.ShowHeader()

	/*data, err := resp.GetHeader("Passport")
	if err != nil {
		fmt.Println(err)
	}
	request.AddHeader("passport", data)
	fmt.Println(data)*/
	//data1 = string(data)
	/*errorCode, err := encrypt.Base64Decode("eyJjb2RlIjoi5Y2aIiwiaWF0IjoxNzk1MDgwMDgyLCJuYmYiOjE3OTUwODAwODJ9")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(errorCode)
	req.AddHeader("passport", errorCode)
	fmt.Println("\n")*/
	resp.ShowBody()
}

/*

{
  "code": "博",
  "iat": 1795080082,
  "nbf": 1795080082
}
*/
