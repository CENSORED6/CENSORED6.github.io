package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main(){
	inputFile, err := ioutil.ReadFile("text.txt")
    if err != nil {
        panic(err)
    }

	

	fmt.Printf("%s\n", string(inputFile))

}