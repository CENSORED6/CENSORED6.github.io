package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hi! Welcome to Muxi507!")
}

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func writeJSON(w http.ResponseWriter, httpStatus int, body any) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(httpStatus)
    //TODO
}  

func OK(w http.ResponseWriter, data any) {
    writeJSON(w, http.StatusOK, Response{
        Code: 200,
        Msg:  "success",
        Data: data,
    })
}

func Fail(w http.ResponseWriter, httpStatus, code int, msg string) {
    writeJSON(w, httpStatus, Response{
        Code: code,
        Msg:  msg,
        Data: nil,
    })
}

  
func DecodeRequest() {
    //TODO
}

func main(){

	http.HandleFunc("/health",helloHandler)

	fmt.Println("启动成功")

	http.ListenAndServe(":8081",nil)
	

}