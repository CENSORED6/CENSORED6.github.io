package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
	TotalPages int `json:"total_pages"`
}

type Post struct {
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
}

func main() {

	var username, password string
	fmt.Scan(&username, &password)
	
	loginJson, err := json.Marshal(LoginRequest{username, password})
	if err != nil {
		panic(err)
	}

	loginRequest, err := http.NewRequest(
		"POST",
		"https://demosite1.muxixyz.com/api/auth/login",
		bytes.NewBuffer(loginJson),
	)
	if err != nil {
		panic(err)
	}

	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	if err != nil {
		panic(err)
	}

	loginResponse, err := client.Do(loginRequest)
	if err != nil {
		panic(err)
	}
	defer loginResponse.Body.Close() 

	pages := 1

	for i := 1; i <= pages; i++ {	

		url := fmt.Sprintf("https://demosite1.muxixyz.com/api/posts?page=%d", i)

		postsRequest, err := http.NewRequest(
			"GET",
			url, 
			nil,
		)

		postsResponse, err := client.Do(postsRequest)
		if err != nil {
			panic(err)
		}
		defer postsResponse.Body.Close() 
		

		body, err := io.ReadAll(postsResponse.Body)
		if err != nil {
			panic(err)
		}

		var posts PostsResponse
		json.Unmarshal(body, &posts)
		
		pages = posts.TotalPages

		fmt.Println("文章：",i)

		for _, post := range posts.Posts {
			
			fmt.Println("标题: ",post.Title)
			fmt.Println("摘要: ",post.Excerpt)
		}
	}
}