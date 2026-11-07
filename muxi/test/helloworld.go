package main

import "fmt"

import "net/http"

func helloHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hi! Welcome to Muxi507!")
}

func main(){
	http.HandleFunc("/health",helloHandler)

	fmt.Println("启动成功")

	http.ListenAndServe(":8081",nil)

}