package main

import (
	"fmt"
	"os"
	request "swipe/internal/core/request"
	response "swipe/internal/core/response"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs at least url as argument")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Println("Sending request to " + url + "\n")
	req := request.Create(url)
	r, e := req.Get()

	res := response.Create(r, e)
	res.Print()

}
