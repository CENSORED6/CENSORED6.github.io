package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/encrypt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key",
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
	resp.ShowBody()
	fmt.Println("\n\n")
	resp.ShowHeader()
	//解码
	fmt.Println("\n\n\n")
	errorCode, err := encrypt.Base64Decode("c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(errorCode)

	//加密
	//fmt.Println(req)
	errorCode = "for {go func(){time.Sleep(1*time.Hour)}()}"
	secretKey := "MuxiStudio203304"
	data, err := encrypt.AESEncryptOutInBase64([]byte(errorCode), []byte(secretKey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	req1, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
		string(data),
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
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
	fmt.Println("\n\n")
	resp1.ShowHeader()

}

/*
c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==
secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()}
106 119 65 97 85 113 73 74 122 121 114 87 99 51 101 65 111 108 52 85 56 104 76 67 52 103 113 108 55 70 102 100 66 78 68 120 110 54 110 66 118 79 76 111 121 120 84 116 108 84 74 70 115 57 73 122 118 56 105 73 98 74 87 116
*/
