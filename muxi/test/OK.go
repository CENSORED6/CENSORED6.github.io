package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {

    req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/hello", nil)

    if err != nil {
        panic(err)
    }


    client := &http.Client{}

    resp, err := client.Do(req)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()


    body, err := io.ReadAll(resp.Body)

    if err != nil {
        panic(err)
    }


    fmt.Println(string(body)) 

}