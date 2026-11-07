package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
   "time"
)

func x(m int)int {

   url := "http://117.72.42.214:8080/api/challenge"
   method := "POST"
	
   var n []string
	n = append(n,"{")
	n = append(n,"seq: ")
   n = append(n,string(m+48))
   n = append(n,"}")
   var qw string
    qw = strings.Join(n,"")
	fmt.Println(m,qw)
   payload := strings.NewReader(qw)
   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6ImNoYWxsZW5nZSIsImV4cCI6MTc5MjkwOTE2MCwiaWF0IjoxNzkyOTA3MzYwLCJqdGkiOiJhNGI2MmU3ODFkMzEwYWZmNDg4NDliODk2ZDRhMDY2YSJ9.MPMdjU_Q_j7ieoiDU3-j8dwomJVkCTVRKBqQI7Zu8iw")
   req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
   req.Header.Add("Content-Type", "application/json")
   req.Header.Add("Accept", "*/*")
   req.Header.Add("Host", "117.72.42.214:8080")
   req.Header.Add("Connection", "keep-alive")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
   return 1
}
func main(){
	for i := 1; i < 11; i++ {

}
